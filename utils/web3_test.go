package utils_test

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/core"

	"github.com/CoopHive/hive/utils"
)

const faucetUrl = "https://example.com/faucet"

func TestCheckInSufficientFunds(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		faucetUrl   string
		expectPanic bool
	}{
		{
			name:        "Test for nil error",
			err:         nil,
			expectPanic: false,
		},
		{
			name:        "Test for core.ErrInsufficientFundsForTransfer",
			err:         errors.New("insufficient funds for transfer"),
			expectPanic: true,
		},
		{
			name:        "Test for error containing 'insufficient funds'",
			err:         errors.New("error: insufficient funds"),
			expectPanic: true,
		},
		{
			name:        "Test for empty faucetUrl with generic error",
			err:         errors.New("generic error"),
			faucetUrl:   "",
			expectPanic: false,
		},
		{
			name:        "Test for non-empty faucetUrl with generic error",
			err:         errors.New("generic error"),
			faucetUrl:   faucetUrl,
			expectPanic: false,
		},

		{
			name:        "Test for core eth insufficient funds for transfer",
			err:         core.ErrInsufficientFundsForTransfer,
			faucetUrl:   faucetUrl,
			expectPanic: true,
		},
		{
			name:        "Test for eth insufficient funds",
			err:         core.ErrInsufficientFunds,
			faucetUrl:   faucetUrl,
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.expectPanic {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			utils.CheckInSufficientFunds(tt.err, tt.faucetUrl)
		})
	}
}
