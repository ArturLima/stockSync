package requests

type ProductRequest struct {
	Products []Product `json:"products" binding:"required"`
}

type Product struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Qtd  int    `json:"qtd" binding:"required"`
}
