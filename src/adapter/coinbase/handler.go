package coinbase

import (
	"cryptodata/model"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func getConnection() (*websocket.Conn, error) {
	flag.Parse()
	log.SetFlags(0)
	addr := flag.String("addr", "ws-feed.exchange.coinbase.com", "http service address")

	webSocketURL := url.URL{Scheme: "wss", Host: *addr, Path: "/"}
	log.Printf("connecting to %s\n", webSocketURL.String())

	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL.String(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("websocket connected successfully")
	return conn, nil
}

type RequestTickersFn func(tickerRequest *model.TickerRequest) (chan []byte, error)

func RequestTickers(tickerRequest *model.TickerRequest) (chan []byte, error) {
	conn, err := getConnection()

	if err != nil {
		return nil, err
	}

	wsResponseChannel := make(chan []byte)
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			wsResponseChannel <- message
		}
	}()

	err = conn.WriteJSON(tickerRequest)
	if err != nil {
		return nil, err
	}

	return wsResponseChannel, nil
}
