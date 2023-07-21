package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGetAnErrorIfTitleIsBlack(t *testing.T) {
	album := Album{}
	assert.Error(t, album.Validate(), "title is required")
}

func TestIfGetAnErrorIfArtistIsBlack(t *testing.T) {
	album := Album{
		Title: "Title",
	}
	assert.Error(t, album.Validate(), "artist is required")
}

func TestIfGetAnErrorIfPriceIsBlank(t *testing.T) {
	album := Album{
		Title:  "Title",
		Artist: "Artist",
	}
	assert.Error(t, album.Validate(), "price must be greater than 0")
}

func TestIfGetAnErrorIfPriceIsNegative(t *testing.T) {
	album := Album{
		Title:  "Title",
		Artist: "Artist",
		Price:  -1,
	}
	assert.Error(t, album.Validate(), "price must be greater than 0")
}

func TestIfAllValidParams(t *testing.T) {
	album := Album{
		Title:  "Title",
		Artist: "Artist",
		Price:  1,
	}
	assert.NoError(t, album.Validate())
}
