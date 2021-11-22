package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func (r *SongRepository) Create(s *models.Song) (int, error) {
	row, err := r.db.Query(`
	INSERT INTO 
	songs(
		title, author_id, url
	)
	VALUES (
		$1, $2, $3
	) returning id;
	`, s.Title, s.Author.ID, s.URL)

	var id int

	if row.Next() {
		row.Scan(&id)
	}
	return id, err
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
	return result, nil
}

func (song *SongRepository) GetFromAtoB() {

}
