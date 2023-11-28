package routes

import (
	"thanhbk113/controllers"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	postController controllers.PostController
}

func NewRoutePost(postController controllers.PostController) PostRoutes {
	return PostRoutes{postController}
}

func (pc *PostRoutes) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.POST("/", pc.postController.CreatePost)
	router.GET("/", pc.postController.GetAllPosts)
	router.PATCH("/:postId", pc.postController.UpdatePost)
	router.GET("/:postId", pc.postController.GetPostById)
	router.DELETE("/:postId", pc.postController.DeletePostById)
}
