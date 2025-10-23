package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type SetRandomStateRequest struct {
	Enable bool `json:"enable"`
}

func SetRandomHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SetRandomStateRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Random(rq.Enable)
	}
}

type SetRepeatStateRequest struct {
	Enable bool `json:"enable"`
}

func SetRepeatHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SetRepeatStateRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Repeat(rq.Enable)
	}
}

type SetSingleStateRequest struct {
	Enable bool `json:"enable"`
}

func SetSingleHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SetSingleStateRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Single(rq.Enable)
	}
}

type SetConsumeStateRequest struct {
	Enable bool `json:"enable"`
}

func SetConsumeHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SetConsumeStateRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Consume(rq.Enable)
	}
}
