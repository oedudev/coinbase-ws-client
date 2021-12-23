package processVWAP

import (
	coinbaseAdapter "cryptodata/adapter/coinbase"
	"cryptodata/model"
	coinbaseProcessor "cryptodata/processor/coinbase"
)

func ExecuteUseCase(tickerRequest *model.TickerRequest,
	requestTickers coinbaseAdapter.RequestTickersFn,
	processCoinbaseResponse coinbaseProcessor.ProcessCoinbaseResponseFn) {

	coinbaseResponseChannel, err := requestTickers(tickerRequest)
	if err != nil {
		panic(err.Error())
	}

	for {
		response := <-coinbaseResponseChannel
		go processCoinbaseResponse(response)
	}
}
