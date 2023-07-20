package entity

type AlbumRepositoryInterface interface {
	Save(album *Album) error
	FindMany() ([]Album, error)
	FindOne(id string) (*Album, error)
	Delete(id string) error
}
