package main

import (
	"OssetianServer/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	db, err := NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	server := echo.New()

	srv := service.NewService(db)

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))

	api := server.Group("/api")
	{
		api.GET("/translate", srv.OssetianTranslate)
		api.GET("/origin", srv.RussianTranslate)
	}

	server.Logger.Fatal(server.Start(":8080"))
}
