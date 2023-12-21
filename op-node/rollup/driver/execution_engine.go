package driver

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/log"
)

/*


The raw execution engine API provides 3 methods
1. FCU
2. GetPayload
3. NewPayload


Above that layer, we have a higher level set of abstraction

1. Insert Unsafe Payload
2. Insert Safe Attributes [Note: if the attributes type is middle of span batch, we don't actually update the safe head...]
3. Update Safe Head
4. Update Finalized Head
5. Start block building (as a sequencer)
6. Complete block building (as a sequencer) [note: This should change to GetBuiltBlock when doing gossip before insertion]
7. Cancel Block Building
8. UpdateUnsafeAndSafeHeads - after sync start w/ reorgs

The only sequence of calls we must prohibit is updating the unsafe head in the middle of block building.
This could happen via InsertUnsafePayload of InsertSafeAttributes.


The engine stores several pieces of information
1. Unsafe block head
2. Safe block head
3. Finalized block head
4. Span batch safe head?



Above the exec engine, we have the following components (probably from the state)
1. Unsafe payload buffer
2. derivation pipeline controller (which is only initialized after we have completed a sync)


*/

type ExecEngine interface {
	GetPayload(ctx context.Context, payloadId eth.PayloadID) (*eth.ExecutionPayload, error)
	ForkchoiceUpdate(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error)
	NewPayload(ctx context.Context, payload *eth.ExecutionPayload) (*eth.PayloadStatusV1, error)
}

var _ derive.EngineControl = (*EngineController)(nil)

type EngineController struct {
	engine ExecEngine // Underlying execution engine RPC
	log    log.Logger

	// Block Head State
	unsafeHead      eth.L2BlockRef
	pendingSafeHead eth.L2BlockRef
	safeHead        eth.L2BlockRef
	finalizedHead   eth.L2BlockRef

	// Building State
	buildingOnto         eth.L2BlockRef
	buildingID           eth.PayloadID
	buildingSafe         bool
	needForkchoiceUpdate bool
}

func NewEngineController(log log.Logger, engine ExecEngine) *EngineController {
	return &EngineController{
		engine: engine,
		log:    log,
		// TODO: Block state. May need to fetch from the L2 Source / have that info be provided.
	}
}

type PayloadCandidate struct {
	parent eth.BlockID
	id     eth.PayloadID
}

// State Getters

func (e *EngineController) UnsafeL2Head() eth.L2BlockRef {
	return e.unsafeHead
}

func (e *EngineController) PendingSafeL2Head() eth.L2BlockRef {
	return e.pendingSafeHead
}

func (e *EngineController) SafeL2Head() eth.L2BlockRef {
	return e.safeHead
}

func (e *EngineController) Finalized() eth.L2BlockRef {
	return e.finalizedHead
}

// V2 APIS

func (e *EngineController) InsertUnsafePayload(ctx context.Context, payload *eth.ExecutionPayload) error {
	// NewPayload
	// FCU w/ updated unsafe head
	return nil
}

func (e *EngineController) InsertSafeAttributes(ctx context.Context, attributes *eth.PayloadAttributes, attrType derive.AttributesType) error {
	// FCU to build (+ rollback unsafe head)
	// NewPayload
	// If attrType == SpanBatchMiddleAttrs, using pending safe head instead
	// FCU updated unsafe & safe head
	return nil
}

func (e *EngineController) UpdateSafeHead(ctx context.Context, safe eth.BlockID) error {
	// FCU updated safe head
	return nil
}

func (e *EngineController) UpdateFinalizedHead(ctx context.Context, finalized eth.BlockID) error {
	// FCU updated safe head
	return nil
}

func (*EngineController) UpdateUnsafeAndSafeHeads(unsafe eth.BlockID, safe eth.BlockID) error {
	// FCU updated safe head
	// Wipe building + pending state
	return nil
}

// V1 APIS

func (e *EngineController) StartPayload(ctx context.Context, parent eth.L2BlockRef, attrs *eth.PayloadAttributes, updateSafe bool) (errType derive.BlockInsertionErrType, err error) {
	// if eq.isEngineSyncing() {
	// 	return BlockInsertTemporaryErr, fmt.Errorf("engine is in progess of p2p sync")
	// }
	// if eq.buildingID != (eth.PayloadID{}) {
	// 	eq.log.Warn("did not finish previous block building, starting new building now", "prev_onto", eq.buildingOnto, "prev_payload_id", eq.buildingID, "new_onto", parent)
	// 	// TODO: maybe worth it to force-cancel the old payload ID here.
	// }
	// fc := eth.ForkchoiceState{
	// 	HeadBlockHash:      parent.Hash,
	// 	SafeBlockHash:      eq.SafeL2Head().Hash,
	// 	FinalizedBlockHash: eq.finalized.Hash,
	// }
	// id, errTyp, err := StartPayload(ctx, eq.engine, fc, attrs)
	// if err != nil {
	// 	return errTyp, err
	// }
	// eq.buildingID = id
	// eq.buildingSafe = updateSafe
	// eq.buildingOnto = parent
	return derive.BlockInsertOK, nil
}

func (e *EngineController) ConfirmPayload(ctx context.Context) (out *eth.ExecutionPayload, errTyp derive.BlockInsertionErrType, err error) {
	// if eq.buildingID == (eth.PayloadID{}) {
	// 	return nil, BlockInsertPrestateErr, fmt.Errorf("cannot complete payload building: not currently building a payload")
	// }
	// if eq.buildingOnto.Hash != eq.UnsafeL2Head().Hash { // E.g. when safe-attributes consolidation fails, it will drop the existing work.
	// 	eq.log.Warn("engine is building block that reorgs previous unsafe head", "onto", eq.buildingOnto, "unsafe", eq.UnsafeL2Head())
	// }
	// fc := eth.ForkchoiceState{
	// 	HeadBlockHash:      common.Hash{}, // gets overridden
	// 	SafeBlockHash:      eq.SafeL2Head().Hash,
	// 	FinalizedBlockHash: eq.finalized.Hash,
	// }
	// // Update the safe head if the payload is built with the last attributes in the batch.
	// updateSafe := eq.buildingSafe && eq.safeAttributes != nil && eq.safeAttributes.isLastInSpan
	// payload, errTyp, err := ConfirmPayload(ctx, eq.log, eq.engine, fc, eq.buildingID, updateSafe)
	// if err != nil {
	// 	return nil, errTyp, fmt.Errorf("failed to complete building on top of L2 chain %s, id: %s, error (%d): %w", eq.buildingOnto, eq.buildingID, errTyp, err)
	// }
	// ref, err := PayloadToBlockRef(payload, &eq.cfg.Genesis)
	// if err != nil {
	// 	return nil, BlockInsertPayloadErr, NewResetError(fmt.Errorf("failed to decode L2 block ref from payload: %w", err))
	// }

	// eq.UnsafeL2Head() = ref
	// eq.engineSyncTarget = ref
	// eq.metrics.RecordL2Ref("l2_unsafe", ref)
	// eq.metrics.RecordL2Ref("l2_engineSyncTarget", ref)

	// if eq.buildingSafe {
	// 	eq.pendingSafeHead = ref
	// 	if updateSafe {
	// 		eq.SafeL2Head() = ref
	// 		eq.postProcessSafeL2()
	// 		eq.metrics.RecordL2Ref("l2_safe", ref)
	// 	}
	// }
	// eq.resetBuildingState()
	// return payload, BlockInsertOK, nil
	return nil, derive.BlockInsertOK, nil
}

func (e *EngineController) CancelPayload(ctx context.Context, force bool) error {
	// if eq.buildingID == (eth.PayloadID{}) { // only cancel if there is something to cancel.
	// 	return nil
	// }
	// // the building job gets wrapped up as soon as the payload is retrieved, there's no explicit cancel in the Engine API
	// eq.log.Error("cancelling old block sealing job", "payload", eq.buildingID)
	// _, err := eq.engine.GetPayload(ctx, eq.buildingID)
	// if err != nil {
	// 	eq.log.Error("failed to cancel block building job", "payload", eq.buildingID, "err", err)
	// 	if !force {
	// 		return err
	// 	}
	// }
	// eq.resetBuildingState()
	return nil
}

func (e *EngineController) BuildingPayload() (onto eth.L2BlockRef, id eth.PayloadID, safe bool) {
	return e.buildingOnto, e.buildingID, e.buildingSafe
}

func (e *EngineController) resetBuildingState() {
	e.buildingID = eth.PayloadID{}
	e.buildingOnto = eth.L2BlockRef{}
	e.buildingSafe = false
}

// checkNewPayloadStatus checks returned status of engine_newPayloadV1 request for next unsafe payload.
// It returns true if the status is acceptable.
func (e *EngineController) checkNewPayloadStatus(status eth.ExecutePayloadStatus) bool {
	// if e.syncCfg.SyncMode == sync.ELSync {
	// 	// Allow SYNCING and ACCEPTED if engine EL sync is enabled
	// 	return status == eth.ExecutionValid || status == eth.ExecutionSyncing || status == eth.ExecutionAccepted
	// }
	return status == eth.ExecutionValid
}

// checkForkchoiceUpdatedStatus checks returned status of engine_forkchoiceUpdatedV1 request for next unsafe payload.
// It returns true if the status is acceptable.
func (e *EngineController) checkForkchoiceUpdatedStatus(status eth.ExecutePayloadStatus) bool {
	// if e.syncCfg.SyncMode == sync.ELSync {
	// 	// Allow SYNCING if engine P2P sync is enabled
	// 	return status == eth.ExecutionValid || status == eth.ExecutionSyncing
	// }
	return status == eth.ExecutionValid
}
