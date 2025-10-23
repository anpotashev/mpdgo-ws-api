package event_handlers

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type ListOutputRs struct {
	Outputs []dto.Output `json:"outputs"`
}

func ListOutputsEventHandleFunc(api mpdapi.MpdApi) EventHandle {
	return func(ctx context.Context) (interface{}, error) {
		outputs, err := api.WithRequestContext(ctx).ListOutputs()
		if err != nil {
			return nil, err
		}
		payload := dto.MapSlice(outputs, dto.MapOutput)
		return ListOutputRs{Outputs: payload}, nil
	}
}
