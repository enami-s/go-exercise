package product

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ProductsResponse struct {
	Products []*Product `json:"products"`
}
