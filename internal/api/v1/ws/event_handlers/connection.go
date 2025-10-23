package event_handlers

import (
	"context"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type GetConnectionStateResponse struct {
	Connected bool `json:"connected"`
}

func GetConnectionState(api mpdapi.MpdApi) EventHandle {
	return func(ctx context.Context) (interface{}, error) {
		return GetConnectionStateResponse{Connected: api.WithRequestContext(ctx).IsConnected()}, nil
	}
}
