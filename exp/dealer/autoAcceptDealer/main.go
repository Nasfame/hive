package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CoopHive/hive/exp/dealer"
)

// AutoAcceptPlugin implements the Dealer interface.
type AutoAcceptPlugin struct {
	dealAgreedChan chan string
	ctx            context.Context
	closed         bool
}

// New creates a new instance of AutoAcceptPlugin.
func New(ctx context.Context) dealer.Dealer {
	return &AutoAcceptPlugin{
		dealAgreedChan: make(chan string),
		ctx:            ctx,
	}
}

// DealMatched is called when a deal is matched.
func (a *AutoAcceptPlugin) DealMatched(dealID string) {
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
func (a *AutoAcceptPlugin) DealsAgreed() <-chan string {
	return a.dealAgreedChan
}

func (a *AutoAcceptPlugin) close() {
	if !a.closed {
		a.closed = true
		close(a.dealAgreedChan)
		log.Println("successfully closed")
	}
}

func main() {
	// Create a context with cancellation ability
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*500))
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
			log.Println("Deal %s is agreed upon\n", dealID)
		case <-ctx.Done():
			log.Println("Context done. Exiting...")
			return
		}
	}
}
