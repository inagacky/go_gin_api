package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/go_gin_sample/apps/controller"
)

func GetRouting() *gin.Engine {

	r := gin.Default()
	u := r.Group("/users")
	{
		ctrl := user.Controller{}
		u.GET("/:id", ctrl.GetUser)
	}

	return r
}
