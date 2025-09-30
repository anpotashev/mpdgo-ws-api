package ws

import "context"

type PlayIdRequest struct {
	Id int `json:"id"`
}

func (p *PlayIdRequest) getPayloadType() payloadType {
	return playId
}

func (p *PlayIdRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).PlayId(p.Id)
	return nil, err
}

type PlayPosRequest struct {
	Pos int `json:"pos"`
}

func (p *PlayPosRequest) getPayloadType() payloadType {
	return playPos
}

func (p *PlayPosRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).PlayPos(p.Pos)
	return nil, err
}

type SeekPosRequest struct {
	Pos      int `json:"pos"`
	SeekTime int `json:"seek_time"`
}

func (s *SeekPosRequest) getPayloadType() payloadType {
	return seekPos
}

func (s *SeekPosRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Seek(s.Pos, s.SeekTime)
	return nil, err
}

type PlayRequest struct{}

func (p *PlayRequest) getPayloadType() payloadType {
	return play
}

func (p *PlayRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Play()
	return nil, err
}

type PauseRequest struct{}

func (p *PauseRequest) getPayloadType() payloadType {
	return pause
}

func (p *PauseRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Pause()
	return nil, err
}

type StopRequest struct{}

func (s *StopRequest) getPayloadType() payloadType {
	return stop
}

func (s *StopRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Stop()
	return nil, err
}

type NextRequest struct{}

func (n *NextRequest) getPayloadType() payloadType {
	return next
}

func (n *NextRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Next()
	return nil, err
}

type PreviousRequest struct{}

func (p *PreviousRequest) getPayloadType() payloadType {
	return previous
}

func (p *PreviousRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Previous()
	return nil, err
}
