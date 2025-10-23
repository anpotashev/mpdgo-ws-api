package ws

type responseType string
type requestType string

const (
	// requestType
	setConnectionState             requestType = "connection/set"
	setOutput                      requestType = "output/set"
	clearCurrentPlaylist           requestType = "current_playlist/clear"
	addToPlaylist                  requestType = "current_playlist/add"
	addToPosPlaylist               requestType = "current_playlist/add_to_pos"
	deleteByPosFromCurrentPlaylist requestType = "current_playlist/delete_by_pos"
	shuffleAllCurrentPlaylist      requestType = "current_playlist/shuffle_all"
	shuffleCurrentPlaylist         requestType = "current_playlist/shuffle"
	moveInCurrentPlaylist          requestType = "current_playlist/move"
	batchMoveInCurrentPlaylist     requestType = "current_playlist/batch_move"
	addStoredToPos                 requestType = "current_playlist/add_stored_to_pos"
	updateTree                     requestType = "tree/update"
	play                           requestType = "playback/play"
	pause                          requestType = "playback/pause"
	stop                           requestType = "playback/stop"
	next                           requestType = "playback/next"
	previous                       requestType = "playback/previous"
	playId                         requestType = "playback/play_id"
	playPos                        requestType = "playback/play_pos"
	seekPos                        requestType = "playback/seek_pos"
	deleteStoredPlaylist           requestType = "stored_playlist/delete_stored_playlist"
	saveCurrentPlaylistAsStored    requestType = "stored_playlist/save_current_playlist_as_stored"
	renameStoredPlaylist           requestType = "stored_playlist/rename_stored_playlist"
	setRandom                      requestType = "status/set_random"
	setRepeat                      requestType = "status/set_repeat"
	setSingle                      requestType = "status/set_single"
	setConsume                     requestType = "status/set_consume"
)
const (
	// responseType
	getConnectionState  responseType = "connection/get"
	listOutputs         responseType = "output/list"
	listCurrentPlaylist responseType = "current_playlist/get"
	getTree             responseType = "tree/get"
	getStoredPlaylists  responseType = "stored_playlist/get_stored_playlists"
	getStatus           responseType = "status/get"
)
