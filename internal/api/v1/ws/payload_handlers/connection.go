package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type SetConnectionStateRqPayload struct {
	Enable bool `json:"enable"`
}

func SetConnectionStateHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SetConnectionStateRqPayload
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		if rq.Enable {
			return api.WithRequestContext(ctx).Connect()
		}
		return api.WithRequestContext(ctx).Disconnect()
	}
}
