package main

import (
	"shlyager/internal/service"
	"shlyager/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := logs.NewLogger(false)
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}
	svc := service.NewService(db, logger)

	router := echo.New()
	api := router.Group("api")

	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.DELETE("/delete/:id", svc.DeleteWords)
	api.PUT("/update/:id", svc.UpdateWords)

	router.Logger.Fatal(router.Start(":1323"))
}
