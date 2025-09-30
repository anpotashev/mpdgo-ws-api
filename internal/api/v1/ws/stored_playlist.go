package ws

import (
	"context"

	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
)

type GetStoredPlaylistsRequest struct {
	Name string `json:"name"`
}

func (req *GetStoredPlaylistsRequest) getPayloadType() payloadType {
	return getStoredPlaylists
}

type GetStoredPlaylistsResponse struct {
	playlists []dto.Playlist
}
type ResetStoredPlaylists struct{}

func (r *ResetStoredPlaylists) getPayloadType() payloadType {
	return getStoredPlaylists
}
func (r *ResetStoredPlaylists) process(ctx context.Context) (interface{}, error) {
	return nil, nil
}
func (req *GetStoredPlaylistsRequest) process(ctx context.Context) (interface{}, error) {
	playlists, err := mpdApi.WithRequestContext(ctx).GetPlaylists()
	if err != nil {
		return nil, err
	}
	payload := dto.MapSlice(playlists, dto.MapPlaylist)
	return GetStoredPlaylistsResponse{playlists: payload}, nil
}

type DeleteStoredPlaylistRequest struct {
	Name string `json:"name"`
}

func (req DeleteStoredPlaylistRequest) getPayloadType() payloadType {
	return deleteStoredPlaylist
}

func (req DeleteStoredPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).DeleteStoredPlaylist(req.Name)
	return nil, err
}

type SaveCurrentPlaylistAsStoredRequest struct {
	Name string `json:"name"`
}

func (req SaveCurrentPlaylistAsStoredRequest) getPayloadType() payloadType {
	return saveCurrentPlaylistAsStored
}

func (req SaveCurrentPlaylistAsStoredRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).SaveCurrentPlaylistAsStored(req.Name)
	return nil, err
}

type RenameStoredPlaylistRequest struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

func (req RenameStoredPlaylistRequest) getPayloadType() payloadType {
	return renameStoredPlaylist
}

func (req RenameStoredPlaylistRequest) process(ctx context.Context) (interface{}, error) {
	err := mpdApi.WithRequestContext(ctx).RenameStoredPlaylist(req.OldName, req.NewName)
	return nil, err
}
