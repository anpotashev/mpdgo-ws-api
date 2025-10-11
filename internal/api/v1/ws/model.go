package ws

import (
	"encoding/json"

	"github.com/google/uuid"
)

type payloadType string

// Константы для поля PayloadType
const (
	subscribe          payloadType = "subscribe"
	setConnectionState payloadType = "connection/set"
	getConnectionState payloadType = "connection/get"

	// outputs
	listOutputs payloadType = "output/list"
	setOutput   payloadType = "output/set"

	// current playlist
	listCurrentPlaylist            payloadType = "current_playlist/get"
	clearCurrentPlaylist           payloadType = "current_playlist/clear"
	addToPlaylist                  payloadType = "current_playlist/add"
	addToPosPlaylist               payloadType = "current_playlist/add_to_pos"
	deleteByPosFromCurrentPlaylist payloadType = "current_playlist/delete_by_pos"
	shuffleAllCurrentPlaylist      payloadType = "current_playlist/shuffle_all"
	shuffleCurrentPlaylist         payloadType = "current_playlist/shuffle"
	moveInCurrentPlaylist          payloadType = "current_playlist/move"
	batchMoveInCurrentPlaylist     payloadType = "current_playlist/batch_move"

	// tree
	getTree    payloadType = "tree/get"
	updateTree payloadType = "tree/update"

	// playback
	play     payloadType = "playback/play"
	pause    payloadType = "playback/pause"
	stop     payloadType = "playback/stop"
	next     payloadType = "playback/next"
	previous payloadType = "playback/previous"
	playId   payloadType = "playback/play_id"
	playPos  payloadType = "playback/play_pos"
	seekPos  payloadType = "playback/seek_pos"

	// stored playlist
	getStoredPlaylists          payloadType = "stored_playlist/get_stored_playlists"
	deleteStoredPlaylist        payloadType = "stored_playlist/delete_stored_playlist"
	saveCurrentPlaylistAsStored payloadType = "stored_playlist/save_current_playlist_as_stored"
	renameStoredPlaylist        payloadType = "stored_playlist/rename_stored_playlist"

	// status | options
	getStatus  payloadType = "status/get"
	setRandom  payloadType = "status/set_random"
	setRepeat  payloadType = "status/set_repeat"
	setSingle  payloadType = "status/set_single"
	setConsume payloadType = "status/set_consume"
)

// Базовый тип сообщения, получаемого по websocket
type WsMessage struct {
	PayloadType payloadType     `json:"@type"`
	RequestId   *uuid.UUID      `json:"requestId,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
}

type ErrorPayload struct {
	Message string `json:"message"`
}

type WsResponse struct {
	PayloadType payloadType `json:"@type,omitempty"`
	RequestId   *uuid.UUID  `json:"requestId,omitempty"`
	Error       *string     `json:"error,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
}
