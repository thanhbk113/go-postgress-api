package routers

import (
	"thanhbk113/pkg/admin/handler"

	i "thanhbk113/pkg/admin/interface"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	postController handler.PostHandler
}

func PostRoute(rg *gin.RouterGroup, postsv i.PostInterface) {

	var (
		g = rg.Group("/posts")
	)
	controller := handler.PostHandler{
		PostService: postsv,
	}

	g.POST("/", controller.CreatePost)
	g.DELETE("/:postId", controller.DeletePostById)
	g.GET("/:postId", controller.GetPostById)
	// router.GET("/", pc.postController.GetAllPosts)
	// router.PATCH("/:postId", pc.postController.UpdatePost)
}
