package route

import (
	"github.com/PsaTyrE/dbe_adzan/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	e.POST("/login", controller.Login)
	e.POST("/register", controller.Register)
	e.GET("/logout", controller.Logout)

	api := e.Group("/api")
	api.GET("/adzan", controller.Index)
	api.GET("/adzan/kota/:id", controller.GetAdzanByKota)
	api.GET("/kota", controller.IndexKota)
	api.GET("/kota/:id", controller.ShowKotaById)
	api.Use(middleware.JWT([]byte("secret"))) // Replace with your JWT secret key
	api.POST("/adzan", controller.Create)
	api.PUT("/adzan/:id", controller.Update)
	api.DELETE("/adzan", controller.Delete)

	api.POST("/kota", controller.CreateKota)
	api.PUT("/kota/:id", controller.UpdateKota)
	api.DELETE("/kota", controller.DeleteKota)

	e.Logger.Fatal(e.Start(":8000"))
}
