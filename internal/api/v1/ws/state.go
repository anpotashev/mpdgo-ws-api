package ws

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
)

type GetStatusRequest struct{}

func (req *GetStatusRequest) getPayloadType() payloadType {
	return getStatus
}

func (req *GetStatusRequest) process(ctx context.Context) (interface{}, error) {
	status, err := mpdApi.WithRequestContext(ctx).Status()
	if err != nil {
		return nil, err
	}
	payload := dto.MapStatus(status)
	return payload, nil
}

type ResetStatus struct{}

func (r *ResetStatus) getPayloadType() payloadType {
	return getStatus
}
func (r *ResetStatus) process(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type SetRandomStateRequest struct {
	Enable bool `json:"enable"`
}

func (req *SetRandomStateRequest) getPayloadType() payloadType {
	return setRandom
}

func (req *SetRandomStateRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Random(req.Enable)
	return nil, err
}

type SetRepeatStateRequest struct {
	Enable bool `json:"enable"`
}

func (req *SetRepeatStateRequest) getPayloadType() payloadType {
	return setRepeat
}

func (req *SetRepeatStateRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Repeat(req.Enable)
	return nil, err
}

type SetSingleStateRequest struct {
	Enable bool `json:"enable"`
}

func (req *SetSingleStateRequest) getPayloadType() payloadType {
	return setSingle
}

func (req *SetSingleStateRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Single(req.Enable)
	return nil, err
}

type SetConsumeStateRequest struct {
	Enable bool `json:"enable"`
}

func (req *SetConsumeStateRequest) getPayloadType() payloadType {
	return setConsume
}

func (req *SetConsumeStateRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Consume(req.Enable)
	return nil, err
}
