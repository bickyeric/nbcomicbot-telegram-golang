package telegram

import (
	"github.com/bickyeric/arumba"
	"github.com/bickyeric/arumba/connection"
	"github.com/bickyeric/arumba/service/comic"
	"github.com/bickyeric/arumba/service/episode"
	"github.com/bickyeric/arumba/service/telegraph"
	"github.com/bickyeric/arumba/telegram/callback"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	log "github.com/sirupsen/logrus"
)

type callbackHandler map[string]callback.Handler

// NewCallbackHandler ...
func NewCallbackHandler(app arumba.Arumba, bot arumba.Bot, kendang connection.IKendang) callback.Handler {
	telegraphCreator := telegraph.NewCreatePage()
	handlers := callbackHandler{}
	handlers[callback.SelectComicCallback] = callback.NewSelectComic(
		bot, bot,
		episode.Search{
			Repo: app.EpisodeRepo,
		},
	)
	handlers[callback.SelectEpisodeCallback] = callback.NewSelectEpisode(
		bot, bot,
		episode.Search{
			Repo: app.EpisodeRepo,
		},
		comic.NewRead(app, kendang, telegraphCreator),
	)
	return handlers
}

func (handler callbackHandler) Handle(event *tgbotapi.CallbackQuery) {
	contextLog := log.WithFields(
		log.Fields{
			"data": event.Data,
		},
	)
	contextLog.Info("Handling callback")
	method, _ := callback.ExtractData(event.Data)
	h, ok := handler[method]

	if ok {
		h.Handle(event)
	} else {
		contextLog.Warn("command not found : " + method)
	}
}
