package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func (song *SongRepository) Create(s *models.Song) error {
	_, err := song.db.Query(`
	INSERT INTO 
	songs(
		title, author_id, url
	)
	VALUES (
		$1, $2, $3
	);
	`, s.Title, s.Author.ID, s.URL)

	if err != nil {
		return err
	}

	return nil
}

func (song *SongRepository) FindByID(id int) (*models.Song, error) {
	rows, err := song.db.Query(`
	SELECT id, title, author_id, url
	FROM songs
	WHERE id = $1;
	`, id)
	if err != nil {
		return nil, err
	}

	var result models.Song
	if rows.Next() {
		rows.Scan(
			&result,
		)
	}
	return &result, nil
}

func (song *SongRepository) GetFromAtoB() {

}
