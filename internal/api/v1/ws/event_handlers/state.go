package event_handlers

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

func GetStatusEventHandleFunc(api mpdapi.MpdApi) EventHandle {
	return func(ctx context.Context) (interface{}, error) {
		status, err := api.WithRequestContext(ctx).Status()
		if err != nil {
			return nil, err
		}
		payload := dto.MapStatus(status)
		return payload, nil
	}
}
