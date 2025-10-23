package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type AddToCurrentPlaylistRequest struct {
	Path string `json:"path"`
}

func AddToCurrentPlaylistHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var req AddToCurrentPlaylistRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Add(req.Path)
	}
}

func ClearHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).Clear()
	}
}

type AddToCurrentPlaylistToPosRequest struct {
	Path string `json:"path"`
	Pos  int    `json:"pos"`
}

func AddToCurrentPlaylistToPosHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq AddToCurrentPlaylistToPosRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).AddToPos(rq.Pos, rq.Path)
	}
}

type DeleteByPosFromCurrentPlaylistRequest struct {
	Pos int `json:"pos"`
}

func DeleteFromCurrentPlaylistByPosHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq DeleteByPosFromCurrentPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).DeleteByPos(rq.Pos)
	}
}

type MoveInCurrentPlaylistRequest struct {
	FromPos int `json:"from_pos"`
	ToPos   int `json:"to_pos"`
}

func MoveInCurrentPlaylistHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq MoveInCurrentPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Move(rq.FromPos, rq.ToPos)
	}
}

type BatchMoveInCurrentPlaylistRequest struct {
	FromStartPos int `json:"from_start_pos"`
	FromEndPos   int `json:"from_end_pos"`
	ToPos        int `json:"to_pos"`
}

func BatchMoveInCurrentPlaylistHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq BatchMoveInCurrentPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).BatchMove(rq.FromStartPos, rq.FromEndPos, rq.ToPos)
	}
}

func ShuffleAllInCurrentPlaylistHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		return api.WithRequestContext(ctx).ShuffleAll()
	}
}

type ShuffleInCurrentPlaylistRequest struct {
	FromPos int `json:"from_pos"`
	ToPos   int `json:"to_pos"`
}

func ShuffleInCurrentPlaylistHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq ShuffleInCurrentPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).Shuffle(rq.FromPos, rq.ToPos)
	}
}

type AddStoredPlaylistToCurrentPlaylistRequest struct {
	Name string `json:"name"`
	Pos  int    `json:"pos"`
}

func AddStoredPlaylistToCurrentPlaylistToPosHandlerFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq AddStoredPlaylistToCurrentPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).AddStoredToPos(rq.Name, rq.Pos)
	}
}
