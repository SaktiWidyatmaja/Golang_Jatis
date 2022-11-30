package entity

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Quantity    int    `json:"quantity"`
	Imageurl    string `json:"image_url"`
	Description string `json:"description"`
	Sku         string `json:"sku"`
}

type Cart struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	BuyQuantity int    `json:"buy_quantity"`
}
