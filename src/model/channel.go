package model

type Channel struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}
