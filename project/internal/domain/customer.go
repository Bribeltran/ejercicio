package domain

type Customer struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Condition string `json:"condition" binding:"required"`
}
