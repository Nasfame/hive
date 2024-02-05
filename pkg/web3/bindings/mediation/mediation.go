// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mediation

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId   string
	Members  SharedStructsDealMembers
	Timeouts SharedStructsDealTimeouts
	Pricing  SharedStructsDealPricing
}

// SharedStructsDealMembers is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealMembers struct {
	Solver           common.Address
	JobCreator       common.Address
	ResourceProvider common.Address
	Mediators        []common.Address
}

// SharedStructsDealPricing is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealPricing struct {
	InstructionPrice          *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsDealTimeout is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeout struct {
	Timeout    *big.Int
	Collateral *big.Int
}

// SharedStructsDealTimeouts is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDealTimeouts struct {
	Agree          SharedStructsDealTimeout
	SubmitResults  SharedStructsDealTimeout
	JudgeResults   SharedStructsDealTimeout
	MediateResults SharedStructsDealTimeout
}

// MediationMetaData contains all meta data concerning the Mediation contract.
var MediationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"MediationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getMediator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"deal\",\"type\":\"tuple\"}],\"name\":\"mediationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b5061004761003c61004c60201b60201c565b61005460201b60201c565b610118565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611a17806101276000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b14610116578063a2bffa0014610134578063a470295814610164578063c57380a21461016e578063f2fde38b1461018c578063f3d3d448146101a8576100a9565b806370bea207146100ae578063715018a6146100ca57806380ffdfe0146100d45780638129fc1c146100f0578063824518aa146100fa575b600080fd5b6100c860048036038101906100c39190611268565b6101c4565b005b6100d261034e565b005b6100ee60048036038101906100e991906112b1565b610362565b005b6100f8610530565b005b610114600480360381019061010f91906112b1565b610669565b005b61011e610837565b60405161012b9190611309565b60405180910390f35b61014e600480360381019061014991906112b1565b610860565b60405161015b9190611309565b60405180910390f35b61016c6108a8565b005b6101766108cd565b6040516101839190611309565b60405180910390f35b6101a660048036038101906101a19190611324565b6108f7565b005b6101c260048036038101906101bd9190611324565b61097a565b005b6101cc610a84565b506000816020015160600151514283600001516040516020016101f09291906113ec565b6040516020818303038152906040528051906020012060001c6102139190611443565b90506000826020015160600151828151811061023257610231611474565b5b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036102ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a290611500565b60405180910390fd5b80600284600001516040516102c09190611520565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507ffd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086836000015182604051610341929190611570565b60405180910390a1505050565b610356610bb4565b6103606000610c32565b565b600073ffffffffffffffffffffffffffffffffffffffff1660028260405161038a9190611520565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361040f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040690611500565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff166002826040516104369190611520565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146104bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b2906115ec565b60405180910390fd5b6104c36108cd565b73ffffffffffffffffffffffffffffffffffffffff166380ffdfe0826040518263ffffffff1660e01b81526004016104fb919061160c565b600060405180830381600087803b15801561051557600080fd5b505af1158015610529573d6000803e3d6000fd5b5050505050565b6000600160169054906101000a900460ff16159050808015610563575060018060159054906101000a900460ff1660ff16105b80610591575061057230610cf6565b158015610590575060018060159054906101000a900460ff1660ff16145b5b6105d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c7906116a0565b60405180910390fd5b60018060156101000a81548160ff021916908360ff160217905550801561060c5760018060166101000a81548160ff0219169083151502179055505b8015610666576000600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161065d9190611712565b60405180910390a15b50565b600073ffffffffffffffffffffffffffffffffffffffff166002826040516106919190611520565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070d90611500565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff1660028260405161073d9190611520565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b9906115ec565b60405180910390fd5b6107ca6108cd565b73ffffffffffffffffffffffffffffffffffffffff1663824518aa826040518263ffffffff1660e01b8152600401610802919061160c565b600060405180830381600087803b15801561081c57600080fd5b505af1158015610830573d6000803e3d6000fd5b5050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006002826040516108729190611520565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6108b0610bb4565b6000600160146101000a81548160ff021916908315150217905550565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6108ff610bb4565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361096e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109659061179f565b60405180910390fd5b61097781610c32565b50565b610982610bb4565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e890611831565b60405180910390fd5b600160149054906101000a900460ff16610a40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a37906118c3565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610b16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0d90611831565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b57610d19565b73ffffffffffffffffffffffffffffffffffffffff1614610bad576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba490611955565b60405180910390fd5b6001905090565b610bbc610d19565b73ffffffffffffffffffffffffffffffffffffffff16610bda610837565b73ffffffffffffffffffffffffffffffffffffffff1614610c30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c27906119c1565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d8382610d3a565b810181811067ffffffffffffffff82111715610da257610da1610d4b565b5b80604052505050565b6000610db5610d21565b9050610dc18282610d7a565b919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff821115610df057610def610d4b565b5b610df982610d3a565b9050602081019050919050565b82818337600083830152505050565b6000610e28610e2384610dd5565b610dab565b905082815260208101848484011115610e4457610e43610dd0565b5b610e4f848285610e06565b509392505050565b600082601f830112610e6c57610e6b610dcb565b5b8135610e7c848260208601610e15565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610eb082610e85565b9050919050565b610ec081610ea5565b8114610ecb57600080fd5b50565b600081359050610edd81610eb7565b92915050565b600067ffffffffffffffff821115610efe57610efd610d4b565b5b602082029050602081019050919050565b600080fd5b6000610f27610f2284610ee3565b610dab565b90508083825260208201905060208402830185811115610f4a57610f49610f0f565b5b835b81811015610f735780610f5f8882610ece565b845260208401935050602081019050610f4c565b5050509392505050565b600082601f830112610f9257610f91610dcb565b5b8135610fa2848260208601610f14565b91505092915050565b600060808284031215610fc157610fc0610d35565b5b610fcb6080610dab565b90506000610fdb84828501610ece565b6000830152506020610fef84828501610ece565b602083015250604061100384828501610ece565b604083015250606082013567ffffffffffffffff81111561102757611026610dc6565b5b61103384828501610f7d565b60608301525092915050565b6000819050919050565b6110528161103f565b811461105d57600080fd5b50565b60008135905061106f81611049565b92915050565b60006040828403121561108b5761108a610d35565b5b6110956040610dab565b905060006110a584828501611060565b60008301525060206110b984828501611060565b60208301525092915050565b600061010082840312156110dc576110db610d35565b5b6110e66080610dab565b905060006110f684828501611075565b600083015250604061110a84828501611075565b602083015250608061111e84828501611075565b60408301525060c061113284828501611075565b60608301525092915050565b60006080828403121561115457611153610d35565b5b61115e6080610dab565b9050600061116e84828501611060565b600083015250602061118284828501611060565b602083015250604061119684828501611060565b60408301525060606111aa84828501611060565b60608301525092915050565b60006101c082840312156111cd576111cc610d35565b5b6111d76080610dab565b9050600082013567ffffffffffffffff8111156111f7576111f6610dc6565b5b61120384828501610e57565b600083015250602082013567ffffffffffffffff81111561122757611226610dc6565b5b61123384828501610fab565b6020830152506040611247848285016110c5565b60408301525061014061125c8482850161113e565b60608301525092915050565b60006020828403121561127e5761127d610d2b565b5b600082013567ffffffffffffffff81111561129c5761129b610d30565b5b6112a8848285016111b6565b91505092915050565b6000602082840312156112c7576112c6610d2b565b5b600082013567ffffffffffffffff8111156112e5576112e4610d30565b5b6112f184828501610e57565b91505092915050565b61130381610ea5565b82525050565b600060208201905061131e60008301846112fa565b92915050565b60006020828403121561133a57611339610d2b565b5b600061134884828501610ece565b91505092915050565b6000819050919050565b61136c6113678261103f565b611351565b82525050565b600081519050919050565b600081905092915050565b60005b838110156113a657808201518184015260208101905061138b565b838111156113b5576000848401525b50505050565b60006113c682611372565b6113d0818561137d565b93506113e0818560208601611388565b80840191505092915050565b60006113f8828561135b565b60208201915061140882846113bb565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061144e8261103f565b91506114598361103f565b92508261146957611468611414565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f6d65646961746f722063616e6e6f742062652030783000000000000000000000600082015250565b60006114ea6016836114a3565b91506114f5826114b4565b602082019050919050565b60006020820190508181036000830152611519816114dd565b9050919050565b600061152c82846113bb565b915081905092915050565b600061154282611372565b61154c81856114a3565b935061155c818560208601611388565b61156581610d3a565b840191505092915050565b6000604082019050818103600083015261158a8185611537565b905061159960208301846112fa565b9392505050565b7f74782e6f726967696e206d75737420626520746865206d65646961746f720000600082015250565b60006115d6601e836114a3565b91506115e1826115a0565b602082019050919050565b60006020820190508181036000830152611605816115c9565b9050919050565b600060208201905081810360008301526116268184611537565b905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600061168a602e836114a3565b91506116958261162e565b604082019050919050565b600060208201905081810360008301526116b98161167d565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b60006116fc6116f76116f2846116c0565b6116d7565b6116ca565b9050919050565b61170c816116e1565b82525050565b60006020820190506117276000830184611703565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006117896026836114a3565b91506117948261172d565b604082019050919050565b600060208201905081810360008301526117b88161177c565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c657220616460008201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b600061181b6035836114a3565b9150611826826117bf565b604082019050919050565b6000602082019050818103600083015261184a8161180e565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e7460008201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b60006118ad6039836114a3565b91506118b882611851565b604082019050919050565b600060208201905081810360008301526118dc816118a0565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e7460008201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b600061193f603b836114a3565b915061194a826118e3565b604082019050919050565b6000602082019050818103600083015261196e81611932565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006119ab6020836114a3565b91506119b682611975565b602082019050919050565b600060208201905081810360008301526119da8161199e565b905091905056fea2646970667358221220bbbae59a29653a16d4dbb91235a1a0883d457ce88ab211d4366cf294b8c1bdd064736f6c634300080f0033",
}

// MediationABI is the input ABI used to generate the binding from.
// Deprecated: Use MediationMetaData.ABI instead.
var MediationABI = MediationMetaData.ABI

// MediationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MediationMetaData.Bin instead.
var MediationBin = MediationMetaData.Bin

// DeployMediation deploys a new Ethereum contract, binding an instance of Mediation to it.
func DeployMediation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mediation, error) {
	parsed, err := MediationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MediationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mediation{MediationCaller: MediationCaller{contract: contract}, MediationTransactor: MediationTransactor{contract: contract}, MediationFilterer: MediationFilterer{contract: contract}}, nil
}

// Mediation is an auto generated Go binding around an Ethereum contract.
type Mediation struct {
	MediationCaller     // Read-only binding to the contract
	MediationTransactor // Write-only binding to the contract
	MediationFilterer   // Log filterer for contract events
}

// MediationCaller is an auto generated read-only Go binding around an Ethereum contract.
type MediationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MediationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MediationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MediationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MediationSession struct {
	Contract     *Mediation        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MediationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MediationCallerSession struct {
	Contract *MediationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MediationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MediationTransactorSession struct {
	Contract     *MediationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MediationRaw is an auto generated low-level Go binding around an Ethereum contract.
type MediationRaw struct {
	Contract *Mediation // Generic contract binding to access the raw methods on
}

// MediationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MediationCallerRaw struct {
	Contract *MediationCaller // Generic read-only contract binding to access the raw methods on
}

// MediationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MediationTransactorRaw struct {
	Contract *MediationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMediation creates a new instance of Mediation, bound to a specific deployed contract.
func NewMediation(address common.Address, backend bind.ContractBackend) (*Mediation, error) {
	contract, err := bindMediation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mediation{MediationCaller: MediationCaller{contract: contract}, MediationTransactor: MediationTransactor{contract: contract}, MediationFilterer: MediationFilterer{contract: contract}}, nil
}

// NewMediationCaller creates a new read-only instance of Mediation, bound to a specific deployed contract.
func NewMediationCaller(address common.Address, caller bind.ContractCaller) (*MediationCaller, error) {
	contract, err := bindMediation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MediationCaller{contract: contract}, nil
}

// NewMediationTransactor creates a new write-only instance of Mediation, bound to a specific deployed contract.
func NewMediationTransactor(address common.Address, transactor bind.ContractTransactor) (*MediationTransactor, error) {
	contract, err := bindMediation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MediationTransactor{contract: contract}, nil
}

// NewMediationFilterer creates a new log filterer instance of Mediation, bound to a specific deployed contract.
func NewMediationFilterer(address common.Address, filterer bind.ContractFilterer) (*MediationFilterer, error) {
	contract, err := bindMediation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MediationFilterer{contract: contract}, nil
}

// bindMediation binds a generic wrapper to an already deployed contract.
func bindMediation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MediationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mediation *MediationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mediation.Contract.MediationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mediation *MediationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.Contract.MediationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mediation *MediationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mediation.Contract.MediationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mediation *MediationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mediation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mediation *MediationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mediation *MediationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mediation.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Mediation *MediationCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mediation.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Mediation *MediationSession) GetControllerAddress() (common.Address, error) {
	return _Mediation.Contract.GetControllerAddress(&_Mediation.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Mediation *MediationCallerSession) GetControllerAddress() (common.Address, error) {
	return _Mediation.Contract.GetControllerAddress(&_Mediation.CallOpts)
}

// GetMediator is a free data retrieval call binding the contract method 0xa2bffa00.
//
// Solidity: function getMediator(string dealId) view returns(address)
func (_Mediation *MediationCaller) GetMediator(opts *bind.CallOpts, dealId string) (common.Address, error) {
	var out []interface{}
	err := _Mediation.contract.Call(opts, &out, "getMediator", dealId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMediator is a free data retrieval call binding the contract method 0xa2bffa00.
//
// Solidity: function getMediator(string dealId) view returns(address)
func (_Mediation *MediationSession) GetMediator(dealId string) (common.Address, error) {
	return _Mediation.Contract.GetMediator(&_Mediation.CallOpts, dealId)
}

// GetMediator is a free data retrieval call binding the contract method 0xa2bffa00.
//
// Solidity: function getMediator(string dealId) view returns(address)
func (_Mediation *MediationCallerSession) GetMediator(dealId string) (common.Address, error) {
	return _Mediation.Contract.GetMediator(&_Mediation.CallOpts, dealId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mediation *MediationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mediation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mediation *MediationSession) Owner() (common.Address, error) {
	return _Mediation.Contract.Owner(&_Mediation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mediation *MediationCallerSession) Owner() (common.Address, error) {
	return _Mediation.Contract.Owner(&_Mediation.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Mediation *MediationTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Mediation *MediationSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Mediation.Contract.DisableChangeControllerAddress(&_Mediation.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Mediation *MediationTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Mediation.Contract.DisableChangeControllerAddress(&_Mediation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mediation *MediationTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mediation *MediationSession) Initialize() (*types.Transaction, error) {
	return _Mediation.Contract.Initialize(&_Mediation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mediation *MediationTransactorSession) Initialize() (*types.Transaction, error) {
	return _Mediation.Contract.Initialize(&_Mediation.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Mediation *MediationTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Mediation *MediationSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationAcceptResult(&_Mediation.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Mediation *MediationTransactorSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationAcceptResult(&_Mediation.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Mediation *MediationTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Mediation *MediationSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRejectResult(&_Mediation.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Mediation *MediationTransactorSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRejectResult(&_Mediation.TransactOpts, dealId)
}

// MediationRequest is a paid mutator transaction binding the contract method 0x70bea207.
//
// Solidity: function mediationRequest((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationTransactor) MediationRequest(opts *bind.TransactOpts, deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "mediationRequest", deal)
}

// MediationRequest is a paid mutator transaction binding the contract method 0x70bea207.
//
// Solidity: function mediationRequest((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationSession) MediationRequest(deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRequest(&_Mediation.TransactOpts, deal)
}

// MediationRequest is a paid mutator transaction binding the contract method 0x70bea207.
//
// Solidity: function mediationRequest((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)) deal) returns()
func (_Mediation *MediationTransactorSession) MediationRequest(deal SharedStructsDeal) (*types.Transaction, error) {
	return _Mediation.Contract.MediationRequest(&_Mediation.TransactOpts, deal)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediation *MediationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediation *MediationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mediation.Contract.RenounceOwnership(&_Mediation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mediation *MediationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mediation.Contract.RenounceOwnership(&_Mediation.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Mediation *MediationTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Mediation *MediationSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.SetControllerAddress(&_Mediation.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Mediation *MediationTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.SetControllerAddress(&_Mediation.TransactOpts, _controllerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mediation *MediationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mediation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mediation *MediationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.TransferOwnership(&_Mediation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mediation *MediationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mediation.Contract.TransferOwnership(&_Mediation.TransactOpts, newOwner)
}

// MediationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mediation contract.
type MediationInitializedIterator struct {
	Event *MediationInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MediationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediationInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MediationInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MediationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediationInitialized represents a Initialized event raised by the Mediation contract.
type MediationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mediation *MediationFilterer) FilterInitialized(opts *bind.FilterOpts) (*MediationInitializedIterator, error) {

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MediationInitializedIterator{contract: _Mediation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mediation *MediationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MediationInitialized) (event.Subscription, error) {

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediationInitialized)
				if err := _Mediation.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mediation *MediationFilterer) ParseInitialized(log types.Log) (*MediationInitialized, error) {
	event := new(MediationInitialized)
	if err := _Mediation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MediationMediationRequestedIterator is returned from FilterMediationRequested and is used to iterate over the raw logs and unpacked data for MediationRequested events raised by the Mediation contract.
type MediationMediationRequestedIterator struct {
	Event *MediationMediationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MediationMediationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediationMediationRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MediationMediationRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MediationMediationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediationMediationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediationMediationRequested represents a MediationRequested event raised by the Mediation contract.
type MediationMediationRequested struct {
	DealId   string
	Mediator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMediationRequested is a free log retrieval operation binding the contract event 0xfd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086.
//
// Solidity: event MediationRequested(string dealId, address mediator)
func (_Mediation *MediationFilterer) FilterMediationRequested(opts *bind.FilterOpts) (*MediationMediationRequestedIterator, error) {

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "MediationRequested")
	if err != nil {
		return nil, err
	}
	return &MediationMediationRequestedIterator{contract: _Mediation.contract, event: "MediationRequested", logs: logs, sub: sub}, nil
}

// WatchMediationRequested is a free log subscription operation binding the contract event 0xfd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086.
//
// Solidity: event MediationRequested(string dealId, address mediator)
func (_Mediation *MediationFilterer) WatchMediationRequested(opts *bind.WatchOpts, sink chan<- *MediationMediationRequested) (event.Subscription, error) {

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "MediationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediationMediationRequested)
				if err := _Mediation.contract.UnpackLog(event, "MediationRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMediationRequested is a log parse operation binding the contract event 0xfd3770121045f9427361660d6eaa8b07a2e45eca6964f5c4f041a28f21084086.
//
// Solidity: event MediationRequested(string dealId, address mediator)
func (_Mediation *MediationFilterer) ParseMediationRequested(log types.Log) (*MediationMediationRequested, error) {
	event := new(MediationMediationRequested)
	if err := _Mediation.contract.UnpackLog(event, "MediationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MediationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mediation contract.
type MediationOwnershipTransferredIterator struct {
	Event *MediationOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MediationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MediationOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MediationOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MediationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MediationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MediationOwnershipTransferred represents a OwnershipTransferred event raised by the Mediation contract.
type MediationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mediation *MediationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MediationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mediation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MediationOwnershipTransferredIterator{contract: _Mediation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mediation *MediationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MediationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mediation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MediationOwnershipTransferred)
				if err := _Mediation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mediation *MediationFilterer) ParseOwnershipTransferred(log types.Log) (*MediationOwnershipTransferred, error) {
	event := new(MediationOwnershipTransferred)
	if err := _Mediation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
