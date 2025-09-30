package ws

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	log "github.com/anpotashev/mpd-ws-api/internal/logger"
)

type ListPlaylistRequest struct {
}

func (req *ListPlaylistRequest) getPayloadType() payloadType {
	return listCurrentPlaylist
}

func (req *ListPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	playlist, err := mpdApi.WithRequestContext(ctx).Playlist()
	if err != nil {
		return nil, err
	}
	payload := dto.MapPlaylist(*playlist)
	return payload, nil
}

type ResetPlaylist struct{}

func (r *ResetPlaylist) getPayloadType() payloadType {
	return listCurrentPlaylist
}
func (r *ResetPlaylist) process(ctx context.Context) (interface{}, error) {
	return nil, nil
}

type AddToCurrentPlaylistRequest struct {
	Path string `json:"path"`
}

func (req *AddToCurrentPlaylistRequest) getPayloadType() payloadType {
	return addToPlaylist
}

func (req *AddToCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Add(req.Path)
	return nil, err
}

type ClearCurrentPlaylistRequest struct {
}

func (req *ClearCurrentPlaylistRequest) getPayloadType() payloadType {
	return clearCurrentPlaylist
}

func (req *ClearCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Clear()
	return nil, err
}

type AddToCurrentPlaylistToPosRequest struct {
	Path string `json:"path"`
	Pos  int    `json:"pos"`
}

func (req *AddToCurrentPlaylistToPosRequest) getPayloadType() payloadType {
	return addToPosPlaylist
}

func (req *AddToCurrentPlaylistToPosRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).AddToPos(req.Pos, req.Path)
	return nil, err
}

type DeleteByPosFromCurrentPlaylistRequest struct {
	Pos int `json:"pos"`
}

func (req *DeleteByPosFromCurrentPlaylistRequest) getPayloadType() payloadType {
	return deleteByPosFromCurrentPlaylist
}

func (req *DeleteByPosFromCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).DeleteByPos(req.Pos)
	return nil, err
}

type MoveInCurrentPlaylistRequest struct {
	FromPos int `json:"from_pos"`
	ToPos   int `json:"to_pos"`
}

func (req *MoveInCurrentPlaylistRequest) getPayloadType() payloadType {
	return moveInCurrentPlaylist
}

func (req *MoveInCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	log.Info("moving ", "from", req.FromPos, "to", req.ToPos)
	err := mpdApi.WithRequestContext(ctx).Move(req.FromPos, req.ToPos)
	return nil, err
}

type BatchMoveInCurrentPlaylistRequest struct {
	FromStartPos int `json:"from_start_pos"`
	FromEndPos   int `json:"from_end_pos"`
	ToPos        int `json:"to_pos"`
}

func (req *BatchMoveInCurrentPlaylistRequest) getPayloadType() payloadType {
	return batchMoveInCurrentPlaylist
}

func (req *BatchMoveInCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).BatchMove(req.FromStartPos, req.FromEndPos, req.ToPos)
	return nil, err
}

type ShuffleAllInCurrentPlaylistRequest struct {
}

func (req *ShuffleAllInCurrentPlaylistRequest) getPayloadType() payloadType {
	return shuffleAllCurrentPlaylist
}

func (req *ShuffleAllInCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).ShuffleAll()
	return nil, err
}

type ShuffleInCurrentPlaylistRequest struct {
	FromPos int `json:"from_pos"`
	ToPos   int `json:"to_pos"`
}

func (req *ShuffleInCurrentPlaylistRequest) getPayloadType() payloadType {
	return shuffleCurrentPlaylist
}

func (req *ShuffleInCurrentPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).Shuffle(req.FromPos, req.ToPos)
	return nil, err
}
