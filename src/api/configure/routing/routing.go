package routing

import (
	"github.com/gin-gonic/gin"
	co "github.com/inagacky/go_gin_api/src/api/controller"
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
			u.POST("", controller.CreateUser)
			u.PUT("/:id", controller.UpdateUser)
			u.DELETE("/:id", controller.DeleteUser)
		}
	}
	return r
}
