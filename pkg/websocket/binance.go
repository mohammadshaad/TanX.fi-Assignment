package websocket

import (
    "log"
    "github.com/gorilla/websocket"
)

func ConnectToBinance() {
    c, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/btcusdt@ticker", nil)
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    for {
        _, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", message)
        // Process the message and check for price alerts
    }
}
