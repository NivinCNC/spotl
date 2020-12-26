package handler

import (
	"github.com/L04DB4L4NC3R/spotify-downloader/scraper/api/middleware"
	"github.com/L04DB4L4NC3R/spotify-downloader/scraper/pkg/core"
	"github.com/gorilla/mux"
)

func RegisterHandler(r *mux.Router, svc core.Service) {
	coreHandler := NewHandler(r, svc)
	middleware.RegisterMiddlewares(r)

	r.Handle("/ping", coreHandler.Health())

	// queue downloads
	r.Handle("/song/{id}/", coreHandler.DownloadSong())
	r.Handle("/playlist/{id}/", coreHandler.DownloadPlaylist())

	// fetch just the metadata
	r.Handle("/meta/song/{id}/", coreHandler.ViewSongMeta())
	r.Handle("/meta/playlist/{id}/", coreHandler.ViewPlaylistMeta())
}
