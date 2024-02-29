package utils

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/core"
	"github.com/rs/zerolog/log"
)

// CheckInsufficientFunds panics if its insufficient funds error
func CheckInsufficientFunds(err error) {
	if errors.Is(err, core.ErrInsufficientFundsForTransfer) {
		log.Err(err).Caller(5).Msgf("CheckInsufficentFunds")
		panic(fmt.Sprintf("CheckInsufficientFunds"))
	}
}
