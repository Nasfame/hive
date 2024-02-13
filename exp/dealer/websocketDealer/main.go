package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"github.com/CoopHive/hive/exp/dealer"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin, you might want to restrict this in production
	},
}

type WebSocketDealer struct {
	ctx          context.Context
	matchedDeals chan string
	agreedDeals  chan string
}

const (
	DealMatched = iota
	DealAgreed
)

// DealMatched called by hive when a deal is matched
func (d *WebSocketDealer) DealMatched(dealID string) {
	d.matchedDeals <- dealID
}

// DealsAgreed sends the deal ID to dealsAgreed channel for hive
func (d *WebSocketDealer) DealsAgreed() <-chan string {
	return d.agreedDeals
}

func (d *WebSocketDealer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}

	go func() {
		for {
			select {

			case <-r.Context().Done():
				log.Println("connection closed")
				return

			case dealID := <-d.matchedDeals:
				log.Println("send: matched deal", dealID)
				err := conn.WriteMessage(websocket.TextMessage, []byte("DealMatched-"+dealID))
				if err != nil {
					fmt.Println("Failed to send deal:", err)
				}

			}
		}
	}()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			if r.Context().Done() != nil {
				return
			}
			continue

		}

		if msgType != websocket.TextMessage {
			fmt.Println("Received non-text message:", msg)
			continue
		}

		log.Println("rcvd dealAgreed", string(msg))

		d.agreedDeals <- string(msg)
	}
}

var PORT = func() int {
	port := os.Getenv("PORT")

	if p, err := strconv.Atoi(port); err != nil {
		return 8080
	} else {
		return p
	}

}()

func New(ctx context.Context) dealer.Dealer {
	d := &WebSocketDealer{
		ctx:          ctx,
		matchedDeals: make(chan string, 1),
		agreedDeals:  make(chan string, 1),
	}

	http.HandleFunc("/", d.handleWebSocket)
	go func() {
		log.Printf("Websocket started on 0.0.0.0:%d\n", PORT)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
	}()
	return d
}

func main() {
	ctxt := context.Background()
	d := New(ctxt).(*WebSocketDealer)

	ticker := time.Tick(time.Second * 5)

	dealID := 0

	for {
		select {
		case <-ticker:
			dealID++
			d.DealMatched(strconv.Itoa(dealID))
		case <-d.ctx.Done():
			log.Println("Main:Context done")
			return
		}
	}

}
