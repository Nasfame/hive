package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CoopHive/hive/exp/dealer"
)

// AutoDealer implements the Services interface.
type AutoDealer struct {
	dealAgreedChan chan string
	ctx            context.Context
	closed         bool
}

// New creates a new instance of AutoDealer.
func NewAutoDealer(ctx context.Context) dealer.Dealer {
	return &AutoDealer{
		dealAgreedChan: make(chan string),
		ctx:            ctx,
	}
}

// DealMatched is called when a deal is matched.
func (a *AutoDealer) DealMatched(dealID string) {
	fmt.Printf("Deal %s is matched\n", dealID)

	// Simulate processing time
	time.Sleep(time.Millisecond * 200)

	if a.closed {
		log.Printf("Cannot process deal %s due to plugin being closed\n", dealID)
		return
	}

	select {
	case <-a.ctx.Done():
		log.Printf("Cannot process deal %s due to stop signal: %v\n", dealID, a.ctx.Err())
		a.close()
	case a.dealAgreedChan <- dealID:
		log.Printf("Deal %s is agreed\n", dealID)
	}
}

// DealsAgreed returns a channel to receive agreed deals.
func (a *AutoDealer) DealsAgreed() <-chan string {
	return a.dealAgreedChan
}

func (a *AutoDealer) close() {
	if !a.closed {
		a.closed = true
		close(a.dealAgreedChan)
		log.Println("successfully closed")
	}
}
