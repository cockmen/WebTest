package service

import (
	"database/sql"
	"shlyager/internal/reports"
	"shlyager/internal/words"

	"github.com/labstack/echo/v4"
)

const (
	InvalidParams       = "invalid params"
	InternalServerError = "internal error"
)

type Service struct {
	db          *sql.DB
	logger      echo.Logger
	wordsRepo   *words.Repo
	reportsRepo *reports.Repo
}

func NewService(db *sql.DB, logger echo.Logger) *Service {
	svc := &Service{
		db:     db,
		logger: logger,
	}
	svc.initRepositories(db)

	return svc
}

func (s *Service) initRepositories(db *sql.DB) {
	s.wordsRepo = words.NewRepo(db)
	s.reportsRepo = reports.NewRepo(db)
}

type Response struct {
	Object       any    `json:"object,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (r *Response) Error() string {
	return r.ErrorMessage
}
func (s *Service) NewError(err string) (int, *Response) {
	return 400, &Response{ErrorMessage: err}
}
