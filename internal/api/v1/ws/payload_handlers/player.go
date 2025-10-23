package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type PlayIdRequest struct {
	Id int `json:"id"`
}

func PlayIdHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq PlayIdRequest
		err := json.Unmarshal(payload, &rq)
		if err != nil {
			return err
		}
		return api.WithRequestContext(ctx).PlayId(rq.Id)
	}
}

type PlayPosRequest struct {
	Pos int `json:"pos"`
}

func PlayPosHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq PlayPosRequest
		err := json.Unmarshal(payload, &rq)
		if err != nil {
			return err
		}
		return api.WithRequestContext(ctx).PlayPos(rq.Pos)
	}
}

type SeekPosRequest struct {
	Pos      int `json:"pos"`
	SeekTime int `json:"seek_time"`
}

func SeekPosHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SeekPosRequest
		err := json.Unmarshal(payload, &rq)
		if err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Seek(rq.Pos, rq.SeekTime)
	}
}

func PrevHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).Previous()
	}
}

func PlayHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).Play()
	}
}

func PauseHandlrFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).Pause()
	}
}

func StopHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).Stop()
	}
}

func NextHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).Next()
	}
}
