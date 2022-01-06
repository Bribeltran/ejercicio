package handler

import (
	"net/http"
	"strconv"

	customer "github.com/Bribeltran/ejercicio/project/internal/customer"
	"github.com/Bribeltran/ejercicio/project/internal/domain"
	"github.com/Bribeltran/ejercicio/project/pkg/web"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	service customer.Service
}

func NewCustomer(sect customer.Service) *Customer {
	return &Customer{service: sect}
}

// CreateCustomer godoc
// @Summary create new customer
// @Tags Customers
// @Description create a new customer with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 201 {data} web.Success "Created"
// @Router /api/v1/customers [post]
func (p *Customer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var prd domain.Customer

		err := c.ShouldBindJSON(&prd)

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Hubo un error al querer actualizar %v", err)
		} else {

			response, err := p.service.Create(c, prd)

			if err != nil {
				web.Error(c, http.StatusBadRequest, "No se pudo cargar el customero %v", err)
			} else {
				web.Success(c, http.StatusCreated, response)
			}
		}
	}
}

// UpdateCustomer godoc
// @Summary update customer
// @Tags Customers
// @Description update a new customer with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {data} web.Success "Updated"
// @Router /api/v1/customers/:id [put]
func (s *Customer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var usr domain.Customer

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
