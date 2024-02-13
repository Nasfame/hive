package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/websocket"

	"github.com/CoopHive/hive/exp/dealer"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin, you might want to restrict this in production
	},
}

type WebSocketDealer struct {
	conn        *websocket.Conn
	agreedDeals chan string
}

const (
	DealMatched = iota
	DealAgreed
)

// DealMatched sends the deal ID to the WebSocket connection.
func (d *WebSocketDealer) DealMatched(dealID string) {
	if d.conn == nil {
		log.Println("No active WebSocket connection.")
		return
	}
	err := d.conn.WriteMessage(websocket.TextMessage, []byte("DealMatched-"+dealID))
	if err != nil {
		fmt.Println("Failed to send deal:", err)
	}
}

// DealsAgreed listens for messages on the WebSocket connection.
func (d *WebSocketDealer) DealsAgreed() <-chan string {
	return d.agreedDeals
}

func (d *WebSocketDealer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	defer func() {
		d.conn = nil
	}()

	if d.conn == nil {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Failed to upgrade connection:", err)
			return
		}
		d.conn = conn
	} else {
		log.Println("Already connected; Max allowed one connection")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Already connected; Max allowed one connection"))
		return
	}
	conn := d.conn

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
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
	d := &WebSocketDealer{agreedDeals: make(chan string, 1)}

	http.HandleFunc("/", d.handleWebSocket)
	go func() {
		log.Println("Server listening on :8080")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
	}()
	return d
}

func main() {
	ctxt := context.Background()
	d := New(ctxt).(*WebSocketDealer)

	d.DealMatched("123")

	select {}

}
