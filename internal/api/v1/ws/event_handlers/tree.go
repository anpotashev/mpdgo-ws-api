package event_handlers

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

func GetTreeEventHandler(api mpdapi.MpdApi) EventHandle {
	return func(ctx context.Context) (interface{}, error) {
		tree, err := api.WithRequestContext(ctx).Tree()
		if err != nil {
			return nil, err
		}
		payload := dto.MapMpdTree(*tree)
		return payload, nil
	}
}
