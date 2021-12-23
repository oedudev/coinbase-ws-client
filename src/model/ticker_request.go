package model

type TickerRequest struct {
	Type     string    `json:"type"`
	Channels []Channel `json:"channels"`
}
