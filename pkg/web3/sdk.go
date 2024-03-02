package web3

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3/bindings/controller"
	"github.com/CoopHive/hive/pkg/web3/bindings/jobcreator"
	"github.com/CoopHive/hive/pkg/web3/bindings/mediation"
	"github.com/CoopHive/hive/pkg/web3/bindings/payments"
	"github.com/CoopHive/hive/pkg/web3/bindings/storage"
	"github.com/CoopHive/hive/pkg/web3/bindings/token"
	"github.com/CoopHive/hive/pkg/web3/bindings/users"
	"github.com/CoopHive/hive/utils"
)

// these are the go-binding wrappers for the various deployed contracts
type Contracts struct {
	Token      *token.Token
	Payments   *payments.Payments
	Storage    *storage.Storage
	Users      *users.Users
	JobCreator *jobcreator.Jobcreator
	Mediation  *mediation.Mediation
	Controller *controller.Controller
}

type Web3SDK struct {
	Options      Web3Options
	PrivateKey   *ecdsa.PrivateKey
	Client       *ethclient.Client
	CallOpts     *bind.CallOpts
	TransactOpts *bind.TransactOpts
	Contracts    *Contracts
}

func NewContracts(
	options Web3Options,
	client *ethclient.Client,
	callOpts *bind.CallOpts,
) (*Contracts, error) {
	controller, err := controller.NewController(common.HexToAddress(options.ControllerAddress), client)
	if err != nil {
		return nil, err
	}

	paymentsAddress := options.PaymentsAddress
	log.Debug().Msgf("paymentsAddress: %s", paymentsAddress)
	if paymentsAddress == "" {
		log.Debug().Msgf("payment address not found")
		loadedPaymentsAddress, err := controller.GetPaymentsAddress(callOpts)
		if err != nil {
			log.Error().Err(err)
			return nil, err
		}
		paymentsAddress = loadedPaymentsAddress.String()
		log.Debug().
			Str("load payments address", paymentsAddress).
			Msgf("")
	}
	payments, err := payments.NewPayments(common.HexToAddress(paymentsAddress), client)
	if err != nil {
		log.Error().Msgf("payments contract connection failed with %v", err)
		return nil, err
	}

	tokenAddress := options.TokenAddress
	log.Debug().Msgf("TokenAddress: %s", tokenAddress)
	if tokenAddress == "" {
		loadedTokenAddress, err := payments.GetTokenAddress(callOpts)
		if err != nil {
			return nil, err
		}
		tokenAddress = loadedTokenAddress.String()
		log.Debug().
			Str("load token address", tokenAddress).
			Msgf("")
	}

	token, err := token.NewToken(common.HexToAddress(tokenAddress), client)
	if err != nil {
		return nil, err
	}

	storageAddress := options.StorageAddress
	log.Debug().Msgf("StorageAddress: %s", storageAddress)
	if storageAddress == "" {
		loadedStorageAddress, err := controller.GetStorageAddress(callOpts)
		if err != nil {
			return nil, err
		}
		storageAddress = loadedStorageAddress.String()
		log.Debug().
			Str("load storage address", storageAddress).
			Msgf("")
	}

	storage, err := storage.NewStorage(common.HexToAddress(storageAddress), client)
	if err != nil {
		return nil, err
	}

	usersAddress := options.UsersAddress
	log.Debug().Msgf("UsersAddress: %s", usersAddress)
	if usersAddress == "" {
		loadedUsersAddress, err := controller.GetUsersAddress(callOpts)
		if err != nil {
			return nil, err
		}
		usersAddress = loadedUsersAddress.String()
		log.Debug().
			Str("load users address", usersAddress).
			Msgf("")
	}

	users, err := users.NewUsers(common.HexToAddress(usersAddress), client)
	if err != nil {
		return nil, err
	}

	jobcreatorAddress := options.JobCreatorAddress
	log.Debug().Msgf("JobCreatorAddress: %s", jobcreatorAddress)
	if jobcreatorAddress == "" {
		loadedJobCreatorAddress, err := controller.GetJobCreatorAddress(callOpts)
		if err != nil {
			return nil, err
		}
		jobcreatorAddress = loadedJobCreatorAddress.String()
		log.Debug().
			Str("load jobcreator address", jobcreatorAddress).
			Msgf("")
	}

	jobCreator, err := jobcreator.NewJobcreator(common.HexToAddress(jobcreatorAddress), client)
	if err != nil {
		return nil, err
	}

	mediationAddress := options.MediationAddress
	log.Debug().Msgf("MediationAddress: %s", mediationAddress)
	if mediationAddress == "" {
		loadedMediationAddress, err := controller.GetMediationAddress(callOpts)
		if err != nil {
			return nil, err
		}
		mediationAddress = loadedMediationAddress.String()
		log.Debug().
			Str("load mediation address", mediationAddress).
			Msgf("")
	}

	mediation, err := mediation.NewMediation(common.HexToAddress(mediationAddress), client)
	if err != nil {
		return nil, err
	}

	return &Contracts{
		Token:      token,
		Payments:   payments,
		Storage:    storage,
		Users:      users,
		JobCreator: jobCreator,
		Mediation:  mediation,
		Controller: controller,
	}, nil
}

func NewContractSDK(options Web3Options) (*Web3SDK, error) {
	// write to console
	log.Debug().Msgf("NewContractSDK: %+v", options)
	client, err := ethclient.Dial(options.RpcURL)
	if err != nil {
		log.Error().Msgf("error for client: %s", err.Error())
		return nil, err
	}
	privateKey, err := ParsePrivateKey(options.PrivateKey)

	log.Debug().Msgf("signer: %s", GetAddress(privateKey))

	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(options.ChainID)))
	if err != nil {
		log.Error().Msgf("error for transactOpts: %s", err.Error())
		return nil, err
	}

	contracts, err := NewContracts(options, client, callOpts)
	if err != nil {
		log.Error().Msgf("error for contracts: %s", err.Error())
		return nil, err
	}
	return &Web3SDK{
		PrivateKey:   privateKey,
		Options:      options,
		Client:       client,
		CallOpts:     callOpts,
		TransactOpts: transactOpts,
		Contracts:    contracts,
	}, nil
}

func (sdk *Web3SDK) getBlockNumber() (uint64, error) {
	var blockNumberHex string
	err := sdk.Client.Client().Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		log.Error().Msgf("error for getBlockNumber: %s", err.Error())
		return 0, err
	}
	blockNumberHex = strings.TrimPrefix(blockNumberHex, "0x")
	return strconv.ParseUint(blockNumberHex, 16, 64)
}

func (sdk *Web3SDK) WaitTx(ctx context.Context, tx *types.Transaction, errCreateTx error) (r *types.Receipt, err error) {
	defer func() {
		sdk.checkInsufficientBalance(err)
	}()

	if errCreateTx != nil {
		system.Error(sdk.Options.Service, "error submitting creatingTx", err)
		// creating transaction also fails if insufficient funds
		return nil, errCreateTx // errors.CombineErrors(fmt.Errorf("failed to create tx"), errCreateTx)
	}

	if tx == nil {
		log.Error().Err(err).Msgf("%v", sdk.Options.Service)
		// panic("invalid tx")
		return nil, fmt.Errorf("invalid tx")
	}

	system.Debug(sdk.Options.Service, "submitted tx", tx.Hash().String())
	// system.DumpObjectDebug(tx)

	return bind.WaitMined(ctx, sdk.Client, tx)
}

func (sdk *Web3SDK) GetAddress() common.Address {
	return crypto.PubkeyToAddress(GetPublicKey(sdk.PrivateKey))
}

func (sdk *Web3SDK) checkInsufficientBalance(err error) {
	// if err != nil {
	// 	log.Error().Caller(2).Err(err).Msgf("CheckInsufficientBalance")
	// }
	utils.CheckInSufficientFunds(err, config.Conf.GetString(enums.GITHUB_REPO))
}
