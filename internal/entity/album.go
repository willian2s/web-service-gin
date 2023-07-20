package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func NewAlbum(title string, artist string, price float64) (*Album, error) {
	album := &Album{
		ID:     uuid.New().String(),
		Title:  title,
		Artist: artist,
		Price:  price,
	}

	err := album.Validate()
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (a *Album) Validate() error {
	if a.Title == "" {
		return errors.New("title is required")
	}
	if a.Artist == "" {
		return errors.New("artist is required")
	}
	if a.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}
