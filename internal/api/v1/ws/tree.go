package ws

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
)

type GetTreeRequest struct{}

func (t *GetTreeRequest) getPayloadType() payloadType {
	return getTree
}

func (t *GetTreeRequest) process(ctx context.Context) (interface{}, error) {
	tree, err := mpdApi.WithRequestContext(ctx).Tree()
	if err != nil {
		return nil, err
	}
	payload := dto.MapMpdTree(*tree)
	return payload, nil
}

type UpdateTree struct {
	Path string `json:"path"`
}

func (t *UpdateTree) getPayloadType() payloadType {
	return updateTree
}

func (t *UpdateTree) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).UpdateDB(t.Path)
	if err != nil {
		return nil, err
	}
	return nil, err
}

type ResetTree struct{}

func (r *ResetTree) getPayloadType() payloadType {
	return getTree
}
func (r *ResetTree) process(ctx context.Context) (interface{}, error) {
	return nil, nil
}
