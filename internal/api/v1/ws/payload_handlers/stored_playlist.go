package payload_handlers

import (
	"context"
	"encoding/json"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type DeleteStoredPlaylistRequest struct {
	Name string `json:"name"`
}

func DeleteStoredPlaylistHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq DeleteStoredPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).DeleteStoredPlaylist(rq.Name)
	}
}

type SaveCurrentPlaylistAsStoredRequest struct {
	Name string `json:"name"`
}

func SaveCurrentPlaylistAsStoredHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq SaveCurrentPlaylistAsStoredRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).SaveCurrentPlaylistAsStored(rq.Name)
	}
}

type RenameStoredPlaylistRequest struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

func RenameStoredPlaylistHandleFunc(api mpdapi.MpdApi) WsPayloadHandleFunc {
	return func(ctx context.Context, payload json.RawMessage) error {
		var rq RenameStoredPlaylistRequest
		if err := json.Unmarshal(payload, &rq); err != nil {
			return err
		}
		return api.WithRequestContext(ctx).RenameStoredPlaylist(rq.OldName, rq.NewName)
	}
}
