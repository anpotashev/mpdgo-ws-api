package event_handlers

import (
	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
	"golang.org/x/net/context"
)

func ListCurrentPlaylistEventHandleFunc(api mpdapi.MpdApi) EventHandle {
	return func(ctx context.Context) (interface{}, error) {
		playlist, err := api.WithRequestContext(ctx).Playlist()
		if err != nil {
			return nil, err
		}
		payload := dto.MapPlaylist(*playlist)
		return payload, nil
	}
}
