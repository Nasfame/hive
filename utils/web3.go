package utils

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/core"
	"github.com/rs/zerolog/log"
)

const msgInsufficientFunds = "PanicOnInsufficientFunds"

// PanicOnInsufficientFunds panics if its insufficient funds error
func PanicOnInsufficientFunds(err error, faucetUrl string) {

	if err == nil {
		return
	}

	if strings.Contains(err.Error(), "actor not found") {
		panic(msgInsufficientFunds + "calibration")
	}

	if errors.Is(err, core.ErrInsufficientFundsForTransfer) {
		log.Err(err).Caller(5).Msgf("CheckInsufficentFunds")

		if faucetUrl != "" {
			log.Info().Msgf("checkout our faucets over here:%v", faucetUrl)
		}
		panic(msgInsufficientFunds)
	}

	if strings.Contains(err.Error(), "insufficient funds") {
		log.Debug().Err(err).Caller(3).Msgf(msgInsufficientFunds + "1")
		panic(err)
	}
}
