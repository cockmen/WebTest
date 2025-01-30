package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Service) GetReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	report, err := repo.RGetReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: report})
}

func (s *Service) CreareNewReport(c echo.Context) error {
	var reportSlice []Report
	err := c.Bind(&reportSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	for _, report := range reportSlice {
		err = repo.RCreateNewReport(report.Title, report.Description)
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "OK")
}

func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	err = repo.RDeleteReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "Deleted")
}

func (s *Service) UpdateReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	var reportSlice []Report
	err = c.Bind(&reportSlice)
	repo := s.reportsRepo
	for _, report := range reportSlice {
		err = repo.RUpdateReport(report.Title, report.Description, id)
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, "Updated")

}
