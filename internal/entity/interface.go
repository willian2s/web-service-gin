package entity

type AlbumRepositoryInterface interface {
	Save(album *Album) error
}
