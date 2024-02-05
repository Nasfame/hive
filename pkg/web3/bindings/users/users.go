// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package users

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

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress common.Address
	MetadataCID string
	Url         string
	Roles       []uint8
}

// UsersMetaData contains all meta data concerning the Users contract.
var UsersMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataCID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612113806200010e6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a15dcc8a11610066578063a15dcc8a146100fa578063a7f96f061461012a578063aeb5ec0114610146578063ebbbca2714610162578063f2fde38b1461019257610093565b80636f77926b14610098578063715018a6146100c85780638129fc1c146100d25780638da5cb5b146100dc575b600080fd5b6100b260048036038101906100ad919061123d565b6101ae565b6040516100bf91906114a9565b60405180910390f35b6100d061040e565b005b6100da610422565b005b6100e461055d565b6040516100f191906114da565b60405180910390f35b610114600480360381019061010f919061151a565b610586565b60405161012191906115f6565b60405180910390f35b610144600480360381019061013f919061151a565b61064b565b005b610160600480360381019061015b919061151a565b610967565b005b61017c60048036038101906101779190611815565b610b6d565b60405161018991906114a9565b60405180910390f35b6101ac60048036038101906101a7919061123d565b610c8e565b005b6101b66110ba565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610266906118eb565b80601f0160208091040260200160405190810160405280929190818152602001828054610292906118eb565b80156102df5780601f106102b4576101008083540402835291602001916102df565b820191906000526020600020905b8154815290600101906020018083116102c257829003601f168201915b505050505081526020016002820180546102f8906118eb565b80601f0160208091040260200160405190810160405280929190818152602001828054610324906118eb565b80156103715780601f1061034657610100808354040283529160200191610371565b820191906000526020600020905b81548152906001019060200180831161035457829003601f168201915b50505050508152602001600382018054806020026020016040519081016040528092919081815260200182805480156103fe57602002820191906000526020600020906000905b82829054906101000a900460ff1660038111156103d8576103d761133e565b5b815260200190600101906020826000010492830192600103820291508084116103b85790505b5050505050815250509050919050565b610416610d11565b6104206000610d8f565b565b60008060159054906101000a900460ff1615905080801561045557506001600060149054906101000a900460ff1660ff16105b80610484575061046430610e53565b15801561048357506001600060149054906101000a900460ff1660ff16145b5b6104c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ba9061199f565b60405180910390fd5b6001600060146101000a81548160ff021916908360ff1602179055508015610501576001600060156101000a81548160ff0219169083151502179055505b801561055a5760008060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516105519190611a11565b60405180910390a15b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606002600083600381111561059f5761059e61133e565b5b60038111156105b1576105b061133e565b5b815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561063f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116105f5575b50505050509050919050565b600073ffffffffffffffffffffffffffffffffffffffff16600160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361071c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071390611a78565b60405180910390fd5b6000806107298332610e76565b915091508061076d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076490611ae4565b60405180910390fd5b60008290505b60016002600086600381111561078c5761078b61133e565b5b600381111561079e5761079d61133e565b5b8152602001908152602001600020805490506107ba9190611b3d565b8110156108e557600260008560038111156107d8576107d761133e565b5b60038111156107ea576107e961133e565b5b81526020019081526020016000206001826108059190611b71565b8154811061081657610815611bc7565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008660038111156108585761085761133e565b5b600381111561086a5761086961133e565b5b8152602001908152602001600020828154811061088a57610889611bc7565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806108dd90611bf6565b915050610773565b50600260008460038111156108fd576108fc61133e565b5b600381111561090f5761090e61133e565b5b815260200190815260200160002080548061092d5761092c611c3e565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b600073ffffffffffffffffffffffffffffffffffffffff16600160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610a38576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2f90611a78565b60405180910390fd5b6000610a448232610e76565b9150508015610a88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7f90611cdf565b60405180910390fd5b610a928232610f96565b610ad1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac890611d4b565b60405180910390fd5b60026000836003811115610ae857610ae761133e565b5b6003811115610afa57610af961133e565b5b8152602001908152602001600020329080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610b756110ba565b600060405180608001604052803273ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815260200184815250905080600160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019081610c4c9190611f0d565b506040820151816002019081610c629190611f0d565b506060820151816003019080519060200190610c7f9291906110f8565b50905050809150509392505050565b610c96610d11565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cfc90612051565b60405180910390fd5b610d0e81610d8f565b50565b610d196110b2565b73ffffffffffffffffffffffffffffffffffffffff16610d3761055d565b73ffffffffffffffffffffffffffffffffffffffff1614610d8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d84906120bd565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008060008060005b60026000886003811115610e9657610e9561133e565b5b6003811115610ea857610ea761133e565b5b815260200190815260200160002080549050811015610f86578573ffffffffffffffffffffffffffffffffffffffff1660026000896003811115610eef57610eee61133e565b5b6003811115610f0157610f0061133e565b5b81526020019081526020016000208281548110610f2157610f20611bc7565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610f735780925060019150610f86565b8080610f7e90611bf6565b915050610e7f565b5081819350935050509250929050565b6000806000905060005b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301805490508110156110a75784600381111561100057610fff61133e565b5b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301828154811061105457611053611bc7565b5b90600052602060002090602091828204019190069054906101000a900460ff1660038111156110865761108561133e565b5b0361109457600191506110a7565b808061109f90611bf6565b915050610fa0565b508091505092915050565b600033905090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016060815260200160608152602001606081525090565b82805482825590600052602060002090601f0160209004810192821561119d5791602002820160005b8382111561116e57835183826101000a81548160ff0219169083600381111561114d5761114c61133e565b5b02179055509260200192600101602081600001049283019260010302611121565b801561119b5782816101000a81549060ff021916905560010160208160000104928301926001030261116e565b505b5090506111aa91906111ae565b5090565b5b808211156111c75760008160009055506001016111af565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061120a826111df565b9050919050565b61121a816111ff565b811461122557600080fd5b50565b60008135905061123781611211565b92915050565b600060208284031215611253576112526111d5565b5b600061126184828501611228565b91505092915050565b611273816111ff565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156112b3578082015181840152602081019050611298565b838111156112c2576000848401525b50505050565b6000601f19601f8301169050919050565b60006112e482611279565b6112ee8185611284565b93506112fe818560208601611295565b611307816112c8565b840191505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004811061137e5761137d61133e565b5b50565b600081905061138f8261136d565b919050565b600061139f82611381565b9050919050565b6113af81611394565b82525050565b60006113c183836113a6565b60208301905092915050565b6000602082019050919050565b60006113e582611312565b6113ef818561131d565b93506113fa8361132e565b8060005b8381101561142b57815161141288826113b5565b975061141d836113cd565b9250506001810190506113fe565b5085935050505092915050565b6000608083016000830151611450600086018261126a565b506020830151848203602086015261146882826112d9565b9150506040830151848203604086015261148282826112d9565b9150506060830151848203606086015261149c82826113da565b9150508091505092915050565b600060208201905081810360008301526114c38184611438565b905092915050565b6114d4816111ff565b82525050565b60006020820190506114ef60008301846114cb565b92915050565b6004811061150257600080fd5b50565b600081359050611514816114f5565b92915050565b6000602082840312156115305761152f6111d5565b5b600061153e84828501611505565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600061157f838361126a565b60208301905092915050565b6000602082019050919050565b60006115a382611547565b6115ad8185611552565b93506115b883611563565b8060005b838110156115e95781516115d08882611573565b97506115db8361158b565b9250506001810190506115bc565b5085935050505092915050565b600060208201905081810360008301526116108184611598565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61165a826112c8565b810181811067ffffffffffffffff8211171561167957611678611622565b5b80604052505050565b600061168c6111cb565b90506116988282611651565b919050565b600067ffffffffffffffff8211156116b8576116b7611622565b5b6116c1826112c8565b9050602081019050919050565b82818337600083830152505050565b60006116f06116eb8461169d565b611682565b90508281526020810184848401111561170c5761170b61161d565b5b6117178482856116ce565b509392505050565b600082601f83011261173457611733611618565b5b81356117448482602086016116dd565b91505092915050565b600067ffffffffffffffff82111561176857611767611622565b5b602082029050602081019050919050565b600080fd5b600061179161178c8461174d565b611682565b905080838252602082019050602084028301858111156117b4576117b3611779565b5b835b818110156117dd57806117c98882611505565b8452602084019350506020810190506117b6565b5050509392505050565b600082601f8301126117fc576117fb611618565b5b813561180c84826020860161177e565b91505092915050565b60008060006060848603121561182e5761182d6111d5565b5b600084013567ffffffffffffffff81111561184c5761184b6111da565b5b6118588682870161171f565b935050602084013567ffffffffffffffff811115611879576118786111da565b5b6118858682870161171f565b925050604084013567ffffffffffffffff8111156118a6576118a56111da565b5b6118b2868287016117e7565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061190357607f821691505b602082108103611916576119156118bc565b5b50919050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000611989602e8361191c565b91506119948261192d565b604082019050919050565b600060208201905081810360008301526119b88161197c565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b60006119fb6119f66119f1846119bf565b6119d6565b6119c9565b9050919050565b611a0b816119e0565b82525050565b6000602082019050611a266000830184611a02565b92915050565b7f55736572206d7573742065786973740000000000000000000000000000000000600082015250565b6000611a62600f8361191c565b9150611a6d82611a2c565b602082019050919050565b60006020820190508181036000830152611a9181611a55565b9050919050565b7f55736572206973206e6f742070617274206f662074686174206c697374000000600082015250565b6000611ace601d8361191c565b9150611ad982611a98565b602082019050919050565b60006020820190508181036000830152611afd81611ac1565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b4882611b04565b9150611b5383611b04565b925082821015611b6657611b65611b0e565b5b828203905092915050565b6000611b7c82611b04565b9150611b8783611b04565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611bbc57611bbb611b0e565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611c0182611b04565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c3357611c32611b0e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f5573657220697320616c72656164792070617274206f662074686174206c697360008201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b6000611cc960218361191c565b9150611cd482611c6d565b604082019050919050565b60006020820190508181036000830152611cf881611cbc565b9050919050565b7f55736572206d7573742068617665207468617420726f6c650000000000000000600082015250565b6000611d3560188361191c565b9150611d4082611cff565b602082019050919050565b60006020820190508181036000830152611d6481611d28565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611dcd7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611d90565b611dd78683611d90565b95508019841693508086168417925050509392505050565b6000611e0a611e05611e0084611b04565b6119d6565b611b04565b9050919050565b6000819050919050565b611e2483611def565b611e38611e3082611e11565b848454611d9d565b825550505050565b600090565b611e4d611e40565b611e58818484611e1b565b505050565b5b81811015611e7c57611e71600082611e45565b600181019050611e5e565b5050565b601f821115611ec157611e9281611d6b565b611e9b84611d80565b81016020851015611eaa578190505b611ebe611eb685611d80565b830182611e5d565b50505b505050565b600082821c905092915050565b6000611ee460001984600802611ec6565b1980831691505092915050565b6000611efd8383611ed3565b9150826002028217905092915050565b611f1682611279565b67ffffffffffffffff811115611f2f57611f2e611622565b5b611f3982546118eb565b611f44828285611e80565b600060209050601f831160018114611f775760008415611f65578287015190505b611f6f8582611ef1565b865550611fd7565b601f198416611f8586611d6b565b60005b82811015611fad57848901518255600182019150602085019450602081019050611f88565b86831015611fca5784890151611fc6601f891682611ed3565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061203b60268361191c565b915061204682611fdf565b604082019050919050565b6000602082019050818103600083015261206a8161202e565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006120a760208361191c565b91506120b282612071565b602082019050919050565b600060208201905081810360008301526120d68161209a565b905091905056fea264697066735822122084322f2f96df7a6960e6f72f1dfaf8a268f050368bf4a5e6800ccd9ba23a984464736f6c634300080f0033",
}

// UsersABI is the input ABI used to generate the binding from.
// Deprecated: Use UsersMetaData.ABI instead.
var UsersABI = UsersMetaData.ABI

// UsersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UsersMetaData.Bin instead.
var UsersBin = UsersMetaData.Bin

// DeployUsers deploys a new Ethereum contract, binding an instance of Users to it.
func DeployUsers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Users, error) {
	parsed, err := UsersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UsersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// Users is an auto generated Go binding around an Ethereum contract.
type Users struct {
	UsersCaller     // Read-only binding to the contract
	UsersTransactor // Write-only binding to the contract
	UsersFilterer   // Log filterer for contract events
}

// UsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsersSession struct {
	Contract     *Users            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsersCallerSession struct {
	Contract *UsersCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsersTransactorSession struct {
	Contract     *UsersTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsersRaw struct {
	Contract *Users // Generic contract binding to access the raw methods on
}

// UsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsersCallerRaw struct {
	Contract *UsersCaller // Generic read-only contract binding to access the raw methods on
}

// UsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsersTransactorRaw struct {
	Contract *UsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsers creates a new instance of Users, bound to a specific deployed contract.
func NewUsers(address common.Address, backend bind.ContractBackend) (*Users, error) {
	contract, err := bindUsers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Users{UsersCaller: UsersCaller{contract: contract}, UsersTransactor: UsersTransactor{contract: contract}, UsersFilterer: UsersFilterer{contract: contract}}, nil
}

// NewUsersCaller creates a new read-only instance of Users, bound to a specific deployed contract.
func NewUsersCaller(address common.Address, caller bind.ContractCaller) (*UsersCaller, error) {
	contract, err := bindUsers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsersCaller{contract: contract}, nil
}

// NewUsersTransactor creates a new write-only instance of Users, bound to a specific deployed contract.
func NewUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*UsersTransactor, error) {
	contract, err := bindUsers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsersTransactor{contract: contract}, nil
}

// NewUsersFilterer creates a new log filterer instance of Users, bound to a specific deployed contract.
func NewUsersFilterer(address common.Address, filterer bind.ContractFilterer) (*UsersFilterer, error) {
	contract, err := bindUsers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsersFilterer{contract: contract}, nil
}

// bindUsers binds a generic wrapper to an already deployed contract.
func bindUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UsersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Users.Contract.UsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.UsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Users *UsersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Users.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Users *UsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Users *UsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Users.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,string,string,uint8[]))
func (_Users *UsersCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "getUser", userAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,string,string,uint8[]))
func (_Users *UsersSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Users.Contract.GetUser(&_Users.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,string,string,uint8[]))
func (_Users *UsersCallerSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Users.Contract.GetUser(&_Users.CallOpts, userAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Users *UsersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Users *UsersSession) Owner() (common.Address, error) {
	return _Users.Contract.Owner(&_Users.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Users *UsersCallerSession) Owner() (common.Address, error) {
	return _Users.Contract.Owner(&_Users.CallOpts)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Users *UsersCaller) ShowUsersInList(opts *bind.CallOpts, serviceType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Users.contract.Call(opts, &out, "showUsersInList", serviceType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Users *UsersSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Users.Contract.ShowUsersInList(&_Users.CallOpts, serviceType)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Users *UsersCallerSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Users.Contract.ShowUsersInList(&_Users.CallOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Users *UsersTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Users *UsersSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.AddUserToList(&_Users.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Users *UsersTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.AddUserToList(&_Users.TransactOpts, serviceType)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Users *UsersTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Users *UsersSession) Initialize() (*types.Transaction, error) {
	return _Users.Contract.Initialize(&_Users.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Users *UsersTransactorSession) Initialize() (*types.Transaction, error) {
	return _Users.Contract.Initialize(&_Users.TransactOpts)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Users *UsersTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Users *UsersSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.RemoveUserFromList(&_Users.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Users *UsersTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Users.Contract.RemoveUserFromList(&_Users.TransactOpts, serviceType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Users *UsersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Users *UsersSession) RenounceOwnership() (*types.Transaction, error) {
	return _Users.Contract.RenounceOwnership(&_Users.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Users *UsersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Users.Contract.RenounceOwnership(&_Users.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Users *UsersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Users *UsersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Users.Contract.TransferOwnership(&_Users.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Users *UsersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Users.Contract.TransferOwnership(&_Users.TransactOpts, newOwner)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xebbbca27.
//
// Solidity: function updateUser(string metadataCID, string url, uint8[] roles) returns((address,string,string,uint8[]))
func (_Users *UsersTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID string, url string, roles []uint8) (*types.Transaction, error) {
	return _Users.contract.Transact(opts, "updateUser", metadataCID, url, roles)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xebbbca27.
//
// Solidity: function updateUser(string metadataCID, string url, uint8[] roles) returns((address,string,string,uint8[]))
func (_Users *UsersSession) UpdateUser(metadataCID string, url string, roles []uint8) (*types.Transaction, error) {
	return _Users.Contract.UpdateUser(&_Users.TransactOpts, metadataCID, url, roles)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xebbbca27.
//
// Solidity: function updateUser(string metadataCID, string url, uint8[] roles) returns((address,string,string,uint8[]))
func (_Users *UsersTransactorSession) UpdateUser(metadataCID string, url string, roles []uint8) (*types.Transaction, error) {
	return _Users.Contract.UpdateUser(&_Users.TransactOpts, metadataCID, url, roles)
}

// UsersInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Users contract.
type UsersInitializedIterator struct {
	Event *UsersInitialized // Event containing the contract specifics and raw log

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
func (it *UsersInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersInitialized)
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
		it.Event = new(UsersInitialized)
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
func (it *UsersInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersInitialized represents a Initialized event raised by the Users contract.
type UsersInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Users *UsersFilterer) FilterInitialized(opts *bind.FilterOpts) (*UsersInitializedIterator, error) {

	logs, sub, err := _Users.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UsersInitializedIterator{contract: _Users.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Users *UsersFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UsersInitialized) (event.Subscription, error) {

	logs, sub, err := _Users.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersInitialized)
				if err := _Users.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Users *UsersFilterer) ParseInitialized(log types.Log) (*UsersInitialized, error) {
	event := new(UsersInitialized)
	if err := _Users.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UsersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Users contract.
type UsersOwnershipTransferredIterator struct {
	Event *UsersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UsersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UsersOwnershipTransferred)
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
		it.Event = new(UsersOwnershipTransferred)
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
func (it *UsersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UsersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UsersOwnershipTransferred represents a OwnershipTransferred event raised by the Users contract.
type UsersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Users *UsersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UsersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Users.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UsersOwnershipTransferredIterator{contract: _Users.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Users *UsersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UsersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Users.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UsersOwnershipTransferred)
				if err := _Users.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Users *UsersFilterer) ParseOwnershipTransferred(log types.Log) (*UsersOwnershipTransferred, error) {
	event := new(UsersOwnershipTransferred)
	if err := _Users.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
