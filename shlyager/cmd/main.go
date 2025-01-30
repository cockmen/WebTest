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
	api.DELETE("/remove/:id", svc.DeleteWord)
	api.PUT("/update/:id", svc.UpdateWord)
	api.GET("/report/:id", svc.GetReportById)
	api.POST("/reports", svc.CreareNewReport)
	api.DELETE("/remove_report/:id", svc.DeleteReport)
	api.PUT("/update_report/:id", svc.UpdateReport)

	router.Logger.Fatal(router.Start(":1323"))
}
