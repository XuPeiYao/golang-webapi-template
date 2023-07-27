package services

type Product struct {
	Id          int
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

type IProductsServiceProvider interface {
	GetProducts() []Product
}
