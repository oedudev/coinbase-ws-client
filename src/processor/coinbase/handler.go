package coinbase

import (
	"cryptodata/model"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

type ProcessCoinbaseResponseFn func(response []byte)

func ProcessCoinbaseResponse(response []byte) {
	ticker := &model.Ticker{}

	err := json.Unmarshal(response, ticker)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	if ticker.Type != "match" {
		return
	}

	tradeMatch := &model.TradeMatch{}

	err = json.Unmarshal(response, tradeMatch)

	vwap, err := processTicker(tradeMatch)
	if err != nil {
		fmt.Printf("error on processTicker: %s\n", err.Error())
	}

	fmt.Printf("(%s) VWAP: %f\n", tradeMatch.ProductID, vwap)
}

var tickerMap map[string][]model.TradeMatch
var globalMutex sync.Mutex

func init() {
	tickerMap = make(map[string][]model.TradeMatch)
}

func processTicker(tradeMatch *model.TradeMatch) (float64, error) {
	globalMutex.Lock()

	tradeMatches, ok := tickerMap[tradeMatch.ProductID]
	if !ok {
		tickerMap[tradeMatch.ProductID] = []model.TradeMatch{}
	}

	tradeMatches = append(tradeMatches, *tradeMatch)
	if len(tradeMatches) >= 200 {
		tradeMatches = tradeMatches[:200]
	}
	tickerMap[tradeMatch.ProductID] = tradeMatches

	globalMutex.Unlock()

	vwap, err := calculateVWAP(tradeMatches)
	if err != nil {
		return 0, err
	}

	return vwap, nil
}

func calculateVWAP(tradeMatches []model.TradeMatch) (float64, error) {

	var globalPrice float64
	var globalVolume float64

	for idx := 0; idx < len(tradeMatches); idx++ {
		tradeMatch := tradeMatches[idx]

		price, err := strconv.ParseFloat(tradeMatch.Price, 8)
		if err != nil {
			return 0, err
		}

		size, err := strconv.ParseFloat(tradeMatch.Size, 8)
		if err != nil {
			return 0, err
		}

		globalPrice = globalPrice + (price * size)
		globalVolume = globalVolume + size
	}

	result := globalPrice / globalVolume

	return result, nil
}
