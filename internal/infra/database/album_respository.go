package database

import (
	"database/sql"
	"errors"

	"github.com/willian2s/web-service-gin/internal/entity"
)

type AlbumRepository struct {
	Db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{
		Db: db,
	}
}

func (repo *AlbumRepository) Save(album *entity.Album) error {
	_, err := repo.Db.Exec(`INSERT INTO albums (
			id,
			title,
			artist,
			price
			) VALUES (?, ?, ?, ?)
		`,
		album.ID,
		album.Title,
		album.Artist,
		album.Price,
	)

	if err != nil {
		return err
	}

	err = repo.Db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (repo *AlbumRepository) FindMany() ([]entity.Album, error) {
	var albums []entity.Album
	rows, err := repo.Db.Query("SELECT * FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album entity.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	err = repo.Db.Close()
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (repo *AlbumRepository) FindOne(id string) (*entity.Album, error) {
	var album entity.Album
	err := repo.Db.QueryRow("SELECT * FROM albums WHERE id = ?", id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err != nil {
		return nil, err
	}

	err = repo.Db.Close()
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func (repo *AlbumRepository) Delete(id string) error {
	var total int
	err := repo.Db.QueryRow("SELECT count(*) FROM albums WHERE id = ?", id).Scan(&total)
	if err != nil {
		return err
	}

	if total == 0 {
		return errors.New("album not found")
	}

	_, err = repo.Db.Exec("DELETE FROM albums WHERE id = ?", id)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	err = repo.Db.Close()
	if err != nil {
		return err
	}

	return nil
}
