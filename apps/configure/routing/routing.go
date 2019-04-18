package routing

import (
	"github.com/gin-gonic/gin"
	co "github.com/go_gin_sample/apps/controller"
)

func GetRouting() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		// ユーザー関係のAPI
		u := v1.Group("/users")
		{
			controller :=  co.UserController{}
			u.GET("/:id", controller.GetUser)
//			u.POST("", controller.CreateUser)
		}
	}


	return r
}
