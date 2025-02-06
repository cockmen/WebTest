package reports

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) RGetReportById(id int) (*Report, error) {
	var report Report
	err := r.db.QueryRow(`SELECT id, title, descriptions, created_at, updated_at FROM reports WHERE id=$1`, id).
		Scan(&report.Id, &report.Title, &report.Description, &report.Created, &report.Updated)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (r *Repo) RCreateNewReport(title, description string) error {
	_, err := r.db.Exec(`INSERT INTO reports (title, descriptions, created_at, updated_at) VALUES ($1, $2, $3, $4)`, title, description, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) RDeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) RUpdateReport(title, description string, id int) error {
	_, err := r.db.Exec(`UPDATE reports SET title = $1, descriptions = $2, updated_at = $4 WHERE id =$3 `, title, description, id, time.Now())
	if err != nil {
		return err
	}
	return nil
}
