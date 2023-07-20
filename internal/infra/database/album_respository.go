package database

import (
	"errors"

	"github.com/google/uuid"
)

// type AlbumRepository struct {
// 	Db *sql.DB
// }

// func NewAlbumRepository(db *sql.DB) *AlbumRepository {
// 	return &AlbumRepository{Db: db}
// }

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albumsCollection = []Album{
	{ID: uuid.New().String(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: uuid.New().String(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: uuid.New().String(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: uuid.New().String(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// func (repo *AlbumRepository) Save(album *entity.Album) error {
// 	_, err := repo.Db.Exec("INSERT INTO albums (id, title, artist, price) VALUES (?, ?, ?, ?)", albumsCollection[0].ID, albumsCollection[0].Title, albumsCollection[0].Artist, albumsCollection[0].Price)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (a *Album) Save() {
	a.ID = uuid.New().String()

	albumsCollection = append(albumsCollection, *a)
}

func (a *Album) FindMany() []Album {
	return albumsCollection
}

func (a *Album) FindOne(id string) (*Album, error) {
	for i, b := range albumsCollection {
		if b.ID == id {
			return &albumsCollection[i], nil
		}
	}
	return nil, errors.New("album not found")
}

func (a *Album) Delete(id string) {
	for i, b := range albumsCollection {
		if b.ID == id {
			albumsCollection = append(albumsCollection[:i], albumsCollection[i+1:]...)
			return
		}
	}
}
