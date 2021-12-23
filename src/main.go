package main

import (
	coinbaseAdapter "cryptodata/adapter/coinbase"
	"cryptodata/model"
	coinbaseProcessor "cryptodata/processor/coinbase"
	"cryptodata/useCase/processVWAP"
)

func main() {
	tickerRequest := &model.TickerRequest{
		Type: "subscribe",
		Channels: []model.Channel{
			{
				Name: "full",
				ProductIDs: []string{
					"BTC-USD",
					"ETH-USD",
					"ETH-BTC",
				},
			},
		},
	}

	processVWAP.ExecuteUseCase(tickerRequest,
		coinbaseAdapter.RequestTickers,
		coinbaseProcessor.ProcessCoinbaseResponse)

}
