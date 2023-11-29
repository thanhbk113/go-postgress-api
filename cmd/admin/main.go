package main

import (
	"log"

	"thanhbk113/internal/config"
	"thanhbk113/pkg/admin/server"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	g := gin.Default()
	server.Bootstrap(g)

	log.Fatal(g.Run(":" + config.GetConfig().ServerPort))
}
