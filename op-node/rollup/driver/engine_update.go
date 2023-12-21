package driver

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// isDepositTx checks an opaqueTx to determine if it is a Deposit Transaction
// It has to return an error in the case the transaction is empty
func isDepositTx(opaqueTx eth.Data) (bool, error) {
	if len(opaqueTx) == 0 {
		return false, errors.New("empty transaction")
	}
	return opaqueTx[0] == types.DepositTxType, nil
}

// lastDeposit finds the index of last deposit at the start of the transactions.
// It walks the transactions from the start until it finds a non-deposit tx.
// An error is returned if any looked at transaction cannot be decoded
func lastDeposit(txns []eth.Data) (int, error) {
	var lastDeposit int
	for i, tx := range txns {
		deposit, err := isDepositTx(tx)
		if err != nil {
			return 0, fmt.Errorf("invalid transaction at idx %d", i)
		}
		if deposit {
			lastDeposit = i
		} else {
			break
		}
	}
	return lastDeposit, nil
}

func sanityCheckPayload(payload *eth.ExecutionPayload) error {
	// Sanity check payload before inserting it
	if len(payload.Transactions) == 0 {
		return errors.New("no transactions in returned payload")
	}
	if payload.Transactions[0][0] != types.DepositTxType {
		return fmt.Errorf("first transaction was not deposit tx. Got %v", payload.Transactions[0][0])
	}
	// Ensure that the deposits are first
	lastDeposit, err := lastDeposit(payload.Transactions)
	if err != nil {
		return fmt.Errorf("failed to find last deposit: %w", err)
	}
	// Ensure no deposits after last deposit
	for i := lastDeposit + 1; i < len(payload.Transactions); i++ {
		tx := payload.Transactions[i]
		deposit, err := isDepositTx(tx)
		if err != nil {
			return fmt.Errorf("failed to decode transaction idx %d: %w", i, err)
		}
		if deposit {
			return fmt.Errorf("deposit tx (%d) after other tx in l2 block with prev deposit at idx %d", i, lastDeposit)
		}
	}
	return nil
}

// StartPayload starts an execution payload building process in the provided Engine, with the given attributes.
// The severity of the error is distinguished to determine whether the same payload attributes may be re-attempted later.
func StartPayload(ctx context.Context, eng ExecEngine, fc eth.ForkchoiceState, attrs *eth.PayloadAttributes) (id eth.PayloadID, errType derive.BlockInsertionErrType, err error) {
	fcRes, err := eng.ForkchoiceUpdate(ctx, &fc, attrs)
	if err != nil {
		var inputErr eth.InputError
		if errors.As(err, &inputErr) {
			switch inputErr.Code {
			case eth.InvalidForkchoiceState:
				return eth.PayloadID{}, derive.BlockInsertPrestateErr, fmt.Errorf("pre-block-creation forkchoice update was inconsistent with engine, need reset to resolve: %w", inputErr.Unwrap())
			case eth.InvalidPayloadAttributes:
				return eth.PayloadID{}, derive.BlockInsertPayloadErr, fmt.Errorf("payload attributes are not valid, cannot build block: %w", inputErr.Unwrap())
			default:
				return eth.PayloadID{}, derive.BlockInsertPrestateErr, fmt.Errorf("unexpected error code in forkchoice-updated response: %w", err)
			}
		} else {
			return eth.PayloadID{}, derive.BlockInsertTemporaryErr, fmt.Errorf("failed to create new block via forkchoice: %w", err)
		}
	}

	switch fcRes.PayloadStatus.Status {
	// TODO(proto): snap sync - specify explicit different error type if node is syncing
	case eth.ExecutionInvalid, eth.ExecutionInvalidBlockHash:
		return eth.PayloadID{}, derive.BlockInsertPayloadErr, eth.ForkchoiceUpdateErr(fcRes.PayloadStatus)
	case eth.ExecutionValid:
		id := fcRes.PayloadID
		if id == nil {
			return eth.PayloadID{}, derive.BlockInsertTemporaryErr, errors.New("nil id in forkchoice result when expecting a valid ID")
		}
		return *id, derive.BlockInsertOK, nil
	default:
		return eth.PayloadID{}, derive.BlockInsertTemporaryErr, eth.ForkchoiceUpdateErr(fcRes.PayloadStatus)
	}
}

// ConfirmPayload ends an execution payload building process in the provided Engine, and persists the payload as the canonical head.
// If updateSafe is true, then the payload will also be recognized as safe-head at the same time.
// The severity of the error is distinguished to determine whether the payload was valid and can become canonical.
func ConfirmPayload(ctx context.Context, log log.Logger, eng ExecEngine, fc eth.ForkchoiceState, id eth.PayloadID, updateSafe bool) (out *eth.ExecutionPayload, errTyp derive.BlockInsertionErrType, err error) {
	payload, err := eng.GetPayload(ctx, id)
	if err != nil {
		// even if it is an input-error (unknown payload ID), it is temporary, since we will re-attempt the full payload building, not just the retrieval of the payload.
		return nil, derive.BlockInsertTemporaryErr, fmt.Errorf("failed to get execution payload: %w", err)
	}
	if err := sanityCheckPayload(payload); err != nil {
		return nil, derive.BlockInsertPayloadErr, err
	}

	status, err := eng.NewPayload(ctx, payload)
	if err != nil {
		return nil, derive.BlockInsertTemporaryErr, fmt.Errorf("failed to insert execution payload: %w", err)
	}
	if status.Status == eth.ExecutionInvalid || status.Status == eth.ExecutionInvalidBlockHash {
		return nil, derive.BlockInsertPayloadErr, eth.NewPayloadErr(payload, status)
	}
	if status.Status != eth.ExecutionValid {
		return nil, derive.BlockInsertTemporaryErr, eth.NewPayloadErr(payload, status)
	}

	fc.HeadBlockHash = payload.BlockHash
	if updateSafe {
		fc.SafeBlockHash = payload.BlockHash
	}
	fcRes, err := eng.ForkchoiceUpdate(ctx, &fc, nil)
	if err != nil {
		var inputErr eth.InputError
		if errors.As(err, &inputErr) {
			switch inputErr.Code {
			case eth.InvalidForkchoiceState:
				// if we succeed to update the forkchoice pre-payload, but fail post-payload, then it is a payload error
				return nil, derive.BlockInsertPayloadErr, fmt.Errorf("post-block-creation forkchoice update was inconsistent with engine, need reset to resolve: %w", inputErr.Unwrap())
			default:
				return nil, derive.BlockInsertPrestateErr, fmt.Errorf("unexpected error code in forkchoice-updated response: %w", err)
			}
		} else {
			return nil, derive.BlockInsertTemporaryErr, fmt.Errorf("failed to make the new L2 block canonical via forkchoice: %w", err)
		}
	}
	if fcRes.PayloadStatus.Status != eth.ExecutionValid {
		return nil, derive.BlockInsertPayloadErr, eth.ForkchoiceUpdateErr(fcRes.PayloadStatus)
	}
	log.Info("inserted block", "hash", payload.BlockHash, "number", uint64(payload.BlockNumber),
		"state_root", payload.StateRoot, "timestamp", uint64(payload.Timestamp), "parent", payload.ParentHash,
		"prev_randao", payload.PrevRandao, "fee_recipient", payload.FeeRecipient,
		"txs", len(payload.Transactions), "update_safe", updateSafe)
	return payload, derive.BlockInsertOK, nil
}
