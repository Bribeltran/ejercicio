package domain

type Product struct {
	ID          int     `json:"id"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}
