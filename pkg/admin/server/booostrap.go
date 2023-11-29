package server

import (
	"thanhbk113/pkg/admin/routers"
	"thanhbk113/pkg/admin/server/initialize"

	"github.com/gin-gonic/gin"
)

func Bootstrap(g *gin.Engine) {
	initialize.Init()
	routers.Init(g)
}
