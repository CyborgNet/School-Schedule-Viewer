package server

import (
	"../config"
	"../controllers"
	"../middlewares"
	"github.com/gin-gonic/contrib/static"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	cookieStore := cookie.NewStore([]byte(config.GetCookiePassword()))

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("SESSION", cookieStore))

	router.Use(static.Serve("/assets/css", static.LocalFile("../assets/template/assets/css", true)))
	router.Use(static.Serve("/assets/js", static.LocalFile("../assets/template/assets/js", true)))
	router.Use(static.Serve("/assets/bootstrap/css", static.LocalFile("../assets/template/assets/bootstrap/css", true)))
	router.Use(static.Serve("/assets/bootstrap/js", static.LocalFile("../assets/template/assets/bootstrap/js", true)))
	router.LoadHTMLGlob("../assets/template/*.html")

	router.Use(middlewares.AuthStudent())

	router.GET("/login", controllers.ShowLoginPage)
	router.POST("/login", controllers.LoginUser)

	router.GET("/", controllers.ShowMainPage)
	router.GET("/index", func(c *gin.Context) {
		c.Redirect(302, "/")
	})

	return router
}
