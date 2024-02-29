package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/core"
	"github.com/rs/zerolog/log"
)

// CheckInSufficientFunds panics if its insufficient funds error
func CheckInSufficientFunds(err error, faucetUrl string) {

	if err == nil {
		return
	}

	if errors.Is(err, core.ErrInsufficientFundsForTransfer) {
		log.Err(err).Caller(5).Msgf("CheckInsufficentFunds")

		if faucetUrl != "" {
			log.Info().Msgf("checkout our faucets over here:%v", faucetUrl)
		}
		panic(fmt.Sprintf("CheckInSufficientFunds"))
	}

	if strings.Contains(err.Error(), "insufficient funds") {
		log.Debug().Err(err).Caller(3).Msgf("CheckInSufficientFunds-I")
		panic(err)
	}
}
