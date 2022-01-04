package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type LanguageRepository struct {
	db *sql.DB
}

func (r *LanguageRepository) Create(l *models.Language) error {
	return r.db.QueryRow(`
	INSERT INTO languages(name)
	VALUES($1)
	RETURNING ID;
	`, l.Name).Scan(&l.ID)
}

func (r *LanguageRepository) ListAll() ([]*models.Language, error) {
	rows, err := r.db.Query(`
	SELECT id, name
	FROM languages;
	`)
	if err != nil {
		return nil, err
	}

	var (
		languages []*models.Language
		id        int
		name      string
	)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&name,
		); err != nil {
			return nil, err
		}
		languages = append(languages, &models.Language{
			ID:   id,
			Name: name,
		})
	}

	return languages, nil
}

func (r *LanguageRepository) Delete(languageID int) error {
	return r.db.QueryRow(`
	DELETE FROM languages
	WHERE id = $1;
	`, languageID).Err()
}
