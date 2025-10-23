package event_handlers

import (
	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
	"golang.org/x/net/context"
)

type GetStoredPlaylistsResponse struct {
	Playlists []dto.Playlist
}

func GetStoredPlaylistsEventHandleFunc(api mpdapi.MpdApi) EventHandle {
	return func(ctx context.Context) (interface{}, error) {
		playlists, err := api.WithRequestContext(ctx).GetPlaylists()
		if err != nil {
			return nil, err
		}
		playlistsDto := dto.MapSlice(playlists, dto.MapPlaylist)
		return GetStoredPlaylistsResponse{Playlists: playlistsDto}, nil
	}
}
