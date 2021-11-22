package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func (r *SongRepository) Create(s *models.Song) error {
	return r.db.QueryRow(`
	INSERT INTO 
	songs(
		title, author_id, url
	)
	VALUES (
		$1, $2, $3
	) RETURNING id;
	`, s.Title, s.Author.ID, s.URL).Scan(&s.ID)
}

func (r *SongRepository) FindByID(id int) (*models.Song, error) {
	rows, err := r.db.Query(`
	SELECT id, title, author_id, url
	FROM songs
	WHERE id = $1;
	`, id)
	if err != nil {
		return nil, err
	}

	var result *models.Song
	if rows.Next() {
		rows.Scan(
			&result.ID,
			&result.Title,
			&result.URL,
		)
	}
	result.Author.ID = id
	return result, nil
}

func (song *SongRepository) GetFromAtoB() {

}
