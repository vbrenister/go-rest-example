package model

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Inventory int    `json:"inventory"`
	Price     int    `json:"price"`
}
