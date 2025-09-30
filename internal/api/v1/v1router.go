package v1

import (
	"github.com/anpotashev/mpd-ws-api/internal/api/v1/ws"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
	"github.com/gorilla/mux"
)

type v1Router struct {
	MpdApi mpdapi.MpdApi
}

func New(router *mux.Router, api mpdapi.MpdApi) {
	r := &v1Router{
		MpdApi: api,
	}
	wsHandler := ws.Init(api)
	r.newConnectionRouter(router.PathPrefix("/connection").Subrouter())
	r.newTreeRouter(router.PathPrefix("/tree").Subrouter())
	r.newOutputRouter(router.PathPrefix("/output").Subrouter())
	r.newCurrentPlaylistRouter(router.PathPrefix("/playlist/current").Subrouter())
	r.newPlayerRouter(router.PathPrefix("/player").Subrouter())
	r.newStatusRouter(router.PathPrefix("/status").Subrouter())
	r.newTreeRouter(router.PathPrefix("tree").Subrouter())
	r.newPlaylistRouter(router.PathPrefix("/playlist").Subrouter())
	router.HandleFunc("/ws", wsHandler)

	//router.HandleFunc("/playlists", r.getPlaylistsHandler)
	//router.HandleFunc("/playlistinfo", r.getPlaylistInfoHandler)
}
