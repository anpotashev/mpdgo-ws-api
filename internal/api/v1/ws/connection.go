package ws

import "context"

type SetConnectionStateRqPayload struct {
	Enable bool `json:"enable"`
}

func (req *SetConnectionStateRqPayload) getPayloadType() payloadType {
	return setConnectionState
}

func (req *SetConnectionStateRqPayload) process(ctx context.Context) (interface{}, error) {
	var err error
	if req.Enable {
		err = mpdApi.WithRequestContext(ctx).Connect()
	} else {
		err = mpdApi.WithRequestContext(ctx).Disconnect()
	}
	return nil, err
}

type GetConnectionStateRequest struct{}

func (req *GetConnectionStateRequest) getPayloadType() payloadType {
	return getConnectionState
}

type GetConnectionStateResponse struct {
	Connected bool `json:"connected"`
}

func (req *GetConnectionStateRequest) process(ctx context.Context) (interface{}, error) {
	connected := mpdApi.WithRequestContext(ctx).IsConnected()
	return GetConnectionStateResponse{Connected: connected}, nil
}
