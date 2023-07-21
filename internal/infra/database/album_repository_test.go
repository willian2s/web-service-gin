package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/willian2s/web-service-gin/internal/entity"
)

type AlbumRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *AlbumRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec(
		`CREATE TABLE IF NOT EXISTS albums (
			id varchar(255) NOT NULL,
			title varchar(255) NOT NULL,
			artist varchar(255) NOT NULL,
			price float NOT NULL,
			PRIMARY KEY (id)
		)`,
	)
	suite.Db = db
}

func (suite *AlbumRepositoryTestSuite) TearDownSuite() {
	suite.Db.Exec("DROP TABLE albums")
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(AlbumRepositoryTestSuite))
}

func (suite *AlbumRepositoryTestSuite) TestSavingAlbum() {
	album, err := entity.NewAlbum("Title", "Artist", 1.0)
	suite.NoError(err)
	repo := NewAlbumRepository(suite.Db)
	err = repo.Save(album)
	suite.NoError(err)

	var albumResult entity.Album
	err = suite.Db.QueryRow(
		"SELECT id, title, artist, price FROM albums WHERE id = ?",
		album.ID,
	).Scan(
		&albumResult.ID,
		&albumResult.Title,
		&albumResult.Artist,
		&albumResult.Price,
	)

	suite.NoError(err)
	suite.Equal(album.ID, albumResult.ID)
	suite.Equal(album.Title, albumResult.Title)
	suite.Equal(album.Artist, albumResult.Artist)
	suite.Equal(album.Price, albumResult.Price)
}
