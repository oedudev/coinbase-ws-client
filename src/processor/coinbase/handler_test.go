package coinbase

import (
	"cryptodata/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkProcessTicker(b *testing.B) {

	for i := 0; i < b.N; i++ {
		processTicker(&model.TradeMatch{
			ProductID: "BTC-USD",
			Size:      "12",
			Price:     "3548",
		})
	}
}

func TestCalculateVWAP(t *testing.T) {

	tradeMatches := []model.TradeMatch{
		{
			ProductID: "BTC-USD",
			Size:      "12",
			Price:     "3548",
		},
		{
			ProductID: "BTC-USD",
			Size:      "8",
			Price:     "3418",
		},
		{
			ProductID: "BTC-USD",
			Size:      "5",
			Price:     "3128",
		},
	}

	vwap, err := calculateVWAP(tradeMatches)
	assert.Nil(t, err, "error testing calculateVWAP")
	assert.Equal(t, 3422.4, vwap)
}
