package ws

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
)

type SetOutputRq struct {
	Id     int  `json:"id"`
	Enable bool `json:"enable"`
}

func (o *SetOutputRq) getPayloadType() payloadType {
	return setOutput
}

func (o *SetOutputRq) process(ctx context.Context) (interface{}, error) {
	var err error
	if o.Enable {
		err = mpdApi.WithRequestContext(ctx).EnableOutput(o.Id)
	} else {
		err = mpdApi.WithRequestContext(ctx).DisableOutput(o.Id)
	}
	return nil, err
}

type ListOutputRequest struct{}

func (o *ListOutputRequest) getPayloadType() payloadType {
	return listOutputs
}

type ListOutputRs struct {
	Outputs []dto.Output `json:"outputs"`
}

func (o *ListOutputRequest) process(ctx context.Context) (interface{}, error) {
	outputs, err := mpdApi.WithRequestContext(ctx).ListOutputs()
	if err != nil {
		return nil, err
	}
	dtoOutputs := dto.MapSlice(outputs, dto.MapOutput)
	return ListOutputRs{Outputs: dtoOutputs}, nil
}

type ResetOutput struct{}

func (r *ResetOutput) getPayloadType() payloadType {
	return listOutputs
}
func (r *ResetOutput) process(ctx context.Context) (interface{}, error) {
	return nil, nil
}
