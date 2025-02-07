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

	//Words
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.DELETE("/word/:id", svc.DeleteWord)
	api.PUT("/word/:id", svc.UpdateWord)

	//Reports
	api.GET("/report/:id", svc.GetReportById)
	api.POST("/reports", svc.CreateNewReport)
	api.DELETE("/report/:id", svc.DeleteReport)
	api.PUT("/report/:id", svc.UpdateReport)

	//Search
	api.GET("/search/ru", svc.SmartSearch)

	router.Logger.Fatal(router.Start(":1323"))
}
