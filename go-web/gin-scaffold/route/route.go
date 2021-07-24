package route

import (
	"gin-scaffold/controller"
	"gin-scaffold/middleware"
	"net/http"
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
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware.Tracing())
	router.Use(middleware.UseCookieSession())
	router.Use(middleware.TimeoutHandler(time.Second * TIME_DURATION))
	// no route
	router.NoRoute(NoRouteResponse)
	// home
	var userController *controller.UserController
	router.Static("/web/assets", "./web/assets")
	router.StaticFS("/web/upload", http.Dir("/web/upload"))
	router.LoadHTMLGlob("web/*.tmpl")
	// login
	router.GET("/login", userController.Login)
	router.POST("/dologin", userController.DoLogin)

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddle())
	{
		auth.GET("/", userController.UserHome)
		auth.GET("/logout", userController.Logout)
		// user
		router.GET("/users", userController.GetAllUsers)
		auth.GET("/user/add", userController.AddUser) //web ui
		router.GET("/user/search", userController.SearchUsersByKeys)
		router.POST("/user/create", userController.CreateUser)
		router.POST("/user/update", userController.UpdateUser)
		router.POST("/user/delete", userController.DeleteUser)
	}
	// api doc
	router.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "USE_SWAGGER"))

}

// no route
func NoRouteResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":  404,
		"error": "oops, page not exists!",
	})
}
