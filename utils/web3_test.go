package utils_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core"

	"github.com/CoopHive/hive/utils"
)

const faucetUrl = "https://example.com/faucet"

func TestCheckInSufficientFunds(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		errString   string
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
		{
			name:        "Test for calibration insufficient funds",
			err:         fmt.Errorf("actor not found"),
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			defer func() {
				r := recover()
				if (r != nil) != tt.expectPanic {
					t.Errorf("Unexpected panic: %v", r)
				}
			}()

			utils.PanicOnInsufficientFunds(tt.err, tt.faucetUrl)
		})
	}
}

func TestPanicOnHTTPUrl(t *testing.T) {
	type TestCase struct {
		name        string
		url         string
		expectPanic bool
	}
	tests := []TestCase{
		{
			name:        "Test for empty URL",
			url:         "",
			expectPanic: true,
		},
		{
			name:        "Test for non-empty HTTP URL",
			url:         "http://example.com",
			expectPanic: true,
		},
		{
			name:        "Test for non-empty HTTPS URL",
			url:         "https://example.com",
			expectPanic: true,
		},
		{
			name:        "Test for non-empty WebSocket URL",
			url:         "ws://example.com",
			expectPanic: false,
		},
		{
			name:        "Test for non-empty Secure WebSocket URL",
			url:         "wss://example.com",
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(tt TestCase) {
				r := recover()
				if (r != nil) != tt.expectPanic {
					t.Errorf("Unexpected panic: %v", r)
					t.Errorf("Test case ID: %s", tt.name)
				}
			}(tt)

			utils.PanicOnHTTPUrl(tt.url)
		})
	}
}
