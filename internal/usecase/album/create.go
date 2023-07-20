package usecase

import "github.com/willian2s/web-service-gin/internal/entity"

type AlbumInputDTO struct {
	Title  string
	Artist string
	Price  float64
}

type AlbumOutputDTO struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

type CreateAlbum struct {
	AlbumRepository entity.AlbumRepositoryInterface
}

func (c *CreateAlbum) Handle(input *AlbumInputDTO) (*AlbumOutputDTO, error) {
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
