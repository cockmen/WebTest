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
	api.DELETE("/word/delete/:id", svc.DeleteWord)
	api.PUT("/word/update/:id", svc.UpdateWord)
	api.GET("/report/:id", svc.GetReportById)
	api.POST("/reports", svc.CreateNewReport)
	api.DELETE("/report/delete/:id", svc.DeleteReport)
	api.PUT("/report/update/:id", svc.UpdateReport)
	api.GET("/search/ru", svc.SmartSearch)

	router.Logger.Fatal(router.Start(":1323"))
}
