// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1ERC721BridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"messenger\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_contract(CrossDomainMessenger)1007\"},{\"astId\":1003,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_uint256)48_storage\"},{\"astId\":1005,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\"},{\"astId\":1006,\"contract\":\"src/L1/L1ERC721Bridge.sol:L1ERC721Bridge\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_contract(SuperchainConfig)1008\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)48_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1007\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(SuperchainConfig)1008\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_mapping(t_uint256,t_bool)))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e mapping(uint256 =\u003e bool)))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1ERC721BridgeStorageLayout = new(solc.StorageLayout)

var L1ERC721BridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100d45760003560e01c8063761f449311610081578063aa5574521161005b578063aa55745214610248578063c0c53b8b1461025b578063c89701a21461026e57600080fd5b8063761f4493146101f35780637f46ddb214610206578063927ede2d1461022457600080fd5b806354fd4d50116100b257806354fd4d501461015e5780635c975abb146101a75780635d93a3fc146101bf57600080fd5b806335e80ab3146100d95780633687011a146101235780633cb747bf14610138575b600080fd5b6033546100f99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61013661013136600461100f565b61028e565b005b6000546100f99062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b61019a6040518060400160405280600581526020017f322e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161011a91906110fd565b6101af61033a565b604051901515815260200161011a565b6101af6101cd366004611117565b603260209081526000938452604080852082529284528284209052825290205460ff1681565b610136610201366004611158565b6103d3565b60015473ffffffffffffffffffffffffffffffffffffffff166100f9565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff166100f9565b6101366102563660046111f0565b610887565b610136610269366004611267565b610943565b6001546100f99073ffffffffffffffffffffffffffffffffffffffff1681565b333b15610322576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4552433732314272696467653a206163636f756e74206973206e6f742065787460448201527f65726e616c6c79206f776e65640000000000000000000000000000000000000060648201526084015b60405180910390fd5b6103328686333388888888610b36565b505050505050565b603354604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa1580156103aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ce91906112b2565b905090565b60005462010000900473ffffffffffffffffffffffffffffffffffffffff16331480156104b55750600154600054604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416936201000090930490921691636e296e45916004808201926020929091908290030181865afa158015610479573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049d91906112d4565b73ffffffffffffffffffffffffffffffffffffffff16145b610541576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f4552433732314272696467653a2066756e6374696f6e2063616e206f6e6c792060448201527f62652063616c6c65642066726f6d20746865206f7468657220627269646765006064820152608401610319565b61054961033a565b156105b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4c314552433732314272696467653a20706175736564000000000000000000006044820152606401610319565b3073ffffffffffffffffffffffffffffffffffffffff881603610655576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4c314552433732314272696467653a206c6f63616c20746f6b656e2063616e6e60448201527f6f742062652073656c66000000000000000000000000000000000000000000006064820152608401610319565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152603260209081526040808320938a1683529281528282208683529052205460ff161515600114610724576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4c314552433732314272696467653a20546f6b656e204944206973206e6f742060448201527f657363726f77656420696e20746865204c3120427269646765000000000000006064820152608401610319565b73ffffffffffffffffffffffffffffffffffffffff87811660008181526032602090815260408083208b8616845282528083208884529091529081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517f42842e0e000000000000000000000000000000000000000000000000000000008152306004820152918616602483015260448201859052906342842e0e90606401600060405180830381600087803b1580156107e457600080fd5b505af11580156107f8573d6000803e3d6000fd5b505050508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1f39bf6707b5d608453e0ae4c067b562bcc4c85c0f562ef5d2c774d2e7f131ac87878787604051610876949392919061133a565b60405180910390a450505050505050565b73ffffffffffffffffffffffffffffffffffffffff851661092a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433732314272696467653a206e667420726563697069656e742063616e6e60448201527f6f742062652061646472657373283029000000000000000000000000000000006064820152608401610319565b61093a8787338888888888610b36565b50505050505050565b600054610100900460ff16158080156109635750600054600160ff909116105b8061097d5750303b15801561097d575060005460ff166001145b610a09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610319565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a6757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416179055610ab18484610e7c565b8015610b1457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff8716610bd9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4c314552433732314272696467653a2072656d6f746520746f6b656e2063616e60448201527f6e6f7420626520616464726573732830290000000000000000000000000000006064820152608401610319565b600063761f449360e01b888a8989898888604051602401610c00979695949392919061137a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000959095169490941790935273ffffffffffffffffffffffffffffffffffffffff8c81166000818152603286528381208e8416825286528381208b82529095529382902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590517f23b872dd000000000000000000000000000000000000000000000000000000008152908a166004820152306024820152604481018890529092506323b872dd90606401600060405180830381600087803b158015610d4057600080fd5b505af1158015610d54573d6000803e3d6000fd5b50506000546001546040517f3dbb202b0000000000000000000000000000000000000000000000000000000081526201000090920473ffffffffffffffffffffffffffffffffffffffff9081169450633dbb202b9350610dbd92911690859089906004016113d7565b600060405180830381600087803b158015610dd757600080fd5b505af1158015610deb573d6000803e3d6000fd5b505050508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167fb7460e2a880f256ebef3406116ff3eee0cee51ebccdc2a40698f87ebb2e9c1a589898888604051610e69949392919061133a565b60405180910390a4505050505050505050565b600054610100900460ff16610f13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610319565b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff94851602179055600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909216179055565b73ffffffffffffffffffffffffffffffffffffffff81168114610faa57600080fd5b50565b803563ffffffff81168114610fc157600080fd5b919050565b60008083601f840112610fd857600080fd5b50813567ffffffffffffffff811115610ff057600080fd5b60208301915083602082850101111561100857600080fd5b9250929050565b60008060008060008060a0878903121561102857600080fd5b863561103381610f88565b9550602087013561104381610f88565b94506040870135935061105860608801610fad565b9250608087013567ffffffffffffffff81111561107457600080fd5b61108089828a01610fc6565b979a9699509497509295939492505050565b6000815180845260005b818110156110b85760208185018101518683018201520161109c565b818111156110ca576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006111106020830184611092565b9392505050565b60008060006060848603121561112c57600080fd5b833561113781610f88565b9250602084013561114781610f88565b929592945050506040919091013590565b600080600080600080600060c0888a03121561117357600080fd5b873561117e81610f88565b9650602088013561118e81610f88565b9550604088013561119e81610f88565b945060608801356111ae81610f88565b93506080880135925060a088013567ffffffffffffffff8111156111d157600080fd5b6111dd8a828b01610fc6565b989b979a50959850939692959293505050565b600080600080600080600060c0888a03121561120b57600080fd5b873561121681610f88565b9650602088013561122681610f88565b9550604088013561123681610f88565b94506060880135935061124b60808901610fad565b925060a088013567ffffffffffffffff8111156111d157600080fd5b60008060006060848603121561127c57600080fd5b833561128781610f88565b9250602084013561129781610f88565b915060408401356112a781610f88565b809150509250925092565b6000602082840312156112c457600080fd5b8151801515811461111057600080fd5b6000602082840312156112e657600080fd5b815161111081610f88565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526060604082015260006113706060830184866112f1565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526113ca60c0830184866112f1565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006114066060830185611092565b905063ffffffff8316604083015294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1ERC721BridgeStorageLayoutJSON), L1ERC721BridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1ERC721Bridge"] = L1ERC721BridgeStorageLayout
	deployedBytecodes["L1ERC721Bridge"] = L1ERC721BridgeDeployedBin
	immutableReferences["L1ERC721Bridge"] = false
}
