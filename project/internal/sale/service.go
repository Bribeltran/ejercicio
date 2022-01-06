package sale

import (
	"github.com/Bribeltran/ejercicio/project/internal/domain"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Create(ctx *gin.Context, section domain.Sale) (int, error)
	Update(ctx *gin.Context, section domain.Sale, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (ser *service) Create(ctx *gin.Context, sale domain.Sale) (int, error) {

	var err error = nil
	if err != nil {
		return 0, err
	}

	p, err := ser.repository.Save(ctx, sale)

	if err != nil {
		return 0, err
	}
	return p, nil
}

func (ser *service) Update(ctx *gin.Context, sale domain.Sale, id int) error {

	err := ser.repository.Update(ctx, sale)

	return err
}
