// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// SharedStructsAgreement is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsAgreement struct {
	State                    uint8
	ResourceProviderAgreedAt *big.Int
	JobCreatorAgreedAt       *big.Int
	DealCreatedAt            *big.Int
	DealAgreedAt             *big.Int
	ResultsSubmittedAt       *big.Int
	ResultsAcceptedAt        *big.Int
	ResultsCheckedAt         *big.Int
	MediationAcceptedAt      *big.Int
	MediationRejectedAt      *big.Int
	TimeoutAgreeAt           *big.Int
	TimeoutSubmitResultsAt   *big.Int
	TimeoutJudgeResultsAt    *big.Int
	TimeoutMediateResultsAt  *big.Int
}

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

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	DealId           string
	ResultsId        string
	DataId           string
	InstructionCount *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealCreatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutAgreeAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"mediators\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.DealMembers\",\"name\":\"members\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"agree\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"submitResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"judgeResults\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealTimeout\",\"name\":\"mediateResults\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.DealTimeouts\",\"name\":\"timeouts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.DealPricing\",\"name\":\"pricing\",\"type\":\"tuple\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"resultsId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutAgree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff0219169083151502179055503480156200004657600080fd5b50620000676200005b6200006d60201b60201c565b6200007560201b60201c565b62000139565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61510080620001496000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80638224ce5f116100f9578063cdd82d1d11610097578063e850be3711610071578063e850be37146104dc578063ec95b967146104f8578063f2fde38b14610528578063f3d3d44814610544576101a9565b8063cdd82d1d14610460578063e707918014610490578063e7b957d1146104c0576101a9565b8063a4702958116100d3578063a4702958146103d8578063a6370b0e146103e2578063b050e74b14610412578063c57380a214610442576101a9565b80638224ce5f1461036e578063824518aa1461039e5780638da5cb5b146103ba576101a9565b8063498cc70d1161016657806373db5c6a1161014057806373db5c6a146102fc578063795f9abf1461032c57806380ffdfe0146103485780638129fc1c14610364576101a9565b8063498cc70d146102a6578063511a9f68146102d6578063715018a6146102f2576101a9565b806311d5af33146101ae5780632244ad2b146101de578063297f9e551461020e5780633955548e1461022a5780633c4135da1461025a57806346834d1e1461028a575b600080fd5b6101c860048036038101906101c391906130fd565b610560565b6040516101d59190613285565b60405180910390f35b6101f860048036038101906101f391906133dc565b610678565b6040516102059190613440565b60405180910390f35b610228600480360381019061022391906133dc565b610691565b005b610244600480360381019061023f9190613491565b610718565b60405161025191906135cc565b60405180910390f35b610274600480360381019061026f91906133dc565b610a25565b6040516102819190613785565b60405180910390f35b6102a4600480360381019061029f91906133dc565b610bff565b005b6102c060048036038101906102bb91906133dc565b610c86565b6040516102cd91906135cc565b60405180910390f35b6102f060048036038101906102eb91906133dc565b610e7e565b005b6102fa610f05565b005b610316600480360381019061031191906133dc565b610f19565b60405161032391906137b0565b60405180910390f35b610346600480360381019061034191906133dc565b610f73565b005b610362600480360381019061035d91906133dc565b610ffa565b005b61036c611081565b005b610388600480360381019061038391906133dc565b6111ba565b60405161039591906137b0565b60405180910390f35b6103b860048036038101906103b391906133dc565b6111fb565b005b6103c2611282565b6040516103cf91906137da565b60405180910390f35b6103e06112ab565b005b6103fc60048036038101906103f79190613a9c565b6112d0565b6040516104099190613da4565b60405180910390f35b61042c60048036038101906104279190613deb565b6119dd565b6040516104399190613440565b60405180910390f35b61044a611a79565b60405161045791906137da565b60405180910390f35b61047a600480360381019061047591906133dc565b611aa3565b6040516104879190613785565b60405180910390f35b6104aa60048036038101906104a591906133dc565b611b99565b6040516104b79190613da4565b60405180910390f35b6104da60048036038101906104d591906133dc565b611ef7565b005b6104f660048036038101906104f191906133dc565b611f7e565b005b610512600480360381019061050d91906133dc565b612005565b60405161051f9190613785565b60405180910390f35b610542600480360381019061053d91906130fd565b6121df565b005b61055e600480360381019061055991906130fd565b612262565b005b6060600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561066d5783829060005260206000200180546105e090613e76565b80601f016020809104026020016040519081016040528092919081815260200182805461060c90613e76565b80156106595780601f1061062e57610100808354040283529160200191610659565b820191906000526020600020905b81548152906001019060200180831161063c57829003601f168201915b5050505050815260200190600101906105c1565b505050509050919050565b60008061068483611b99565b6000015151119050919050565b61069961236c565b506106a58160026119dd565b6106e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106db90613f04565b60405180910390fd5b426005826040516106f59190613f60565b90815260200160405180910390206006018190555061071581600361249c565b50565b610720612e15565b61072861236c565b506107348560016119dd565b610773576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076a90613fc3565b60405180910390fd5b426005866040516107849190613f60565b9081526020016040518091039020600501819055506107a485600261249c565b6040518060800160405280868152602001858152602001848152602001838152506006866040516107d59190613f60565b908152602001604051809103902060008201518160000190816107f8919061418f565b50602082015181600101908161080e919061418f565b506040820151816002019081610824919061418f565b50606082015181600301559050506006856040516108429190613f60565b908152602001604051809103902060405180608001604052908160008201805461086b90613e76565b80601f016020809104026020016040519081016040528092919081815260200182805461089790613e76565b80156108e45780601f106108b9576101008083540402835291602001916108e4565b820191906000526020600020905b8154815290600101906020018083116108c757829003601f168201915b505050505081526020016001820180546108fd90613e76565b80601f016020809104026020016040519081016040528092919081815260200182805461092990613e76565b80156109765780601f1061094b57610100808354040283529160200191610976565b820191906000526020600020905b81548152906001019060200180831161095957829003601f168201915b5050505050815260200160028201805461098f90613e76565b80601f01602080910402602001604051908101604052809291908181526020018280546109bb90613e76565b8015610a085780601f106109dd57610100808354040283529160200191610a08565b820191906000526020600020905b8154815290600101906020018083116109eb57829003601f168201915b505050505081526020016003820154815250509050949350505050565b610a2d612e3d565b610a3561236c565b50610a3f82610678565b610a7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a75906142ad565b60405180910390fd5b6000600583604051610a909190613f60565b90815260200160405180910390206002015414610ae2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad990614319565b60405180910390fd5b42600583604051610af39190613f60565b908152602001604051809103902060020181905550610b1182612522565b600582604051610b219190613f60565b9081526020016040518091039020604051806101c00160405290816000820160009054906101000a900460ff16600a811115610b6057610b5f6135ee565b5b600a811115610b7257610b716135ee565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b610c0761236c565b50610c138160026119dd565b610c52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4990613f04565b60405180910390fd5b42600582604051610c639190613f60565b908152602001604051809103902060070181905550610c8381600461249c565b50565b610c8e612e15565b600682604051610c9e9190613f60565b9081526020016040518091039020604051806080016040529081600082018054610cc790613e76565b80601f0160208091040260200160405190810160405280929190818152602001828054610cf390613e76565b8015610d405780601f10610d1557610100808354040283529160200191610d40565b820191906000526020600020905b815481529060010190602001808311610d2357829003601f168201915b50505050508152602001600182018054610d5990613e76565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8590613e76565b8015610dd25780601f10610da757610100808354040283529160200191610dd2565b820191906000526020600020905b815481529060010190602001808311610db557829003601f168201915b50505050508152602001600282018054610deb90613e76565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1790613e76565b8015610e645780601f10610e3957610100808354040283529160200191610e64565b820191906000526020600020905b815481529060010190602001808311610e4757829003601f168201915b505050505081526020016003820154815250509050919050565b610e8661236c565b50610e928160016119dd565b610ed1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec890613fc3565b60405180910390fd5b42600582604051610ee29190613f60565b9081526020016040518091039020600b0181905550610f0281600861249c565b50565b610f0d6125db565b610f176000612659565b565b6000600682604051610f2b9190613f60565b908152602001604051809103902060030154600383604051610f4d9190613f60565b9081526020016040518091039020600d0160000154610f6c9190614368565b9050919050565b610f7b61236c565b50610f878160006119dd565b610fc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbd9061440e565b60405180910390fd5b42600582604051610fd79190613f60565b9081526020016040518091039020600a0181905550610ff781600761249c565b50565b61100261236c565b5061100e8160046119dd565b61104d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110449061447a565b60405180910390fd5b4260058260405161105e9190613f60565b90815260200160405180910390206009018190555061107e81600661249c565b50565b6000600160169054906101000a900460ff161590508080156110b4575060018060159054906101000a900460ff1660ff16105b806110e257506110c33061271d565b1580156110e1575060018060159054906101000a900460ff1660ff16145b5b611121576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111189061450c565b60405180910390fd5b60018060156101000a81548160ff021916908360ff160217905550801561115d5760018060166101000a81548160ff0219169083151502179055505b80156111b7576000600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516111ae9190614574565b60405180910390a15b50565b60006111c582610f19565b6003836040516111d59190613f60565b9081526020016040518091039020600d01600201546111f49190614368565b9050919050565b61120361236c565b5061120f8160046119dd565b61124e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112459061447a565b60405180910390fd5b4260058260405161125f9190613f60565b90815260200160405180910390206008018190555061127f81600561249c565b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6112b36125db565b6000600160146101000a81548160ff021916908315150217905550565b6112d8612ebe565b6112e061236c565b506112ec8560006119dd565b61132b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113229061440e565b60405180910390fd5b61133484612740565b61133d836128e4565b61134685610678565b1561138857600061135686611b99565b905061136681602001518661297d565b611374816040015185612bfd565b611382816060015184612c49565b50611684565b6040518060800160405280868152602001858152602001848152602001838152506003866040516113b99190613f60565b908152602001604051809103902060008201518160000190816113dc919061418f565b5060208201518160010160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190805190602001906114d7929190612ef8565b505050604082015181600501600082015181600001600082015181600001556020820151816001015550506020820151816002016000820151816000015560208201518160010155505060408201518160040160008201518160000155602082015181600101555050606082015181600601600082015181600001556020820151816001015550505050606082015181600d0160008201518160000155602082015181600101556040820151816002015560608201518160030155505090505060046000856040015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208590806001815401808255809150506001900390600052602060002001600090919091909150908161160c919061418f565b5060046000856020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085908060018154018082558091505060019003906000526020600020016000909190919091509081611682919061418f565b505b6003856040516116949190613f60565b90815260200160405180910390206040518060800160405290816000820180546116bd90613e76565b80601f01602080910402602001604051908101604052809291908181526020018280546116e990613e76565b80156117365780601f1061170b57610100808354040283529160200191611736565b820191906000526020600020905b81548152906001019060200180831161171957829003601f168201915b50505050508152602001600182016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382018054806020026020016040519081016040528092919081815260200182805480156118d657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161188c575b5050505050815250508152602001600582016040518060800160405290816000820160405180604001604052908160008201548152602001600182015481525050815260200160028201604051806040016040529081600082015481526020016001820154815250508152602001600482016040518060400160405290816000820154815260200160018201548152505081526020016006820160405180604001604052908160008201548152602001600182015481525050815250508152602001600d8201604051806080016040529081600082015481526020016001820154815260200160028201548152602001600382015481525050815250509050949350505050565b60006119e883610678565b611a1b576000600a811115611a00576119ff6135ee565b5b82600a811115611a1357611a126135ee565b5b149050611a73565b81600a811115611a2e57611a2d6135ee565b5b600584604051611a3e9190613f60565b908152602001604051809103902060000160009054906101000a900460ff16600a811115611a6f57611a6e6135ee565b5b1490505b92915050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611aab612e3d565b600582604051611abb9190613f60565b9081526020016040518091039020604051806101c00160405290816000820160009054906101000a900460ff16600a811115611afa57611af96135ee565b5b600a811115611b0c57611b0b6135ee565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b611ba1612ebe565b600382604051611bb19190613f60565b9081526020016040518091039020604051806080016040529081600082018054611bda90613e76565b80601f0160208091040260200160405190810160405280929190818152602001828054611c0690613e76565b8015611c535780601f10611c2857610100808354040283529160200191611c53565b820191906000526020600020905b815481529060010190602001808311611c3657829003601f168201915b50505050508152602001600182016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201805480602002602001604051908101604052809291908181526020018280548015611df357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611da9575b5050505050815250508152602001600582016040518060800160405290816000820160405180604001604052908160008201548152602001600182015481525050815260200160028201604051806040016040529081600082015481526020016001820154815250508152602001600482016040518060400160405290816000820154815260200160018201548152505081526020016006820160405180604001604052908160008201548152602001600182015481525050815250508152602001600d8201604051806080016040529081600082015481526020016001820154815260200160028201548152602001600382015481525050815250509050919050565b611eff61236c565b50611f0b8160046119dd565b611f4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f419061447a565b60405180910390fd5b42600582604051611f5b9190613f60565b9081526020016040518091039020600d0181905550611f7b81600a61249c565b50565b611f8661236c565b50611f928160026119dd565b611fd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fc890613f04565b60405180910390fd5b42600582604051611fe29190613f60565b9081526020016040518091039020600c018190555061200281600961249c565b50565b61200d612e3d565b61201561236c565b5061201f82610678565b61205e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612055906142ad565b60405180910390fd5b60006005836040516120709190613f60565b908152602001604051809103902060010154146120c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b9906145db565b60405180910390fd5b426005836040516120d39190613f60565b9081526020016040518091039020600101819055506120f182612522565b6005826040516121019190613f60565b9081526020016040518091039020604051806101c00160405290816000820160009054906101000a900460ff16600a8111156121405761213f6135ee565b5b600a811115612152576121516135ee565b5b8152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b8201548152602001600c8201548152602001600d820154815250509050919050565b6121e76125db565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612256576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161224d9061466d565b60405180910390fd5b61225f81612659565b50565b61226a6125db565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036122d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122d0906146ff565b60405180910390fd5b600160149054906101000a900460ff16612328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161231f90614791565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036123fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123f5906146ff565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661243f612d75565b73ffffffffffffffffffffffffffffffffffffffff1614612495576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161248c90614823565b60405180910390fd5b6001905090565b806005836040516124ad9190613f60565b908152602001604051809103902060000160006101000a81548160ff0219169083600a8111156124e0576124df6135ee565b5b02179055507f10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386828260405161251692919061488b565b60405180910390a15050565b60006005826040516125349190613f60565b90815260200160405180910390206001015414158015612576575060006005826040516125619190613f60565b90815260200160405180910390206002015414155b156125b1574260058260405161258c9190613f60565b9081526020016040518091039020600401819055506125ac81600161249c565b6125d8565b426005826040516125c29190613f60565b9081526020016040518091039020600301819055505b50565b6125e3612d75565b73ffffffffffffffffffffffffffffffffffffffff16612601611282565b73ffffffffffffffffffffffffffffffffffffffff1614612657576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161264e90614907565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600073ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16036127b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127aa90614973565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1603612826576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161281d906149df565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1603612899576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161289090614a4b565b60405180910390fd5b6000816060015151116128e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128d890614ab7565b60405180910390fd5b50565b60008160000151602001511461292f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161292690614b23565b60405180910390fd5b60008160600151602001511461297a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161297190614b8f565b60405180910390fd5b50565b806040015173ffffffffffffffffffffffffffffffffffffffff16826040015173ffffffffffffffffffffffffffffffffffffffff16146129f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129ea90614bfb565b60405180910390fd5b806020015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff1614612a69576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a6090614c67565b60405180910390fd5b806000015173ffffffffffffffffffffffffffffffffffffffff16826000015173ffffffffffffffffffffffffffffffffffffffff1614612adf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ad690614cd3565b60405180910390fd5b80606001515182606001515114612b2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b2290614d3f565b60405180910390fd5b60005b826060015151811015612bf85781606001518181518110612b5257612b51614d5f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1683606001518281518110612b8757612b86614d5f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614612be5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bdc90614dda565b60405180910390fd5b8080612bf090614dfa565b915050612b2e565b505050565b612c0f82600001518260000151612d7d565b612c2182602001518260200151612d7d565b612c3382604001518260400151612d7d565b612c4582606001518260600151612d7d565b5050565b8060000151826000015114612c93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c8a90614e8e565b60405180910390fd5b8060200151826020015114612cdd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cd490614efa565b60405180910390fd5b8060400151826040015114612d27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d1e90614f66565b60405180910390fd5b8060600151826060015114612d71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d6890614fd2565b60405180910390fd5b5050565b600033905090565b8060000151826000015114612dc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dbe9061503e565b60405180910390fd5b8060200151826020015114612e11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e08906150aa565b60405180910390fd5b5050565b6040518060800160405280606081526020016060815260200160608152602001600081525090565b604051806101c001604052806000600a811115612e5d57612e5c6135ee565b5b8152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604051806080016040528060608152602001612ed8612f82565b8152602001612ee5612fec565b8152602001612ef261302c565b81525090565b828054828255906000526020600020908101928215612f71579160200282015b82811115612f705782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190612f18565b5b509050612f7e9190613054565b5090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081525090565b6040518060800160405280612fff613071565b815260200161300c613071565b8152602001613019613071565b8152602001613026613071565b81525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b5b8082111561306d576000816000905550600101613055565b5090565b604051806040016040528060008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006130ca8261309f565b9050919050565b6130da816130bf565b81146130e557600080fd5b50565b6000813590506130f7816130d1565b92915050565b60006020828403121561311357613112613095565b5b6000613121848285016130e8565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015613190578082015181840152602081019050613175565b8381111561319f576000848401525b50505050565b6000601f19601f8301169050919050565b60006131c182613156565b6131cb8185613161565b93506131db818560208601613172565b6131e4816131a5565b840191505092915050565b60006131fb83836131b6565b905092915050565b6000602082019050919050565b600061321b8261312a565b6132258185613135565b93508360208202850161323785613146565b8060005b85811015613273578484038952815161325485826131ef565b945061325f83613203565b925060208a0199505060018101905061323b565b50829750879550505050505092915050565b6000602082019050818103600083015261329f8184613210565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6132e9826131a5565b810181811067ffffffffffffffff82111715613308576133076132b1565b5b80604052505050565b600061331b61308b565b905061332782826132e0565b919050565b600067ffffffffffffffff821115613347576133466132b1565b5b613350826131a5565b9050602081019050919050565b82818337600083830152505050565b600061337f61337a8461332c565b613311565b90508281526020810184848401111561339b5761339a6132ac565b5b6133a684828561335d565b509392505050565b600082601f8301126133c3576133c26132a7565b5b81356133d384826020860161336c565b91505092915050565b6000602082840312156133f2576133f1613095565b5b600082013567ffffffffffffffff8111156134105761340f61309a565b5b61341c848285016133ae565b91505092915050565b60008115159050919050565b61343a81613425565b82525050565b60006020820190506134556000830184613431565b92915050565b6000819050919050565b61346e8161345b565b811461347957600080fd5b50565b60008135905061348b81613465565b92915050565b600080600080608085870312156134ab576134aa613095565b5b600085013567ffffffffffffffff8111156134c9576134c861309a565b5b6134d5878288016133ae565b945050602085013567ffffffffffffffff8111156134f6576134f561309a565b5b613502878288016133ae565b935050604085013567ffffffffffffffff8111156135235761352261309a565b5b61352f878288016133ae565b92505060606135408782880161347c565b91505092959194509250565b6135558161345b565b82525050565b6000608083016000830151848203600086015261357882826131b6565b9150506020830151848203602086015261359282826131b6565b915050604083015184820360408601526135ac82826131b6565b91505060608301516135c1606086018261354c565b508091505092915050565b600060208201905081810360008301526135e6818461355b565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600b811061362e5761362d6135ee565b5b50565b600081905061363f8261361d565b919050565b600061364f82613631565b9050919050565b61365f81613644565b82525050565b6101c08201600082015161367c6000850182613656565b50602082015161368f602085018261354c565b5060408201516136a2604085018261354c565b5060608201516136b5606085018261354c565b5060808201516136c8608085018261354c565b5060a08201516136db60a085018261354c565b5060c08201516136ee60c085018261354c565b5060e082015161370160e085018261354c565b5061010082015161371661010085018261354c565b5061012082015161372b61012085018261354c565b5061014082015161374061014085018261354c565b5061016082015161375561016085018261354c565b5061018082015161376a61018085018261354c565b506101a082015161377f6101a085018261354c565b50505050565b60006101c08201905061379b6000830184613665565b92915050565b6137aa8161345b565b82525050565b60006020820190506137c560008301846137a1565b92915050565b6137d4816130bf565b82525050565b60006020820190506137ef60008301846137cb565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff82111561381a576138196132b1565b5b602082029050602081019050919050565b600080fd5b600061384361383e846137ff565b613311565b905080838252602082019050602084028301858111156138665761386561382b565b5b835b8181101561388f578061387b88826130e8565b845260208401935050602081019050613868565b5050509392505050565b600082601f8301126138ae576138ad6132a7565b5b81356138be848260208601613830565b91505092915050565b6000608082840312156138dd576138dc6137f5565b5b6138e76080613311565b905060006138f7848285016130e8565b600083015250602061390b848285016130e8565b602083015250604061391f848285016130e8565b604083015250606082013567ffffffffffffffff811115613943576139426137fa565b5b61394f84828501613899565b60608301525092915050565b600060408284031215613971576139706137f5565b5b61397b6040613311565b9050600061398b8482850161347c565b600083015250602061399f8482850161347c565b60208301525092915050565b600061010082840312156139c2576139c16137f5565b5b6139cc6080613311565b905060006139dc8482850161395b565b60008301525060406139f08482850161395b565b6020830152506080613a048482850161395b565b60408301525060c0613a188482850161395b565b60608301525092915050565b600060808284031215613a3a57613a396137f5565b5b613a446080613311565b90506000613a548482850161347c565b6000830152506020613a688482850161347c565b6020830152506040613a7c8482850161347c565b6040830152506060613a908482850161347c565b60608301525092915050565b6000806000806101c08587031215613ab757613ab6613095565b5b600085013567ffffffffffffffff811115613ad557613ad461309a565b5b613ae1878288016133ae565b945050602085013567ffffffffffffffff811115613b0257613b0161309a565b5b613b0e878288016138c7565b9350506040613b1f878288016139ab565b925050610140613b3187828801613a24565b91505092959194509250565b613b46816130bf565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000613b848383613b3d565b60208301905092915050565b6000602082019050919050565b6000613ba882613b4c565b613bb28185613b57565b9350613bbd83613b68565b8060005b83811015613bee578151613bd58882613b78565b9750613be083613b90565b925050600181019050613bc1565b5085935050505092915050565b6000608083016000830151613c136000860182613b3d565b506020830151613c266020860182613b3d565b506040830151613c396040860182613b3d565b5060608301518482036060860152613c518282613b9d565b9150508091505092915050565b604082016000820151613c74600085018261354c565b506020820151613c87602085018261354c565b50505050565b61010082016000820151613ca46000850182613c5e565b506020820151613cb76040850182613c5e565b506040820151613cca6080850182613c5e565b506060820151613cdd60c0850182613c5e565b50505050565b608082016000820151613cf9600085018261354c565b506020820151613d0c602085018261354c565b506040820151613d1f604085018261354c565b506060820151613d32606085018261354c565b50505050565b60006101c0830160008301518482036000860152613d5682826131b6565b91505060208301518482036020860152613d708282613bfb565b9150506040830151613d856040860182613c8d565b506060830151613d99610140860182613ce3565b508091505092915050565b60006020820190508181036000830152613dbe8184613d38565b905092915050565b600b8110613dd357600080fd5b50565b600081359050613de581613dc6565b92915050565b60008060408385031215613e0257613e01613095565b5b600083013567ffffffffffffffff811115613e2057613e1f61309a565b5b613e2c858286016133ae565b9250506020613e3d85828601613dd6565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680613e8e57607f821691505b602082108103613ea157613ea0613e47565b5b50919050565b600082825260208201905092915050565b7f526573756c74735375626d697474656400000000000000000000000000000000600082015250565b6000613eee601083613ea7565b9150613ef982613eb8565b602082019050919050565b60006020820190508181036000830152613f1d81613ee1565b9050919050565b600081905092915050565b6000613f3a82613156565b613f448185613f24565b9350613f54818560208601613172565b80840191505092915050565b6000613f6c8284613f2f565b915081905092915050565b7f4465616c41677265656400000000000000000000000000000000000000000000600082015250565b6000613fad600a83613ea7565b9150613fb882613f77565b602082019050919050565b60006020820190508181036000830152613fdc81613fa0565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026140457fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614008565b61404f8683614008565b95508019841693508086168417925050509392505050565b6000819050919050565b600061408c6140876140828461345b565b614067565b61345b565b9050919050565b6000819050919050565b6140a683614071565b6140ba6140b282614093565b848454614015565b825550505050565b600090565b6140cf6140c2565b6140da81848461409d565b505050565b5b818110156140fe576140f36000826140c7565b6001810190506140e0565b5050565b601f8211156141435761411481613fe3565b61411d84613ff8565b8101602085101561412c578190505b61414061413885613ff8565b8301826140df565b50505b505050565b600082821c905092915050565b600061416660001984600802614148565b1980831691505092915050565b600061417f8383614155565b9150826002028217905092915050565b61419882613156565b67ffffffffffffffff8111156141b1576141b06132b1565b5b6141bb8254613e76565b6141c6828285614102565b600060209050601f8311600181146141f957600084156141e7578287015190505b6141f18582614173565b865550614259565b601f19841661420786613fe3565b60005b8281101561422f5784890151825560018201915060208501945060208101905061420a565b8683101561424c5784890151614248601f891682614155565b8355505b6001600288020188555050505b505050505050565b7f4465616c20646f6573206e6f7420657869737400000000000000000000000000600082015250565b6000614297601383613ea7565b91506142a282614261565b602082019050919050565b600060208201905081810360008301526142c68161428a565b9050919050565b7f4a432068617320616c7265616479206167726565640000000000000000000000600082015250565b6000614303601583613ea7565b915061430e826142cd565b602082019050919050565b60006020820190508181036000830152614332816142f6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006143738261345b565b915061437e8361345b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156143b7576143b6614339565b5b828202905092915050565b7f4465616c4e65676f74696174696e670000000000000000000000000000000000600082015250565b60006143f8600f83613ea7565b9150614403826143c2565b602082019050919050565b60006020820190508181036000830152614427816143eb565b9050919050565b7f526573756c7473436865636b6564000000000000000000000000000000000000600082015250565b6000614464600e83613ea7565b915061446f8261442e565b602082019050919050565b6000602082019050818103600083015261449381614457565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60006144f6602e83613ea7565b91506145018261449a565b604082019050919050565b60006020820190508181036000830152614525816144e9565b9050919050565b6000819050919050565b600060ff82169050919050565b600061455e6145596145548461452c565b614067565b614536565b9050919050565b61456e81614543565b82525050565b60006020820190506145896000830184614565565b92915050565b7f52502068617320616c7265616479206167726565640000000000000000000000600082015250565b60006145c5601583613ea7565b91506145d08261458f565b602082019050919050565b600060208201905081810360008301526145f4816145b8565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000614657602683613ea7565b9150614662826145fb565b604082019050919050565b600060208201905081810360008301526146868161464a565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c657220616460008201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b60006146e9603583613ea7565b91506146f48261468d565b604082019050919050565b60006020820190508181036000830152614718816146dc565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e7460008201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b600061477b603983613ea7565b91506147868261471f565b604082019050919050565b600060208201905081810360008301526147aa8161476e565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e7460008201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b600061480d603b83613ea7565b9150614818826147b1565b604082019050919050565b6000602082019050818103600083015261483c81614800565b9050919050565b600061484e82613156565b6148588185613ea7565b9350614868818560208601613172565b614871816131a5565b840191505092915050565b61488581613644565b82525050565b600060408201905081810360008301526148a58185614843565b90506148b4602083018461487c565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006148f1602083613ea7565b91506148fc826148bb565b602082019050919050565b60006020820190508181036000830152614920816148e4565b9050919050565b7f5250206d697373696e6700000000000000000000000000000000000000000000600082015250565b600061495d600a83613ea7565b915061496882614927565b602082019050919050565b6000602082019050818103600083015261498c81614950565b9050919050565b7f4a43206d697373696e6700000000000000000000000000000000000000000000600082015250565b60006149c9600a83613ea7565b91506149d482614993565b602082019050919050565b600060208201905081810360008301526149f8816149bc565b9050919050565b7f536f6c766572206d697373696e67000000000000000000000000000000000000600082015250565b6000614a35600e83613ea7565b9150614a40826149ff565b602082019050919050565b60006020820190508181036000830152614a6481614a28565b9050919050565b7f4d65646961746f7273203c3d2030000000000000000000000000000000000000600082015250565b6000614aa1600e83613ea7565b9150614aac82614a6b565b602082019050919050565b60006020820190508181036000830152614ad081614a94565b9050919050565b7f4167726565206465706f736974206d7573742062652030000000000000000000600082015250565b6000614b0d601783613ea7565b9150614b1882614ad7565b602082019050919050565b60006020820190508181036000830152614b3c81614b00565b9050919050565b7f4d656469617465206465706f736974206d757374206265203000000000000000600082015250565b6000614b79601983613ea7565b9150614b8482614b43565b602082019050919050565b60006020820190508181036000830152614ba881614b6c565b9050919050565b7f5250000000000000000000000000000000000000000000000000000000000000600082015250565b6000614be5600283613ea7565b9150614bf082614baf565b602082019050919050565b60006020820190508181036000830152614c1481614bd8565b9050919050565b7f4a43000000000000000000000000000000000000000000000000000000000000600082015250565b6000614c51600283613ea7565b9150614c5c82614c1b565b602082019050919050565b60006020820190508181036000830152614c8081614c44565b9050919050565b7f536f6c7665720000000000000000000000000000000000000000000000000000600082015250565b6000614cbd600683613ea7565b9150614cc882614c87565b602082019050919050565b60006020820190508181036000830152614cec81614cb0565b9050919050565b7f4d65646961746f72730000000000000000000000000000000000000000000000600082015250565b6000614d29600983613ea7565b9150614d3482614cf3565b602082019050919050565b60006020820190508181036000830152614d5881614d1c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4d65646961746f72000000000000000000000000000000000000000000000000600082015250565b6000614dc4600883613ea7565b9150614dcf82614d8e565b602082019050919050565b60006020820190508181036000830152614df381614db7565b9050919050565b6000614e058261345b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614e3757614e36614339565b5b600182019050919050565b7f5072696365000000000000000000000000000000000000000000000000000000600082015250565b6000614e78600583613ea7565b9150614e8382614e42565b602082019050919050565b60006020820190508181036000830152614ea781614e6b565b9050919050565b7f5061796d656e7400000000000000000000000000000000000000000000000000600082015250565b6000614ee4600783613ea7565b9150614eef82614eae565b602082019050919050565b60006020820190508181036000830152614f1381614ed7565b9050919050565b7f526573756c747300000000000000000000000000000000000000000000000000600082015250565b6000614f50600783613ea7565b9150614f5b82614f1a565b602082019050919050565b60006020820190508181036000830152614f7f81614f43565b9050919050565b7f4d6564696174696f6e0000000000000000000000000000000000000000000000600082015250565b6000614fbc600983613ea7565b9150614fc782614f86565b602082019050919050565b60006020820190508181036000830152614feb81614faf565b9050919050565b7f54696d656f757400000000000000000000000000000000000000000000000000600082015250565b6000615028600783613ea7565b915061503382614ff2565b602082019050919050565b600060208201905081810360008301526150578161501b565b9050919050565b7f436f6c6c61746572616c00000000000000000000000000000000000000000000600082015250565b6000615094600a83613ea7565b915061509f8261505e565b602082019050919050565b600060208201905081810360008301526150c381615087565b905091905056fea26469706673582212200d2467404b80197dc5e0085217e4cfb7f4cd31a03c6e68855a72fbd8f715202064736f6c634300080f0033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCaller) GetAgreement(opts *bind.CallOpts, dealId string) (SharedStructsAgreement, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAgreement", dealId)

	if err != nil {
		return *new(SharedStructsAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsAgreement)).(*SharedStructsAgreement)

	return out0, err

}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetAgreement(dealId string) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0xcdd82d1d.
//
// Solidity: function getAgreement(string dealId) view returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetAgreement(dealId string) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Storage *StorageCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Storage *StorageSession) GetControllerAddress() (common.Address, error) {
	return _Storage.Contract.GetControllerAddress(&_Storage.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Storage *StorageCallerSession) GetControllerAddress() (common.Address, error) {
	return _Storage.Contract.GetControllerAddress(&_Storage.CallOpts)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCaller) GetDeal(opts *bind.CallOpts, dealId string) (SharedStructsDeal, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0xe7079180.
//
// Solidity: function getDeal(string dealId) view returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageCallerSession) GetDeal(dealId string) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(string[])
func (_Storage *StorageCaller) GetDealsForParty(opts *bind.CallOpts, party common.Address) ([]string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDealsForParty", party)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(string[])
func (_Storage *StorageSession) GetDealsForParty(party common.Address) ([]string, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(string[])
func (_Storage *StorageCallerSession) GetDealsForParty(party common.Address) ([]string, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetJobCost is a free data retrieval call binding the contract method 0x73db5c6a.
//
// Solidity: function getJobCost(string dealId) view returns(uint256)
func (_Storage *StorageCaller) GetJobCost(opts *bind.CallOpts, dealId string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getJobCost", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCost is a free data retrieval call binding the contract method 0x73db5c6a.
//
// Solidity: function getJobCost(string dealId) view returns(uint256)
func (_Storage *StorageSession) GetJobCost(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetJobCost is a free data retrieval call binding the contract method 0x73db5c6a.
//
// Solidity: function getJobCost(string dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetJobCost(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,string,uint256))
func (_Storage *StorageCaller) GetResult(opts *bind.CallOpts, dealId string) (SharedStructsResult, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResult", dealId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,string,uint256))
func (_Storage *StorageSession) GetResult(dealId string) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x498cc70d.
//
// Solidity: function getResult(string dealId) view returns((string,string,string,uint256))
func (_Storage *StorageCallerSession) GetResult(dealId string) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0x8224ce5f.
//
// Solidity: function getResultsCollateral(string dealId) view returns(uint256)
func (_Storage *StorageCaller) GetResultsCollateral(opts *bind.CallOpts, dealId string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResultsCollateral", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResultsCollateral is a free data retrieval call binding the contract method 0x8224ce5f.
//
// Solidity: function getResultsCollateral(string dealId) view returns(uint256)
func (_Storage *StorageSession) GetResultsCollateral(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0x8224ce5f.
//
// Solidity: function getResultsCollateral(string dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetResultsCollateral(dealId string) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x2244ad2b.
//
// Solidity: function hasDeal(string dealId) view returns(bool)
func (_Storage *StorageCaller) HasDeal(opts *bind.CallOpts, dealId string) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hasDeal", dealId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeal is a free data retrieval call binding the contract method 0x2244ad2b.
//
// Solidity: function hasDeal(string dealId) view returns(bool)
func (_Storage *StorageSession) HasDeal(dealId string) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x2244ad2b.
//
// Solidity: function hasDeal(string dealId) view returns(bool)
func (_Storage *StorageCallerSession) HasDeal(dealId string) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// IsState is a free data retrieval call binding the contract method 0xb050e74b.
//
// Solidity: function isState(string dealId, uint8 state) view returns(bool)
func (_Storage *StorageCaller) IsState(opts *bind.CallOpts, dealId string, state uint8) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isState", dealId, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xb050e74b.
//
// Solidity: function isState(string dealId, uint8 state) view returns(bool)
func (_Storage *StorageSession) IsState(dealId string, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// IsState is a free data retrieval call binding the contract method 0xb050e74b.
//
// Solidity: function isState(string dealId, uint8 state) view returns(bool)
func (_Storage *StorageCallerSession) IsState(dealId string, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Storage *StorageTransactor) AcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Storage *StorageSession) AcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x297f9e55.
//
// Solidity: function acceptResult(string dealId) returns()
func (_Storage *StorageTransactorSession) AcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns((string,string,string,uint256))
func (_Storage *StorageTransactor) AddResult(opts *bind.TransactOpts, dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addResult", dealId, resultsId, dataId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns((string,string,string,uint256))
func (_Storage *StorageSession) AddResult(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x3955548e.
//
// Solidity: function addResult(string dealId, string resultsId, string dataId, uint256 instructionCount) returns((string,string,string,uint256))
func (_Storage *StorageTransactorSession) AddResult(dealId string, resultsId string, dataId string, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, dataId, instructionCount)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeJobCreator(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x3c4135da.
//
// Solidity: function agreeJobCreator(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeJobCreator(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeResourceProvider(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0xec95b967.
//
// Solidity: function agreeResourceProvider(string dealId) returns((uint8,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeResourceProvider(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Storage *StorageTransactor) CheckResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "checkResult", dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Storage *StorageSession) CheckResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x46834d1e.
//
// Solidity: function checkResult(string dealId) returns()
func (_Storage *StorageTransactorSession) CheckResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Storage.Contract.DisableChangeControllerAddress(&_Storage.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Storage.Contract.DisableChangeControllerAddress(&_Storage.TransactOpts)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0xa6370b0e.
//
// Solidity: function ensureDeal(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageTransactor) EnsureDeal(opts *bind.TransactOpts, dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "ensureDeal", dealId, members, timeouts, pricing)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0xa6370b0e.
//
// Solidity: function ensureDeal(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageSession) EnsureDeal(dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, members, timeouts, pricing)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0xa6370b0e.
//
// Solidity: function ensureDeal(string dealId, (address,address,address,address[]) members, ((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)) timeouts, (uint256,uint256,uint256,uint256) pricing) returns((string,(address,address,address,address[]),((uint256,uint256),(uint256,uint256),(uint256,uint256),(uint256,uint256)),(uint256,uint256,uint256,uint256)))
func (_Storage *StorageTransactorSession) EnsureDeal(dealId string, members SharedStructsDealMembers, timeouts SharedStructsDealTimeouts, pricing SharedStructsDealPricing) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, members, timeouts, pricing)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactorSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Storage *StorageTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Storage *StorageSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x824518aa.
//
// Solidity: function mediationAcceptResult(string dealId) returns()
func (_Storage *StorageTransactorSession) MediationAcceptResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Storage *StorageTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Storage *StorageSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x80ffdfe0.
//
// Solidity: function mediationRejectResult(string dealId) returns()
func (_Storage *StorageTransactorSession) MediationRejectResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetControllerAddress(&_Storage.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetControllerAddress(&_Storage.TransactOpts, _controllerAddress)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutAgree(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutAgree", dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Storage *StorageSession) TimeoutAgree(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutAgree(&_Storage.TransactOpts, dealId)
}

// TimeoutAgree is a paid mutator transaction binding the contract method 0x795f9abf.
//
// Solidity: function timeoutAgree(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutAgree(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutAgree(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Storage *StorageSession) TimeoutJudgeResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0xe850be37.
//
// Solidity: function timeoutJudgeResult(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutJudgeResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Storage *StorageSession) TimeoutMediateResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0xe7b957d1.
//
// Solidity: function timeoutMediateResult(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutMediateResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Storage *StorageTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Storage *StorageSession) TimeoutSubmitResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x511a9f68.
//
// Solidity: function timeoutSubmitResult(string dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutSubmitResult(dealId string) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// StorageDealStateChangeIterator is returned from FilterDealStateChange and is used to iterate over the raw logs and unpacked data for DealStateChange events raised by the Storage contract.
type StorageDealStateChangeIterator struct {
	Event *StorageDealStateChange // Event containing the contract specifics and raw log

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
func (it *StorageDealStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDealStateChange)
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
		it.Event = new(StorageDealStateChange)
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
func (it *StorageDealStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDealStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDealStateChange represents a DealStateChange event raised by the Storage contract.
type StorageDealStateChange struct {
	DealId string
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string dealId, uint8 state)
func (_Storage *StorageFilterer) FilterDealStateChange(opts *bind.FilterOpts) (*StorageDealStateChangeIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DealStateChange")
	if err != nil {
		return nil, err
	}
	return &StorageDealStateChangeIterator{contract: _Storage.contract, event: "DealStateChange", logs: logs, sub: sub}, nil
}

// WatchDealStateChange is a free log subscription operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string dealId, uint8 state)
func (_Storage *StorageFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *StorageDealStateChange) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DealStateChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDealStateChange)
				if err := _Storage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
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

// ParseDealStateChange is a log parse operation binding the contract event 0x10ca3d89184491f5d8de422bd36534abe6eb4b4aa1429a261bdb5ff1dd9ac386.
//
// Solidity: event DealStateChange(string dealId, uint8 state)
func (_Storage *StorageFilterer) ParseDealStateChange(log types.Log) (*StorageDealStateChange, error) {
	event := new(StorageDealStateChange)
	if err := _Storage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
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
		it.Event = new(StorageOwnershipTransferred)
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
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
