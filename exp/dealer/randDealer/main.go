package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/CoopHive/hive/exp/dealer"
)

// randDealer implements the Dealer interface.
type randDealer struct {
	dealAgreedChan chan string
	ctx            context.Context
	closed         bool
	randGenr       *rand.Rand
}

// New creates a new instance of randDealer.
func New(ctx context.Context) dealer.Dealer {
	return &randDealer{
		dealAgreedChan: make(chan string),
		ctx:            ctx,
		randGenr:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// DealMatched is called when a deal is matched.
func (a *randDealer) DealMatched(dealID string) {
	log.Printf("Deal %s matched\n", dealID)

	// Simulate processing time
	time.Sleep(time.Millisecond * 200)

	if a.closed {
		log.Printf("Cannot process deal %s due to plugin being closed\n", dealID)
		return
	}

	agreeCond := a.randGenr.Intn(2) == 0

	if agreeCond {
		log.Printf("Deal:%s is agreed\n", dealID)
		a.agree(dealID)
		return
	}
	log.Printf("Deal:%s not agreed\n", dealID)
}

func (a *randDealer) agree(dealID string) {
	select {
	case <-a.ctx.Done():
		log.Printf("Cannot process deal %s due to stop signal: %v\n", dealID, a.ctx.Err())
		a.close()
	case a.dealAgreedChan <- dealID:
		log.Printf("Deal %s successfully agreed\n", dealID)
	}
}

// DealsAgreed returns a channel to receive agreed deals.
func (a *randDealer) DealsAgreed() <-chan string {
	return a.dealAgreedChan
}

func (a *randDealer) close() {
	if !a.closed {
		a.closed = true
		close(a.dealAgreedChan)
		log.Println("successfully closed")
	}
}

func main() {
	// Create a context with cancellation ability
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()

	// Instantiate the plugin
	dealer := New(ctx)

	// Simulate deals being matched
	go func() {
		for i := 0; i < 100; i++ {
			dealID := fmt.Sprintf("deal-%d", i)
			dealer.DealMatched(dealID)
		}
	}()

	time.Sleep(time.Second)
	// Receive agreed deals
RECV_AGREE_DEALS:
	for {
		select {
		case dealID, ok := <-dealer.DealsAgreed():
			if !ok {
				log.Println("Channel closed. Exiting...")
				break RECV_AGREE_DEALS
			}
			log.Printf("Deal %s rcvd \n", dealID)
		case <-ctx.Done():
			log.Println("Context done. Exiting...")
			return
		}
	}
}
