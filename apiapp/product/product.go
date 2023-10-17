package product

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ProductListResponse struct {
	Products []*Product `json:"products"`
}
