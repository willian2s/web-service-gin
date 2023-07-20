package usecase

import (
	"fmt"

	"github.com/willian2s/web-service-gin/internal/entity"
)

type AlbumInputDTO struct {
	Title  string
	Artist string
	Price  float64
}

type AlbumOutputDTO struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type AlbumUsecase struct {
	AlbumRepository entity.AlbumRepositoryInterface
}

func (c *AlbumUsecase) Save(input *AlbumInputDTO) (*AlbumOutputDTO, error) {
	album, err := entity.NewAlbum(input.Title, input.Artist, input.Price)
	if err != nil {
		return nil, err
	}

	err = c.AlbumRepository.Save(album)
	if err != nil {
		return nil, err
	}

	return &AlbumOutputDTO{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}, nil
}

func (c *AlbumUsecase) FindMany() ([]entity.Album, error) {
	findAlbums := []entity.Album{}
	albums, err := c.AlbumRepository.FindMany()
	fmt.Println(albums)
	if err != nil {
		return nil, err
	}

	for _, album := range albums {
		findAlbums = append(findAlbums, entity.Album{
			ID:     album.ID,
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		})
	}

	return findAlbums, nil
}

func (c *AlbumUsecase) FindOne(id string) (*entity.Album, error) {
	album, err := c.AlbumRepository.FindOne(id)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (c *AlbumUsecase) Delete(id string) error {
	err := c.AlbumRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
