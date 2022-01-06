package domain

type Sale struct {
	ID        int     `json:"id"`
	IdInvoice int     `json:"id_invoice" binding:"required"`
	IdProduct int     `json:"id_product" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required"`
}
