package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/Rei-Suzuki1729/spajam_stech_api/controllers"
)

func Initialize() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/posts", controllers.GetPosts)
	e.GET("/posts/:id", controllers.GetPostById)
	e.POST("/posts", controllers.CreatePost)
	e.PUT("/posts/:id", controllers.UpdatePost)
	e.DELETE("/posts/:id", controllers.DeletePost)

	e.GET("/posts/ranking", controllers.GetPostRanking)
	
	return e

}
