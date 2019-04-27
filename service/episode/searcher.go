package episode

import (
	"github.com/bickyeric/arumba/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Search ...
type Search struct {
	Repo repository.IEpisode
}

// Perform ...
func (s Search) Perform(comicID primitive.ObjectID, bound ...float64) ([][]float64, error) {
	totalEpisode, err := s.Repo.Count(comicID, bound...)
	if err != nil {
		return nil, err
	}
	noGroup := [][]float64{}

	number := 5
	index := 0
	for i := 0; i < number; i++ {
		member := (totalEpisode - index) / (number - i)
		if member < 1 {
			continue
		}
		noRange := []float64{}
		no, err := s.Repo.No(comicID, index, bound...)
		if err != nil {
			return nil, err
		}
		noRange = append(noRange, no)

		index += member
		if member > 1 {
			no, err := s.Repo.No(comicID, index-1, bound...)
			if err != nil {
				return nil, err
			}
			noRange = append(noRange, no)
		}
		noGroup = append(noGroup, noRange)
	}
	return noGroup, nil
}
