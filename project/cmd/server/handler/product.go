package handler

import (
	"net/http"
	"strconv"

	"github.com/Bribeltran/ejercicio/project/internal/domain"
	product "github.com/Bribeltran/ejercicio/project/internal/product"
	"github.com/Bribeltran/ejercicio/project/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service product.Service
}

func NewProduct(sect product.Service) *Product {
	return &Product{service: sect}
}

// CreateProduct godoc
// @Summary create new product
// @Tags Products
// @Description create a new product with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 201 {data} web.Success "Created"
// @Router /api/v1/products [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var prd domain.Product

		err := c.ShouldBindJSON(&prd)

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Hubo un error al querer actualizar %v", err)
		} else {

			response, err := p.service.Create(c, prd)

			if err != nil {
				web.Error(c, http.StatusBadRequest, "No se pudo cargar el producto %v", err)
			} else {
				web.Success(c, http.StatusCreated, response)
			}
		}
	}
}

// UpdateProduct godoc
// @Summary update product
// @Tags Products
// @Description update a new product with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {data} web.Success "Updated"
// @Router /api/v1/products/:id [put]
func (s *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var usr domain.Product

		err := c.ShouldBindJSON(&usr)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Hubo un error al querer actualizar %v", err)
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			web.Error(c, http.StatusBadRequest, "Hubo un error al querer actualizar %v", err)
		} else {

			err := s.service.Update(c, usr, int(id))
			if err != nil {
				web.Error(c, http.StatusConflict, err.Error())

			} else {
				usr.ID = int(id)
				web.Success(c, http.StatusOK, usr)
			}
		}
	}
}
