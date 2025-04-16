package routes

import (
	"github.com/gin-gonic/gin"
	"web_go/internal/controllers"
	"web_go/internal/repositories"
	"web_go/internal/services"
	"web_go/pkg/jwt"
	"web_go/pkg/mysql"
)

func registerAPI(router *gin.Engine) {
	// 初始化各层依赖
	db := mysql.GetDB()

	// 版本化路由组‌
	v1 := router.Group("/api/v1")
	v1.Use(jwt.AuthMiddleware())
	{
		userRepository := repositories.NewUserRepository(db)
		userService := services.NewUserService(userRepository)
		userController := controllers.NewUserController(userService)
		users := v1.Group("/user")
		{
			/*{
				api.GET("/users", controllers.GetUsers) // 用户列表接口
				api.GET("/user/:id", controllers.GetUser) // 用户详情接口
				api.POST("/user", controllers.CreateUser) // 创建用户接口
				api.PUT("/user/:id", controllers.UpdateUser) // 更新用户接口
				api.DELETE("/user/:id", controllers.DeleteUser) // 删除用户接口
			}*/
			users.POST("/register", userController.CreateCtrl) // 用户注册接口
			users.POST("/login", userController.SignInCtrl)    // 用户登录接口
			users.GET("/:id", userController.GetCtrl)          // 用户详情接口
		}
	}

}
