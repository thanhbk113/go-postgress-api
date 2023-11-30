package main

import (
	"fmt"
	"log"
	"os"
	"thanhbk113/docs/admin"
	"thanhbk113/internal/config"
	"thanhbk113/pkg/admin/server"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go-Postgres - Admin API
// @version 1.0
// @description All APIs for Go-Postgres admin.
// @description
// @description ******************************
// @description - Add description
// @description ******************************
// @description
// @termsOfService https://bag-manage.vn
// @contact.name Dev team
// @contact.url https://bag-manage.vn
// @basePath /admin

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	g := gin.Default()
	server.Bootstrap(g)

	domain := os.Getenv("DOMAIN_ADMIN")
	admin.SwaggerInfo.Host = domain
	fmt.Println("Swagger host : ", admin.SwaggerInfo.BasePath+"/swagger/*")
	g.GET(admin.SwaggerInfo.BasePath+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	fmt.Println(aurora.Green("Swagger url : " + domain + admin.SwaggerInfo.BasePath + "/swagger/index.html"))

	log.Fatal(g.Run(":" + config.GetConfig().ServerPort))
}
