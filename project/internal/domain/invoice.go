package domain

type Invoice struct {
	ID         int     `json:"id"`
	DateTime   string  `json:"data_time" binding:"required"`
	IdCustomer int     `json:"id_customer" binding:"required"`
	Total      float64 `json:"total" binding:"required"`
}
