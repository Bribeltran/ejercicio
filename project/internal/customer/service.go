package customer

import (
	"github.com/Bribeltran/ejercicio/project/internal/domain"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Create(ctx *gin.Context, customer domain.Customer) (int, error)
	Update(ctx *gin.Context, customer domain.Customer, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) Create(ctx *gin.Context, customer domain.Customer) (int, error) {

	var err error = nil
	if err != nil {
		return 0, err
	}

	p, err := ser.repository.Save(ctx, customer)

	if err != nil {
		return 0, err
	}
	return p, nil
}

func (ser *service) Update(ctx *gin.Context, customer domain.Customer, id int) error {

	err := ser.repository.Update(ctx, customer)

	return err
}
