package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type UpdateTreeRequest struct {
	Path string `json:"path"`
}

func UpdateTreeHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq UpdateTreeRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).UpdateDB(rq.Path)
	}
}
