package route

import (
	"gin-scaffold/controller"
	_ "gin-scaffold/docs"
	"gin-scaffold/middleware"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	TIME_DURATION = 10
)

func DefinitionRoute(router *gin.Engine) {
	// set run mode
	gin.SetMode(gin.DebugMode)
	// middleware
	router.Use(middleware.UseCookieSession())
	router.Use(middleware.TimeoutHandler(time.Second * TIME_DURATION))
	// route
	// user
	var userController *controller.UserController
	router.GET("/users", userController.GetAllUsers)
	router.GET("/user/search", userController.SearchUsersByKeys)
	router.POST("/user/create", userController.CreateUser)
	router.POST("/user/update", userController.UpdateUser)
	router.POST("/user/update/usertype", userController.UpdateUserByUserType)
	router.POST("/user/delete", userController.DeleteUser)

	// api doc
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "USE_SWAGGER"))

}
