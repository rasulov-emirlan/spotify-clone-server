package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type CountryRepository struct {
	db *sql.DB
}

func (r *CountryRepository) Create(c *models.Country) error {
	return r.db.QueryRow(`
	INSERT INTO countries(name)
	VALUES($1)
	RETURNING id;
	`, c.Name).Scan(&c.ID)
}

func (r *CountryRepository) ListAll() ([]*models.Country, error) {
	rows, err := r.db.Query(`
	SELECT id, name
	FROM countries;
	`)
	if err != nil {
		return nil, err
	}

	var (
		countries []*models.Country
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
		countries = append(countries, &models.Country{
			ID:   id,
			Name: name,
		})
	}
	return countries, nil
}

func (r *CountryRepository) Delete(countryID int) error {
	return r.db.QueryRow(`
	DELETE FROM countries
	WHERE id = $1;
	`, countryID).Err()
}
