package words

import (
	"database/sql"
	"fmt"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) RGetWordById(id int) (*Word, error) {
	var word Word
	err := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id=$1`, id).
		Scan(&word.Id, &word.Title, &word.Translation)
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) RDeleteWord(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) RUpdateWord(title, translate string, id int) error {
	_, err := r.db.Exec(`UPDATE ru_en SET title=$1, translation=$2 WHERE id=$3 `, title, translate, id)
	if err != nil {
		return err
	}

	return nil
}

// Функция умного поиска слов
func (r *Repo) RSmartSearch(title string) ([]Word, error) {
	rows, err := r.db.Query(`SELECT id, title, translation AS similarity FROM ru_en WHERE word_similarity(title, $1)>0.1 ORDER BY similarity DESC LIMIT 100`, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []Word
	for rows.Next() {
		var w Word
		if err := rows.Scan(&w.Id, &w.Title, &w.Translation); err != nil {
			return nil, err
		}
		words = append(words, w)
	}
	if len(words) == 0 {
		return nil, fmt.Errorf("no similar words found")
	}
	return words, nil

}
