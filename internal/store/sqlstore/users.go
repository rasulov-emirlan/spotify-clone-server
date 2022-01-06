package sqlstore

import (
	"database/sql"
	"spotify-clone/server/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u *models.User) error {
	err := r.db.QueryRow(
		`
		INSERT INTO users (username, full_name, birth_date, password, email)
		VALUES($1, $2, $3, $4, $5) returning id;
		`,
		u.UserName,
		u.FullName,
		u.BirthDate,
		u.Password,
		u.Email,
	).Scan(&u.ID)
	u.Password = ""
	return err
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	SELECT id, username, password, email
	FROM users
		WHERE id=$1;
	`, id).Scan(&u.ID, &u.UserName, &u.Email)
	return &u, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`
	SELECT id, username, password, email
	FROM users
		WHERE email = $1;
	`, email).Scan(&u.ID, &u.UserName, &u.Password, &u.Email)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query(`
	SELECT r.name
	FROM users_roles ur
	INNER JOIN roles r
		ON ur.role_id = r.id
		WHERE ur.user_id = $1; 
	`, u.ID)
	if err != nil {
		return nil, err
	}
	var role string

	for rows.Next() {
		rows.Scan(
			&role,
		)
		u.Roles = append(u.Roles, role)
	}
	return &u, err
}

func (r *UserRepository) BanByID(id int) error {
	return nil
}

func (r *UserRepository) FindByName(name string) (*models.User, error) {
	var u models.User

	return &u, r.db.QueryRow(`
	SELECT id, username, profile_picture_url
	FROM users
		WHERE username @@ to_tsquery('english', $1)
	LIMIT 1;
	`, name).Scan(&u.ID, &u.UserName, &u.ProfilePictureURL)
}

// CODE that goes after is all about favorite songs,
// favorite authors and denied songs
// it does not affect the users table itself

func (r *UserRepository) AddFavoriteSong(songID, userID int) error {
	return r.db.QueryRow(`
	INSERT INTO favorite_songs(
		song_id, user_id)
		VALUES ($1, $2);
	`, songID, userID).Err()
}

func (r *UserRepository) ListFavoriteSongs(userID, limit, offset int) ([]*models.Song, error) {
	rows, err := r.db.Query(`
	SELECT s.id, s.name, s.cover_picture_url, s.song_url, u.id as user_id, u.username
	FROM favorite_songs fs
	INNER JOIN songs s
		ON s.id = fs.song_id
	INNER JOIN users u
		ON u.id = fs.user_id
	WHERE fs.user_id = $1
	LIMIT $2 OFFSET $3; 
	`, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	var (
		songs                                          []*models.Song
		songID, authorID                               int
		songName, authorName, coverPictureURL, songURL string
	)

	for rows.Next() {
		if err := rows.Scan(
			&songID,
			&songName,
			&coverPictureURL,
			&songURL,
			&authorID,
			&authorName,
			// &authorPFPURL,
		); err != nil {
			return nil, err
		}
		songs = append(songs, &models.Song{
			ID:       songID,
			Name:     songName,
			CoverURL: coverPictureURL,
			URL:      songURL,
			Author: models.User{
				ID:       authorID,
				UserName: authorName,
				// ProfilePictureURL: authorPFPURL,
			},
		})
	}
	return songs, nil
}

func (r *UserRepository) RemoveFromFavoriteSongs(userID, songID int) error {
	return r.db.QueryRow(`
	DELETE FROM favorite_songs
	WHERE song_id = $1
	AND user_id = $2;
	`, songID, userID).Err()
}

func (r *UserRepository) AddFavoriteAuthor(userID, authorID int) error {
	return r.db.QueryRow(`
	INSERT INTO favorite_authors(
		user_id, author_id)
		VALUES ($1, $2);
	`, userID, authorID).Err()
}

func (r *UserRepository) ListFavoriteAuthors(userID, limit, offset int) ([]*models.User, error) {
	rows, err := r.db.Query(`
	SELECT u.id, u.username
	FROM favorite_authors fa
	INNER JOIN users u
		ON fa.author_id = u.id
	WHERE fa.user_id = $1`, userID)
	if err != nil {
		return nil, err
	}

	var (
		authors        []*models.User
		authorID       int
		authorUsername string
	)

	for rows.Next() {
		if err := rows.Scan(
			&authorID,
			&authorUsername,
		); err != nil {
			return nil, err
		}
		authors = append(authors, &models.User{
			ID:       authorID,
			UserName: authorUsername,
		})
	}
	return authors, nil
}

func (r *UserRepository) RemoveFromFavoriteAuthors(userID, authorID int) error {
	return r.db.QueryRow(`
	DELETE FROM favorite_authors
		WHERE	user_id = $1
		AND		author_id = $2;
	`, userID, authorID).Err()
}
