package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/anpotashev/mpd-ws-api/internal/api/v1/ws/event_handlers"
	"github.com/anpotashev/mpd-ws-api/internal/api/v1/ws/payload_handlers"
	log "github.com/anpotashev/mpd-ws-api/internal/logger"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type MpdWS struct {
	ctx      context.Context
	upgrader websocket.Upgrader
	clients  map[*client]struct{}
	sync.Mutex
	payloadHandleFunc map[requestType]payload_handlers.WsPayloadHandleFunc
	api               mpdapi.MpdApi
	eventHandleFunc   map[mpdapi.MpdEventType]map[responseType]event_handlers.EventHandle
}

// открытое websocket соединение
type client struct {
	conn *websocket.Conn
	send chan interface{}
	ctx  context.Context
}

func (cl *client) writePump() {
	for msg := range cl.send {
		cl.conn.WriteJSON(msg)
	}
}

func (cl *client) sendError(err error, requestID *uuid.UUID) {
	errorMsg := err.Error()
	cl.send <- WsResponse{
		Error:     &errorMsg,
		RequestId: requestID,
	}
}

func NewWsHandler(api mpdapi.MpdApi) *MpdWS {
	mpdWS := &MpdWS{
		ctx: context.Background(),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		clients: make(map[*client]struct{}),
		api:     api,
	}
	mpdWS.fillPayloadHandlers()
	mpdWS.fillEventHandlers()
	mpdWS.subscribe()
	mpdWS.startUpdatingStatus()
	return mpdWS
}

func (m *MpdWS) fillPayloadHandlers() {
	m.payloadHandleFunc = map[requestType]payload_handlers.WsPayloadHandleFunc{
		setConnectionState:             payload_handlers.SetConnectionStateHandleFunc(m.api),
		setOutput:                      payload_handlers.SetOutputHandleFunc(m.api),
		play:                           payload_handlers.PlayHandleFunc(m.api),
		pause:                          payload_handlers.PauseHandlrFunc(m.api),
		stop:                           payload_handlers.StopHandleFunc(m.api),
		next:                           payload_handlers.NextHandleFunc(m.api),
		previous:                       payload_handlers.PrevHandleFunc(m.api),
		playId:                         payload_handlers.PlayIdHandleFunc(m.api),
		playPos:                        payload_handlers.PlayPosHandleFunc(m.api),
		clearCurrentPlaylist:           payload_handlers.ClearHandlerFunc(m.api),
		addToPlaylist:                  payload_handlers.AddToCurrentPlaylistHandlerFunc(m.api),
		addToPosPlaylist:               payload_handlers.AddToCurrentPlaylistToPosHandlerFunc(m.api),
		deleteByPosFromCurrentPlaylist: payload_handlers.DeleteFromCurrentPlaylistByPosHandlerFunc(m.api),
		shuffleAllCurrentPlaylist:      payload_handlers.ShuffleAllInCurrentPlaylistHandlerFunc(m.api),
		shuffleCurrentPlaylist:         payload_handlers.ShuffleInCurrentPlaylistHandlerFunc(m.api),
		moveInCurrentPlaylist:          payload_handlers.MoveInCurrentPlaylistHandlerFunc(m.api),
		batchMoveInCurrentPlaylist:     payload_handlers.BatchMoveInCurrentPlaylistHandlerFunc(m.api),
		addStoredToPos:                 payload_handlers.AddStoredPlaylistToCurrentPlaylistToPosHandlerFunc(m.api),
		seekPos:                        payload_handlers.SeekPosHandleFunc(m.api),
		deleteStoredPlaylist:           payload_handlers.DeleteStoredPlaylistHandleFunc(m.api),
		saveCurrentPlaylistAsStored:    payload_handlers.SaveCurrentPlaylistAsStoredHandleFunc(m.api),
		renameStoredPlaylist:           payload_handlers.RenameStoredPlaylistHandleFunc(m.api),
		setRandom:                      payload_handlers.SetRandomHandleFunc(m.api),
		setRepeat:                      payload_handlers.SetRepeatHandleFunc(m.api),
		setSingle:                      payload_handlers.SetSingleHandleFunc(m.api),
		setConsume:                     payload_handlers.SetConsumeHandleFunc(m.api),
		updateTree:                     payload_handlers.UpdateTreeHandleFunc(m.api),
		//"": payload_handlers.SeekPosHandleFunc(m.api),
	}
}

func (m *MpdWS) HandleFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {

	}
	cl := &client{
		conn: conn,
		send: make(chan interface{}, 100),
		ctx:  r.Context(),
	}
	m.clients[cl] = struct{}{}
	go cl.writePump()
	m.sendCurrentData(cl)
	m.startListening(cl)
}

func (m *MpdWS) sendCurrentData(cl *client) {
	// при установке соединения просто действуем как при получении события ON_CONNECT и ON_DISCONNECT
	event := mpdapi.ON_DISCONNECT
	if m.api.WithRequestContext(cl.ctx).IsConnected() {
		event = mpdapi.ON_CONNECT
	}
	for payloadType, handle := range m.eventHandleFunc[event] {
		payload, _ := handle(cl.ctx)
		cl.send <- WsResponse{
			Type:    payloadType,
			Payload: payload,
		}
	}
}

func (m *MpdWS) startListening(cl *client) {
	defer func() {
		delete(m.clients, cl)
		close(cl.send)
		cl.conn.Close()
	}()
	for {
		var msg WsRequest
		err := cl.conn.ReadJSON(&msg)
		if err != nil {
			cl.sendError(fmt.Errorf("error parsing json: %w", err), nil)
		}
		if handler, ok := m.payloadHandleFunc[msg.Type]; ok {
			if err = handler(cl.ctx, msg.Payload); err != nil {
				cl.sendError(err, msg.RequestId)
			}
		}
	}
}

func (m *MpdWS) subscribe() {
	ch := m.api.Subscribe(100 * time.Millisecond)
	go func() {
		for {
			select {
			case e := <-ch:
				handlers := m.eventHandleFunc[e]
				go func() {
					for key, handle := range handlers {
						payload, err := handle(m.ctx)
						if err == nil {
							for cl := range m.clients {
								cl.send <- WsResponse{
									Type:    key,
									Payload: payload,
								}
							}
						}
					}
				}()
			}
		}
	}()
}

func (m *MpdWS) fillEventHandlers() {

	// Суррогатный метод, возвращает nil
	resetFunc := func(ctx context.Context) (interface{}, error) {
		return nil, nil
	}

	m.eventHandleFunc = map[mpdapi.MpdEventType]map[responseType]event_handlers.EventHandle{
		mpdapi.ON_CONNECT: {
			getConnectionState:  event_handlers.GetConnectionState(m.api),
			getTree:             event_handlers.GetTreeEventHandler(m.api),
			listCurrentPlaylist: event_handlers.ListCurrentPlaylistEventHandleFunc(m.api),
			getStoredPlaylists:  event_handlers.GetStoredPlaylistsEventHandleFunc(m.api),
			listOutputs:         event_handlers.ListOutputsEventHandleFunc(m.api),
			getStatus:           event_handlers.GetStatusEventHandleFunc(m.api),
		},
		mpdapi.ON_DISCONNECT: {
			getConnectionState:  event_handlers.GetConnectionState(m.api),
			getTree:             resetFunc,
			listCurrentPlaylist: resetFunc,
			getStoredPlaylists:  resetFunc,
			listOutputs:         resetFunc,
			getStatus:           resetFunc,
		},
		mpdapi.ON_DATABASE_CHANGED: {
			getTree: event_handlers.GetTreeEventHandler(m.api),
		},
		mpdapi.ON_UPDATE_CHANGED: {
			getTree: event_handlers.GetTreeEventHandler(m.api),
		},
		mpdapi.ON_STORED_PLAYLIST_CHANGED: {
			getStoredPlaylists: event_handlers.GetStoredPlaylistsEventHandleFunc(m.api),
		},
		mpdapi.ON_PLAYLIST_CHANGED: {
			listCurrentPlaylist: event_handlers.ListCurrentPlaylistEventHandleFunc(m.api),
		},
		mpdapi.ON_PLAYER_CHANGED: {
			getStatus: event_handlers.GetStatusEventHandleFunc(m.api),
		},
		mpdapi.ON_MIXER_CHANGED: {},
		mpdapi.ON_OUTPUT_CHANGED: {
			listOutputs: event_handlers.ListOutputsEventHandleFunc(m.api)},
		mpdapi.ON_OPTIONS_CHANGED: {
			getStatus: event_handlers.GetStatusEventHandleFunc(m.api)},
		mpdapi.ON_PARTITION_CHANGED:    {},
		mpdapi.ON_STICKER_CHANGED:      {},
		mpdapi.ON_SUBSCRIPTION_CHANGED: {},
		mpdapi.ON_MESSAGE_CHANGED:      {},
	}
}

func (m *MpdWS) startUpdatingStatus() {
	playing := false
	go func() {
		ch := m.api.Subscribe(100 * time.Millisecond)
		for {
			event := <-ch
			switch event {
			case mpdapi.ON_DISCONNECT:
				playing = false
				break
			case mpdapi.ON_CONNECT:
			case mpdapi.ON_PLAYER_CHANGED:
				status, err := m.api.Status()
				if err == nil {
					playing = *(status.State) == "play"
				}
				break
			}
		}
	}()
	go func() {
		for {
			if playing && len(m.clients) > 0 {
				state, err := event_handlers.GetStatusEventHandleFunc(m.api)(m.ctx)
				if err == nil {
					for cl := range m.clients {
						b, _ := json.Marshal(state)
						log.Info("====sending state", "state", string(b))
						cl.send <- WsResponse{
							Type:    getStatus,
							Payload: state,
						}
					}
				}
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}
