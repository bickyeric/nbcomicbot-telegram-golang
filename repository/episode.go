package repository

import (
	"database/sql"
	"fmt"

	"github.com/bickyeric/arumba/model"
)

// IEpisode ...
type IEpisode interface {
	Count(comicID int) (int, error)
	No(comicID, offset int) (float64, error)
	FindByNo(comicID int, no float64) (*model.Episode, error)
	GetLink(episodeID, sourceID int) (string, error)
	GetSources(episodeID int) []int
	Insert(*model.Episode) error
	InsertLink(episodeID, sourceID int, link string) error
}

type episodeRepository struct {
	*sql.DB
}

// NewEpisode ...
func NewEpisode(db *sql.DB) IEpisode {
	return episodeRepository{db}
}

func (repo episodeRepository) Count(comicID int) (int, error) {
	totalEpisode := 0
	row := repo.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM episodes WHERE comic_id=%d`, comicID))
	err := row.Scan(&totalEpisode)
	return totalEpisode, err
}

func (repo episodeRepository) No(comicID, offset int) (float64, error) {
	var (
		no float64 = 0.0
	)

	row := repo.QueryRow(fmt.Sprintf(`SELECT no FROM episodes WHERE comic_id=%d LIMIT %d,1`, comicID, offset))
	err := row.Scan(&no)
	return no, err
}

func (repo episodeRepository) InsertLink(episodeID, sourceID int, link string) error {
	_, err := repo.Exec("INSERT INTO episode_source(source_id, episode_id, link) VALUES(?,?,?)", sourceID, episodeID, link)
	return err
}

func (repo episodeRepository) GetLink(episodeID, sourceID int) (string, error) {
	link := ""
	row := repo.QueryRow("SELECT link FROM episode_source WHERE source_id=? AND episode_id=?", sourceID, episodeID)
	err := row.Scan(&link)
	return link, err
}

func (repo episodeRepository) Insert(episode *model.Episode) error {
	res, err := repo.Exec("INSERT INTO episodes(no, name, created_at, updated_at, comic_id) VALUES(?,?,?,?,?)", episode.No, episode.Name, episode.CreatedAt, episode.UpdatedAt, episode.ComicID)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	episode.ID = int(id)
	return nil
}

func (repo episodeRepository) FindByNo(comicID int, no float64) (*model.Episode, error) {
	episode := new(model.Episode)
	row := repo.QueryRow("SELECT * FROM episodes WHERE comic_id=? AND no=?", comicID, no)
	err := row.Scan(&episode.ID, &episode.No, &episode.Name, &episode.CreatedAt, &episode.UpdatedAt, &episode.ComicID)
	return episode, err
}

func (repo episodeRepository) GetSources(episodeID int) []int {
	sourceIds := []int{}
	rows, err := repo.Query("SELECT source_id FROM episode_source WHERE episode_id=?", episodeID)
	if err != nil {
		return sourceIds
	}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		sourceIds = append(sourceIds, id)
	}
	return sourceIds
}
