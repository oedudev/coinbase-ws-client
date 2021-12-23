package model

import "time"

type TradeMatch struct {
	Type        string    `json:"type"`
	TradeID     int64     `json:"trade_id"`
	Sequence    int64     `json:"sequence"`
	MakeOrderID string    `json:"make_order_id"`
	TakeOrderID string    `json:"take_order_id"`
	Time        time.Time `json:"time"`
	ProductID   string    `json:"product_id"`
	Size        string    `json:"size"`
	Price       string    `json:"price"`
	Side        string    `json:"side"`
}
