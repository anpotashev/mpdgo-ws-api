package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/anpotashev/mpdgo/pkg/mpdapi"
	"github.com/gorilla/websocket"
)

var mpdApi mpdapi.MpdApi

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Интерфейс, имплементациями которого являются указатели на структуру payload запросов
// эти же имплементации могут использовать и для подготовки данных, отправляемых при получении idle событий.
type processable interface {
	process(ctx context.Context) (interface{}, error)
	getPayloadType() payloadType
}

// открытое websocket соединение
type client struct {
	conn   *websocket.Conn
	send   chan WsResponse
	topics map[mpdapi.MpdEventType]bool
	ctx    context.Context
}

// глобальный объект, хранящий открытие websocket соединения
var hub = struct {
	clients map[*client]struct{}
	sync.Mutex
}{
	clients: make(map[*client]struct{}),
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	cl := client{
		conn: conn,
		send: make(chan WsResponse, 100),
		ctx:  r.Context(),
	}
	hub.clients[&cl] = struct{}{}
	go cl.writePump()
	cl.sendCurrentMpdData()
	cl.readPump()
}

func (cl *client) writePump() {
	for msg := range cl.send {
		cl.conn.WriteJSON(msg)
	}
}

func (cl *client) readPump() {
	defer func() {
		delete(hub.clients, cl)
		close(cl.send)
		cl.conn.Close()
	}()
	for {
		var msg WsMessage
		err := cl.conn.ReadJSON(&msg)
		if err != nil {
			errMsg := fmt.Errorf("error parsing json: %w", err).Error()
			response := WsResponse{
				Error: &errMsg,
			}
			cl.send <- response
		}
		cl.processMsg(&msg)
	}
}

// Мапа используемая при десереализации payload-ов запросов. Ключ - значение из поля payloadType, значение - фабрика указателей на структуру payload
var typeToPayload = map[payloadType]func() processable{
	subscribe: func() processable { return &SubscribeRqPayload{} },
	// connection
	setConnectionState: func() processable { return &SetConnectionStateRqPayload{} },
	getConnectionState: func() processable { return &GetConnectionStateRequest{} },
	// outputs
	listOutputs: func() processable { return &ListOutputRequest{} },
	setOutput:   func() processable { return &SetOutputRq{} },
	// current playlist
	listCurrentPlaylist:            func() processable { return &ListPlaylistRequest{} },
	clearCurrentPlaylist:           func() processable { return &ClearCurrentPlaylistRequest{} },
	addToPlaylist:                  func() processable { return &AddToCurrentPlaylistRequest{} },
	addToPosPlaylist:               func() processable { return &AddToCurrentPlaylistToPosRequest{} },
	deleteByPosFromCurrentPlaylist: func() processable { return &DeleteByPosFromCurrentPlaylistRequest{} },
	shuffleAllCurrentPlaylist:      func() processable { return &ShuffleAllInCurrentPlaylistRequest{} },
	shuffleCurrentPlaylist:         func() processable { return &ShuffleInCurrentPlaylistRequest{} },
	moveInCurrentPlaylist:          func() processable { return &MoveInCurrentPlaylistRequest{} },
	batchMoveInCurrentPlaylist:     func() processable { return &BatchMoveInCurrentPlaylistRequest{} },
	// tree
	getTree: func() processable { return &GetTreeRequest{} },
	// playback
	play:     func() processable { return &PlayRequest{} },
	pause:    func() processable { return &PauseRequest{} },
	stop:     func() processable { return &StopRequest{} },
	next:     func() processable { return &NextRequest{} },
	previous: func() processable { return &PreviousRequest{} },
	playId:   func() processable { return &PlayIdRequest{} },
	playPos:  func() processable { return &PlayPosRequest{} },
	seekPos:  func() processable { return &SeekPosRequest{} },
	// stored playlist
	getStoredPlaylists:          func() processable { return &GetStoredPlaylistsRequest{} },
	deleteStoredPlaylist:        func() processable { return &DeleteStoredPlaylistRequest{} },
	saveCurrentPlaylistAsStored: func() processable { return &SaveCurrentPlaylistAsStoredRequest{} },
	renameStoredPlaylist:        func() processable { return &RenameStoredPlaylistRequest{} },
	// status | options
	getStatus:  func() processable { return &GetStatusRequest{} },
	setRandom:  func() processable { return &SetRandomStateRequest{} },
	setRepeat:  func() processable { return &SetRepeatStateRequest{} },
	setSingle:  func() processable { return &SetSingleStateRequest{} },
	setConsume: func() processable { return &SetConsumeStateRequest{} },
}

func (cl *client) processMsg(msg *WsMessage) {
	//lint:ignore SA1029 ignore
	ctx := context.WithValue(context.Background(), "requestId", msg.RequestId.String())
	var err error
	var payload processable
	var typeFound = false
	for k, v := range typeToPayload {
		if msg.PayloadType == k {
			typeFound = true
			payload = v()
			err := json.Unmarshal(msg.Payload, payload)
			if err != nil {
				cl.sendError(msg, err)
				return
			}
			break
		}
	}
	if !typeFound {
		cl.sendError(msg, fmt.Errorf("unknown payload type %s", msg.PayloadType))
		return
	}
	subscribeRqPayload, ok := payload.(*SubscribeRqPayload)
	if ok {
		err = cl.subscribe(subscribeRqPayload.Topics)
		if err != nil {
			cl.sendError(msg, err)
			return
		}
	}
	result, err := payload.process(ctx)
	if err != nil {
		cl.sendError(msg, err)
		return
	}
	response := WsResponse{
		PayloadType: msg.PayloadType,
		RequestId:   msg.RequestId,
		Error:       nil,
		Payload:     result,
	}
	cl.send <- response
}

func (cl *client) sendError(msg *WsMessage, err error) {
	errMsg := err.Error()
	response := WsResponse{
		RequestId: msg.RequestId,
		Error:     &errMsg,
	}
	cl.send <- response
}

func Init(api mpdapi.MpdApi) http.HandlerFunc {
	mpdApi = api
	go listenMpdEvent()
	return wsHandler
}
