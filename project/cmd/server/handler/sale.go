package handler

import (
	"net/http"
	"strconv"

	"github.com/Bribeltran/ejercicio/project/internal/domain"
	sale "github.com/Bribeltran/ejercicio/project/internal/sale"
	"github.com/Bribeltran/ejercicio/project/pkg/web"
	"github.com/gin-gonic/gin"
)

type Sale struct {
	service sale.Service
}

func NewSale(sect sale.Service) *Sale {
	return &Sale{service: sect}
}

// CreateSale godoc
// @Summary create new sale
// @Tags Sales
// @Description create a new sale with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 201 {data} web.Success "Created"
// @Router /api/v1/sales [post]
func (p *Sale) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var prd domain.Sale

		err := c.ShouldBindJSON(&prd)

		if err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Hubo un error al querer actualizar %v", err)
		} else {

			response, err := p.service.Create(c, prd)

			if err != nil {
				web.Error(c, http.StatusBadRequest, "No se pudo cargar el saleo %v", err)
			} else {
				web.Success(c, http.StatusCreated, response)
			}
		}
	}
}

// UpdateSale godoc
// @Summary update sale
// @Tags Sales
// @Description update a new sale with the body parameters required
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {data} web.Success "Updated"
// @Router /api/v1/sales/:id [put]
func (s *Sale) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var usr domain.Sale

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
