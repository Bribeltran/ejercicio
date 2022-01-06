package routes

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	/* 	r.buildSellerRoutes()
	   	r.buildProductRoutes()
	   	r.buildSectionRoutes()
	   	r.buildWarehouseRoutes()
	   	r.buildEmployeeRoutes()
	   	r.buildBuyerRoutes() */
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

/* func (r *router) buildSellerRoutes() {
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := handler.NewSeller(service)

	sellers := r.rg.Group("/sellers")

	sellers.GET("/", handler.GetAll())
	sellers.GET("/:id", handler.Get())
	sellers.POST("/", TokenAuthMiddleware(), handler.Create())
	sellers.PUT("/:id", TokenAuthMiddleware(), handler.Update())
	sellers.DELETE("/:id", TokenAuthMiddleware(), handler.Delete())

}

func (r *router) buildProductRoutes() {
	repo := products.NewRepository(r.db)
	service := products.NewService(repo)
	p := handler.NewProduct(&service)

	pr := r.rg.Group("/products")
	pr.POST("/", TokenAuthMiddleware(), p.Create())
	pr.GET("/:id", p.Get())
	pr.GET("/", p.GetAll())
	pr.PATCH("/:id", TokenAuthMiddleware(), p.Update())
	pr.DELETE("/:id", TokenAuthMiddleware(), p.Delete())
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := handler.NewSection(service)

	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.POST("/sections", TokenAuthMiddleware(), handler.Create())
	r.rg.PATCH("/sections/:id", TokenAuthMiddleware(), handler.Update())
	r.rg.DELETE("/sections/:id", TokenAuthMiddleware(), handler.Delete())
}

func (r *router) buildWarehouseRoutes() {

	warehouseGroup := r.rg.Group("/warehouses")

	repo := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repo)
	handler := handler.NewWarehouse(service)
	warehouseGroup.GET("", handler.GetAll())
	warehouseGroup.GET(":id", handler.GetById())
	warehouseGroup.DELETE(":id", TokenAuthMiddleware(), handler.Delete())
	warehouseGroup.PATCH(":id", TokenAuthMiddleware(), handler.Update())
	warehouseGroup.POST("", TokenAuthMiddleware(), handler.Create())
}

func (r *router) buildEmployeeRoutes() {
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	controller := handler.NewEmployee(service)
	employeeURL := r.rg.Group("/employees")
	employeeURL.GET("/", controller.GetAll())
	employeeURL.GET("/:id", controller.GetById())
	employeeURL.POST("/", TokenAuthMiddleware(), controller.Create())
	employeeURL.PATCH("/:id", TokenAuthMiddleware(), controller.Update())
	employeeURL.DELETE("/:id", TokenAuthMiddleware(), controller.Delete())
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository(r.db)
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)
	r.rg.GET("/buyers", handler.GetAll())
	r.rg.POST("/buyers", TokenAuthMiddleware(), handler.Create())
	r.rg.GET("/buyers/:id", handler.Get())
	r.rg.PATCH("/buyers/:id", TokenAuthMiddleware(), handler.Update())
	r.rg.DELETE("/buyers/:id", TokenAuthMiddleware(), handler.Delete())
}
*/
