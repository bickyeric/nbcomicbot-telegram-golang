package main

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/updater"
	"github.com/bickyeric/arumba/updater/source"
	"github.com/subosito/gotenv"
)

func main() {

	gotenv.Load(".env")

	bot := arumba.NewBot()
	db := connection.NewMysql()
	kendang := connection.NewKendang()

	app := arumba.New(db)
	updater := updater.NewRunner(
		bot,
		kendang,
		episode.UpdateSaver{app.ComicRepo, app.EpisodeRepo, app.PageRepo},
	)

	mangacan := source.Mangacan{}
	updater.Run(mangacan)
	// gocron.Every(1).Minute().Do(updater.Run, mangacan)

	// mangatail := source.Mangatail{}
	// updater.Run(mangatail)

	// <-gocron.Start()
}
