package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type SetOutputRq struct {
	Id     int  `json:"id"`
	Enable bool `json:"enable"`
}

func SetOutputHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SetOutputRq
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		if rq.Enable {
			return api.WithRequestContext(ctx).EnableOutput(rq.Id)
		}
		return api.WithRequestContext(ctx).DisableOutput(rq.Id)
	}
}
