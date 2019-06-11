package comic_test

import (
	"testing"

	"github.com/bickyeric/arumba/mocks"
	"github.com/bickyeric/arumba/model"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	comicRepo := mocks.NewMockIComic(ctrl)
	assert.NotPanics(t, func() {
		comic.NewSearch(comicRepo)
	})
}

func TestPerform(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	comicRepo := mocks.NewMockIComic(ctrl)
	comicRepo.EXPECT().FindAll("One Piece").Return([]model.Comic{model.Comic{Name: "One Piece"}}, nil)

	searcher := comic.NewSearch(comicRepo)
	comics, err := searcher.Perform("One Piece")
	assert.Nil(t, err)
	assert.Len(t, comics, 1)
}