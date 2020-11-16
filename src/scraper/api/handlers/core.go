package handler

import (
	"net/http"

	"github.com/L04DB4L4NC3R/spotify-downloader/scraper/api/views"
	"github.com/L04DB4L4NC3R/spotify-downloader/scraper/core"
	"github.com/gorilla/mux"
)

type handler struct {
	router  *mux.Router
	service core.Service
}

func (h *handler) ViewAlbumMeta() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) ViewSongMeta() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) ViewProgressOfAlbumDownload() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) PauseAlbumDownload() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) ResumeAlbumDownload() http.Handler {
	panic("not implemented") // TODO: Implement
}

// player
func (h *handler) PlayPauseSong() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) DownloadSong() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) DownloadAlbum() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) SyncAlbum() http.Handler {
	panic("not implemented") // TODO: Implement
}

func (h *handler) Health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		views.Fill(w, "pong", nil, http.StatusOK)
	})
}

func NewHandler(r *mux.Router, svc core.Service) Handler {
	return &handler{
		router:  r,
		service: svc,
	}
}