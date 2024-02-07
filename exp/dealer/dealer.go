package dealer

import (
	"context"
)

// Dealer defines the interface for the plugin.
type Dealer interface {
	DealMatched(dealID string)
	DealsAgreed() <-chan string
}

type New func(ctx context.Context) Dealer
