package routing

import (
	"github.com/gin-gonic/gin"
	co "github.com/inagacky/go_gin_api/src/api/controller"
	"go.uber.org/dig"
)

func GetRouting(c *dig.Container) *gin.Engine {

	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		// ユーザー関係のAPI
		u := v1.Group("/users")
		{
			if err := c.Invoke(func(controller co.UserController) {
				u.GET("/:id", controller.GetUser)
				u.POST("", controller.CreateUser)
				u.PUT("/:id", controller.UpdateUser)
				u.DELETE("/:id", controller.DeleteUser)
			}); err != nil {
				panic(err)
			}
		}
	}
	return r
}
