// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jobcreator

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

// JobcreatorMetaData contains all meta data concerning the Jobcreator contract.
var JobcreatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"calling_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"}],\"name\":\"JobAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextJobID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"runJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"setRequiredDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"}],\"name\":\"submitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff02191690831515021790555034801561002a57600080fd5b5061004761003c61004c60201b60201c565b61005460201b60201c565b610118565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611edd806101276000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063a470295811610097578063d2a715c011610066578063d2a715c014610226578063f2fde38b14610244578063f3d3d44814610260578063fb7cfdd71461027c576100f5565b8063a4702958146101b2578063c4d66de8146101bc578063c57380a2146101d8578063c75555fa146101f6576100f5565b806358e56db4116100d357806358e56db4146101505780636c0f1f581461016e578063715018a61461018a5780638da5cb5b14610194576100f5565b806310fe9ae8146100fa57806326a4e8d2146101185780634c526c7614610134575b600080fd5b61010261029a565b60405161010f9190610eef565b60405180910390f35b610132600480360381019061012d9190610f4a565b6102c4565b005b61014e60048036038101906101499190610fad565b6103e2565b005b610158610438565b6040516101659190610fe9565b60405180910390f35b6101886004803603810190610183919061114a565b610442565b005b610192610542565b005b61019c610556565b6040516101a99190610eef565b60405180910390f35b6101ba61057f565b005b6101d660048036038101906101d19190610f4a565b6105a4565b005b6101e06106ef565b6040516101ed9190610eef565b60405180910390f35b610210600480360381019061020b91906112bb565b610719565b60405161021d9190610fe9565b60405180910390f35b61022e61099e565b60405161023b9190610fe9565b60405180910390f35b61025e60048036038101906102599190610f4a565b6109a4565b005b61027a60048036038101906102759190610f4a565b610a27565b005b610284610b31565b6040516102919190610fe9565b60405180910390f35b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102cc610b37565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361033b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610332906113a3565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6103ea610bb5565b506000811161042e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104259061140f565b60405180910390fd5b8060048190555050565b6000600454905090565b61044a610bb5565b50600060066000858152602001908152602001600020905060008160000154036104a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a09061147b565b60405180910390fd5b8060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636c0f1f588585856040518463ffffffff1660e01b815260040161050a93929190611512565b600060405180830381600087803b15801561052457600080fd5b505af1158015610538573d6000803e3d6000fd5b5050505050505050565b61054a610b37565b6105546000610ce5565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610587610b37565b6000600160146101000a81548160ff021916908315150217905550565b6000600160169054906101000a900460ff161590508080156105d7575060018060159054906101000a900460ff1660ff16105b8061060557506105e630610da9565b158015610604575060018060159054906101000a900460ff1660ff16145b5b610644576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063b906115c9565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156106805760018060166101000a81548160ff0219169083151502179055505b610689826102c4565b600060058190555080156106eb576000600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516106e2919061163b565b60405180910390a15b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600454600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e846107656106ef565b6040518363ffffffff1660e01b8152600401610782929190611656565b602060405180830381865afa15801561079f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c39190611694565b1015610804576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fb9061170d565b60405180910390fd5b6001600554610813919061175c565b6005819055506040518060a0016040528060055481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152602001858152602001848152506006600060055481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301908161093091906119b4565b50608082015181600401908051906020019061094d929190610dd4565b509050507faa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a6005543384878760405161098a959493929190611b92565b60405180910390a160055490509392505050565b60055481565b6109ac610b37565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1290611c65565b60405180910390fd5b610a2481610ce5565b50565b610a2f610b37565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9590611cf7565b60405180910390fd5b600160149054906101000a900460ff16610aed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae490611d89565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60045481565b610b3f610dcc565b73ffffffffffffffffffffffffffffffffffffffff16610b5d610556565b73ffffffffffffffffffffffffffffffffffffffff1614610bb3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610baa90611df5565b60405180910390fd5b565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3e90611cf7565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610c88610dcc565b73ffffffffffffffffffffffffffffffffffffffff1614610cde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd590611e87565b60405180910390fd5b6001905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600033905090565b828054828255906000526020600020908101928215610e1c579160200282015b82811115610e1b578251829081610e0b91906119b4565b5091602001919060010190610df4565b5b509050610e299190610e2d565b5090565b5b80821115610e4d5760008181610e449190610e51565b50600101610e2e565b5090565b508054610e5d906117e1565b6000825580601f10610e6f5750610e8e565b601f016020900490600052602060002090810190610e8d9190610e91565b5b50565b5b80821115610eaa576000816000905550600101610e92565b5090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ed982610eae565b9050919050565b610ee981610ece565b82525050565b6000602082019050610f046000830184610ee0565b92915050565b6000604051905090565b600080fd5b600080fd5b610f2781610ece565b8114610f3257600080fd5b50565b600081359050610f4481610f1e565b92915050565b600060208284031215610f6057610f5f610f14565b5b6000610f6e84828501610f35565b91505092915050565b6000819050919050565b610f8a81610f77565b8114610f9557600080fd5b50565b600081359050610fa781610f81565b92915050565b600060208284031215610fc357610fc2610f14565b5b6000610fd184828501610f98565b91505092915050565b610fe381610f77565b82525050565b6000602082019050610ffe6000830184610fda565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110578261100e565b810181811067ffffffffffffffff821117156110765761107561101f565b5b80604052505050565b6000611089610f0a565b9050611095828261104e565b919050565b600067ffffffffffffffff8211156110b5576110b461101f565b5b6110be8261100e565b9050602081019050919050565b82818337600083830152505050565b60006110ed6110e88461109a565b61107f565b90508281526020810184848401111561110957611108611009565b5b6111148482856110cb565b509392505050565b600082601f83011261113157611130611004565b5b81356111418482602086016110da565b91505092915050565b60008060006060848603121561116357611162610f14565b5b600061117186828701610f98565b935050602084013567ffffffffffffffff81111561119257611191610f19565b5b61119e8682870161111c565b925050604084013567ffffffffffffffff8111156111bf576111be610f19565b5b6111cb8682870161111c565b9150509250925092565b600067ffffffffffffffff8211156111f0576111ef61101f565b5b602082029050602081019050919050565b600080fd5b6000611219611214846111d5565b61107f565b9050808382526020820190506020840283018581111561123c5761123b611201565b5b835b8181101561128357803567ffffffffffffffff81111561126157611260611004565b5b80860161126e898261111c565b8552602085019450505060208101905061123e565b5050509392505050565b600082601f8301126112a2576112a1611004565b5b81356112b2848260208601611206565b91505092915050565b6000806000606084860312156112d4576112d3610f14565b5b600084013567ffffffffffffffff8111156112f2576112f1610f19565b5b6112fe8682870161111c565b935050602084013567ffffffffffffffff81111561131f5761131e610f19565b5b61132b8682870161128d565b925050604061133c86828701610f35565b9150509250925092565b600082825260208201905092915050565b7f546f6b656e206164647265737300000000000000000000000000000000000000600082015250565b600061138d600d83611346565b915061139882611357565b602082019050919050565b600060208201905081810360008301526113bc81611380565b9050919050565b7f4d696e206465706f736974000000000000000000000000000000000000000000600082015250565b60006113f9600b83611346565b9150611404826113c3565b602082019050919050565b60006020820190508181036000830152611428816113ec565b9050919050565b7f4a6f62206e6f7420666f756e6400000000000000000000000000000000000000600082015250565b6000611465600d83611346565b91506114708261142f565b602082019050919050565b6000602082019050818103600083015261149481611458565b9050919050565b600081519050919050565b60005b838110156114c45780820151818401526020810190506114a9565b838111156114d3576000848401525b50505050565b60006114e48261149b565b6114ee8185611346565b93506114fe8185602086016114a6565b6115078161100e565b840191505092915050565b60006060820190506115276000830186610fda565b818103602083015261153981856114d9565b9050818103604083015261154d81846114d9565b9050949350505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60006115b3602e83611346565b91506115be82611557565b604082019050919050565b600060208201905081810360008301526115e2816115a6565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b600061162561162061161b846115e9565b611600565b6115f3565b9050919050565b6116358161160a565b82525050565b6000602082019050611650600083018461162c565b92915050565b600060408201905061166b6000830185610ee0565b6116786020830184610ee0565b9392505050565b60008151905061168e81610f81565b92915050565b6000602082840312156116aa576116a9610f14565b5b60006116b88482850161167f565b91505092915050565b7f546f6b656e20616c6c6f77616e6365206e6f7420656e6f756768000000000000600082015250565b60006116f7601a83611346565b9150611702826116c1565b602082019050919050565b60006020820190508181036000830152611726816116ea565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061176782610f77565b915061177283610f77565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156117a7576117a661172d565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806117f957607f821691505b60208210810361180c5761180b6117b2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118747fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611837565b61187e8683611837565b95508019841693508086168417925050509392505050565b60006118b16118ac6118a784610f77565b611600565b610f77565b9050919050565b6000819050919050565b6118cb83611896565b6118df6118d7826118b8565b848454611844565b825550505050565b600090565b6118f46118e7565b6118ff8184846118c2565b505050565b5b81811015611923576119186000826118ec565b600181019050611905565b5050565b601f8211156119685761193981611812565b61194284611827565b81016020851015611951578190505b61196561195d85611827565b830182611904565b50505b505050565b600082821c905092915050565b600061198b6000198460080261196d565b1980831691505092915050565b60006119a4838361197a565b9150826002028217905092915050565b6119bd8261149b565b67ffffffffffffffff8111156119d6576119d561101f565b5b6119e082546117e1565b6119eb828285611927565b600060209050601f831160018114611a1e5760008415611a0c578287015190505b611a168582611998565b865550611a7e565b601f198416611a2c86611812565b60005b82811015611a5457848901518255600182019150602085019450602081019050611a2f565b86831015611a715784890151611a6d601f89168261197a565b8355505b6001600288020188555050505b505050505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000611ace8261149b565b611ad88185611ab2565b9350611ae88185602086016114a6565b611af18161100e565b840191505092915050565b6000611b088383611ac3565b905092915050565b6000602082019050919050565b6000611b2882611a86565b611b328185611a91565b935083602082028501611b4485611aa2565b8060005b85811015611b805784840389528151611b618582611afc565b9450611b6c83611b10565b925060208a01995050600181019050611b48565b50829750879550505050505092915050565b600060a082019050611ba76000830188610fda565b611bb46020830187610ee0565b611bc16040830186610ee0565b8181036060830152611bd381856114d9565b90508181036080830152611be78184611b1d565b90509695505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611c4f602683611346565b9150611c5a82611bf3565b604082019050919050565b60006020820190508181036000830152611c7e81611c42565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c657220616460008201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b6000611ce1603583611346565b9150611cec82611c85565b604082019050919050565b60006020820190508181036000830152611d1081611cd4565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e7460008201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b6000611d73603983611346565b9150611d7e82611d17565b604082019050919050565b60006020820190508181036000830152611da281611d66565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611ddf602083611346565b9150611dea82611da9565b602082019050919050565b60006020820190508181036000830152611e0e81611dd2565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e7460008201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b6000611e71603b83611346565b9150611e7c82611e15565b604082019050919050565b60006020820190508181036000830152611ea081611e64565b905091905056fea26469706673582212203930a855df191172367bcfee342242de53f1d5eca88173e8422548b7359f369b64736f6c634300080f0033",
}

// JobcreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use JobcreatorMetaData.ABI instead.
var JobcreatorABI = JobcreatorMetaData.ABI

// JobcreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobcreatorMetaData.Bin instead.
var JobcreatorBin = JobcreatorMetaData.Bin

// DeployJobcreator deploys a new Ethereum contract, binding an instance of Jobcreator to it.
func DeployJobcreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jobcreator, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobcreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// Jobcreator is an auto generated Go binding around an Ethereum contract.
type Jobcreator struct {
	JobcreatorCaller     // Read-only binding to the contract
	JobcreatorTransactor // Write-only binding to the contract
	JobcreatorFilterer   // Log filterer for contract events
}

// JobcreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobcreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobcreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobcreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobcreatorSession struct {
	Contract     *Jobcreator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobcreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobcreatorCallerSession struct {
	Contract *JobcreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobcreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobcreatorTransactorSession struct {
	Contract     *JobcreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobcreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobcreatorRaw struct {
	Contract *Jobcreator // Generic contract binding to access the raw methods on
}

// JobcreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobcreatorCallerRaw struct {
	Contract *JobcreatorCaller // Generic read-only contract binding to access the raw methods on
}

// JobcreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobcreatorTransactorRaw struct {
	Contract *JobcreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobcreator creates a new instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreator(address common.Address, backend bind.ContractBackend) (*Jobcreator, error) {
	contract, err := bindJobcreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// NewJobcreatorCaller creates a new read-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorCaller(address common.Address, caller bind.ContractCaller) (*JobcreatorCaller, error) {
	contract, err := bindJobcreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorCaller{contract: contract}, nil
}

// NewJobcreatorTransactor creates a new write-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*JobcreatorTransactor, error) {
	contract, err := bindJobcreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorTransactor{contract: contract}, nil
}

// NewJobcreatorFilterer creates a new log filterer instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*JobcreatorFilterer, error) {
	contract, err := bindJobcreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobcreatorFilterer{contract: contract}, nil
}

// bindJobcreator binds a generic wrapper to an already deployed contract.
func bindJobcreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.JobcreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) GetRequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getRequiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) NextJobID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "nextJobID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "requiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactor) RunJob(opts *bind.TransactOpts, module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "runJob", module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactor) SetRequiredDeposit(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setRequiredDeposit", cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorTransactor) SubmitResults(opts *bind.TransactOpts, id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "submitResults", id, dealId, dataId)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorSession) SubmitResults(id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.Contract.SubmitResults(&_Jobcreator.TransactOpts, id, dealId, dataId)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorTransactorSession) SubmitResults(id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.Contract.SubmitResults(&_Jobcreator.TransactOpts, id, dealId, dataId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// JobcreatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Jobcreator contract.
type JobcreatorInitializedIterator struct {
	Event *JobcreatorInitialized // Event containing the contract specifics and raw log

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
func (it *JobcreatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorInitialized)
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
		it.Event = new(JobcreatorInitialized)
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
func (it *JobcreatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorInitialized represents a Initialized event raised by the Jobcreator contract.
type JobcreatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*JobcreatorInitializedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JobcreatorInitializedIterator{contract: _Jobcreator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JobcreatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorInitialized)
				if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Jobcreator *JobcreatorFilterer) ParseInitialized(log types.Log) (*JobcreatorInitialized, error) {
	event := new(JobcreatorInitialized)
	if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorJobAddedIterator is returned from FilterJobAdded and is used to iterate over the raw logs and unpacked data for JobAdded events raised by the Jobcreator contract.
type JobcreatorJobAddedIterator struct {
	Event *JobcreatorJobAdded // Event containing the contract specifics and raw log

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
func (it *JobcreatorJobAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorJobAdded)
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
		it.Event = new(JobcreatorJobAdded)
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
func (it *JobcreatorJobAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorJobAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorJobAdded represents a JobAdded event raised by the Jobcreator contract.
type JobcreatorJobAdded struct {
	Id              *big.Int
	CallingContract common.Address
	Payee           common.Address
	Module          string
	Inputs          []string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterJobAdded is a free log retrieval operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) FilterJobAdded(opts *bind.FilterOpts) (*JobcreatorJobAddedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return &JobcreatorJobAddedIterator{contract: _Jobcreator.contract, event: "JobAdded", logs: logs, sub: sub}, nil
}

// WatchJobAdded is a free log subscription operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) WatchJobAdded(opts *bind.WatchOpts, sink chan<- *JobcreatorJobAdded) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorJobAdded)
				if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
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

// ParseJobAdded is a log parse operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) ParseJobAdded(log types.Log) (*JobcreatorJobAdded, error) {
	event := new(JobcreatorJobAdded)
	if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Jobcreator contract.
type JobcreatorOwnershipTransferredIterator struct {
	Event *JobcreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JobcreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorOwnershipTransferred)
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
		it.Event = new(JobcreatorOwnershipTransferred)
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
func (it *JobcreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorOwnershipTransferred represents a OwnershipTransferred event raised by the Jobcreator contract.
type JobcreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JobcreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JobcreatorOwnershipTransferredIterator{contract: _Jobcreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JobcreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorOwnershipTransferred)
				if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Jobcreator *JobcreatorFilterer) ParseOwnershipTransferred(log types.Log) (*JobcreatorOwnershipTransferred, error) {
	event := new(JobcreatorOwnershipTransferred)
	if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
