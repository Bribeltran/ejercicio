package main

import (
	"database/sql"
	"os"

	"github.com/Bribeltran/ejercicio/project/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"

	"github.com/joho/godotenv"
)

// @title MELI Bootcamp API - Group 1
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	db, _ := sql.Open("mysql", "test_db_user:Test_DB#123@/melisprint")
	_ = godotenv.Load()

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	//no existe docs aun
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}
