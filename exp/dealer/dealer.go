package dealer

import (
	"context"
)

// Dealer defines the interface for a plugin that manages deals within the Hive application.
// Implementations of this interface are expected to handle matched deals and provide a channel
// for agreed deals.
type Dealer interface {
	// DealMatched processes a deal identified by its unique ID.
	DealMatched(dealID string)

	// DealsAgreed returns a read-only channel that emits IDs of deals that have been agreed upon.
	DealsAgreed() <-chan string
}

// New is a constructor type for creating new instances of the Dealer interface.
// It takes a context and returns an implementation of the Dealer interface.
type New func(ctx context.Context) Dealer
