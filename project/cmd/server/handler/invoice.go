package handler

import (
	"net/http"
	"strconv"

	"github.com/Bribeltran/ejercicio/project/internal/domain"
	invoice "github.com/Bribeltran/ejercicio/project/internal/invoice"
	"github.com/Bribeltran/ejercicio/project/pkg/web"
	"github.com/gin-gonic/gin"
)

type Invoice struct {
	service invoice.Service
}

func NewInvoice(sect invoice.Service) *Invoice {
	return &Invoice{service: sect}
}

// CreateInvoice godoc
// @Summary create new invoice
// @Tags Invoices
// @Description create a new invoice with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 201 {data} web.Success "Created"
// @Router /api/v1/invoices [post]
func (p *Invoice) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var prd domain.Invoice

		err := c.ShouldBindJSON(&prd)

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Hubo un error al querer actualizar %v", err)
		} else {

			response, err := p.service.Create(c, prd)

			if err != nil {
				web.Error(c, http.StatusBadRequest, "No se pudo cargar el invoiceo %v", err)
			} else {
				web.Success(c, http.StatusCreated, response)
			}
		}
	}
}

// UpdateInvoice godoc
// @Summary update invoice
// @Tags Invoices
// @Description update a new invoice with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {data} web.Success "Updated"
// @Router /api/v1/invoices/:id [put]
func (s *Invoice) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var usr domain.Invoice

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
