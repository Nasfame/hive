package autoacceptdeals

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

// AutoAcceptPlugin implements the Services interface.
type AutoAcceptPlugin struct {
	dealAgreedChan chan string
	ctx            context.Context
	closed         *atomic.Bool
}

// New creates a new instance of AutoAcceptPlugin.
func New(ctx context.Context) *AutoAcceptPlugin {
	return &AutoAcceptPlugin{
		dealAgreedChan: make(chan string, 10),
		ctx:            ctx,
		closed:         &atomic.Bool{},
	}
}

// DealsMatched is called when a deal is matched.
func (a *AutoAcceptPlugin) DealsMatched(dealID string) {
	fmt.Printf("Deal %s is matched\n", dealID)
	// Simulate processing time
	time.Sleep(time.Millisecond * 200)
	select {
	case <-a.ctx.Done():
		log.Printf("Cannot process deal %s due to stop signal: %v", dealID, a.ctx.Err())
		a.close()
		return

	case a.dealAgreedChan <- dealID:
		log.Printf("Deal %s is agreed", dealID)
	}
}

// DealsAgreed returns a channel to receive agreed deals.
func (a *AutoAcceptPlugin) DealsAgreed() <-chan string {
	return a.dealAgreedChan
}

func (a *AutoAcceptPlugin) close() {
	if a.closed.Load() {
		return
	}
	for len(a.dealAgreedChan) != 0 {
		<-a.dealAgreedChan
	}
	close(a.dealAgreedChan)
	a.closed.Store(true)
	log.Println("successfully closed")
}

func main() {
	// Create a context with cancellation ability
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))

	// Instantiate the plugin
	dealer := New(ctx)

	// Simulate deals being matched
	go func() {
		for i := 0; i < 100; i++ {
			dealID := fmt.Sprintf("deal-%d", i)
			dealer.DealsMatched(dealID)
		}
	}()

	time.Sleep(time.Second)
	// Receive agreed deals
RECV_AGREE_DEALS:
	for i := 0; i < 100; i++ {
		select {
		case dealID, ok := <-dealer.DealsAgreed():

			if !ok {
				fmt.Println("Channel closed. Exiting...")
				break RECV_AGREE_DEALS
			}

			fmt.Printf("Deal %s is agreed upon\n", dealID)
		case <-ctx.Done():
			fmt.Println("Context done. Exiting...")
			return
		}
	}
}
