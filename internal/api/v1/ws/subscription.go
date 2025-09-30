package ws

import (
	"context"
	"fmt"
	"time"

	log "github.com/anpotashev/mpd-ws-api/internal/logger"
	"github.com/anpotashev/mpdgo/pkg/mpdapi"
)

type SubscribeRsPayload struct{}

// мапа, где ключ - это idle событие, а значение - имплементация processable
// при получении idle события, значение используется чтобы получить payload и payloadType
// для отправки broadcast уведомления  всем подписавшимся
var eventTypeToMpdProcessableMap = map[mpdapi.MpdEventType][]processable{
	mpdapi.ON_CONNECT:                 {&GetConnectionStateRequest{}, &GetTreeRequest{}, &ListPlaylistRequest{}, &ListOutputRequest{}, &GetStatusRequest{}, &GetStoredPlaylistsRequest{}},
	mpdapi.ON_DISCONNECT:              {&GetConnectionStateRequest{}, &ResetTree{}, &ResetPlaylist{}, &ResetOutput{}, &ResetStatus{}, &ResetStoredPlaylists{}},
	mpdapi.ON_DATABASE_CHANGED:        {&GetTreeRequest{}},
	mpdapi.ON_PLAYLIST_CHANGED:        {&ListPlaylistRequest{}, &GetStatusRequest{}},
	mpdapi.ON_OUTPUT_CHANGED:          {&ListOutputRequest{}},
	mpdapi.ON_OPTIONS_CHANGED:         {&GetStatusRequest{}},
	mpdapi.ON_STORED_PLAYLIST_CHANGED: {&GetStoredPlaylistsRequest{}},
	mpdapi.ON_PLAYER_CHANGED:          {&GetStatusRequest{}},
}

//var eventTypeToTopicNameMap = map[mpdapi.MpdEventType]string{
//	mpdapi.ON_CONNECT:                 "on_connect",
//	mpdapi.ON_DISCONNECT:              "on_disconnect",
//	mpdapi.ON_DATABASE_CHANGED:        "database",
//	//mpdapi.ON_UPDATE_CHANGED:          "update",
//	mpdapi.ON_STORED_PLAYLIST_CHANGED: "stored_playlist",
//	mpdapi.ON_PLAYLIST_CHANGED:        "playlist",
//	mpdapi.ON_PLAYER_CHANGED:          "player",
//	//mpdapi.ON_MIXER_CHANGED:           "mixer",
//	mpdapi.ON_OUTPUT_CHANGED:          "output",
//	mpdapi.ON_OPTIONS_CHANGED:         "options",
//	//mpdapi.ON_PARTITION_CHANGED:       "partition",
//	//mpdapi.ON_STICKER_CHANGED:         "sticker",
//	//mpdapi.ON_SUBSCRIPTION_CHANGED:    "subscription",
//	//mpdapi.ON_MESSAGE_CHANGED:         "message",
//}

// мапа исползуемая для получения значения mpdapi.MpdEventType по занчению topic
// из запросов пользователей на подписку на события
var topicNameToEventTypeMap = map[string]mpdapi.MpdEventType{
	"on_connect":    mpdapi.ON_CONNECT,
	"on_disconnect": mpdapi.ON_DISCONNECT,
	"database":      mpdapi.ON_DATABASE_CHANGED,
	//"update":          mpdapi.ON_UPDATE_CHANGED,
	"stored_playlist": mpdapi.ON_STORED_PLAYLIST_CHANGED,
	"playlist":        mpdapi.ON_PLAYLIST_CHANGED,
	"player":          mpdapi.ON_PLAYER_CHANGED,
	//"mixer":           mpdapi.ON_MIXER_CHANGED,
	"output":  mpdapi.ON_OUTPUT_CHANGED,
	"options": mpdapi.ON_OPTIONS_CHANGED,
	//"partition":       mpdapi.ON_PARTITION_CHANGED,
	//"sticker":         mpdapi.ON_STICKER_CHANGED,
	//"subscription":    mpdapi.ON_SUBSCRIPTION_CHANGED,
	//"message":         mpdapi.ON_MESSAGE_CHANGED,
}

type SubscribeRqPayload struct {
	Topics []string `json:"topics"`
}

func (req *SubscribeRqPayload) getPayloadType() payloadType {
	return subscribe
}

func (req *SubscribeRqPayload) process(ctx context.Context) (interface{}, error) {
	return SubscribeRsPayload{}, nil
}

func (cl *client) sendCurrentMpdData() {
	var processables []processable
	if mpdApi.IsConnected() {
		processables = eventTypeToMpdProcessableMap[mpdapi.ON_CONNECT]
	} else {
		processables = eventTypeToMpdProcessableMap[mpdapi.ON_CONNECT]
	}
	for _, p := range processables {
		fmt.Println("processing", p)
		payloadType := p.getPayloadType()
		payload, err := p.process(cl.ctx)
		if err != nil {
			errMsg := err.Error()
			cl.send <- WsResponse{PayloadType: payloadType, Error: &errMsg}
		}
		cl.send <- WsResponse{PayloadType: payloadType, Payload: payload}
	}
}

func (cl *client) subscribe(topics []string) error {
	var subscriptionMap = map[mpdapi.MpdEventType]bool{}
	for _, topic := range topics {
		eventType, ok := topicNameToEventTypeMap[topic]
		if !ok {
			return fmt.Errorf("unknown topic: %s", topic)
		}
		subscriptionMap[eventType] = true
	}
	cl.topics = subscriptionMap
	return nil
}

func listenMpdEvent() {
	eventChannel := mpdApi.Subscribe(100 * time.Millisecond)
	playing := false
	checkStatus := func() {
		status, err := mpdApi.Status()
		if err != nil {
			log.Error("error fetching new status", err)
		}
		playing = status.State != nil && *(status.State) == "play"
	}
	if mpdApi.IsConnected() {
		checkStatus()
	}
	go func() {
		for range time.NewTicker(time.Second).C {
			if playing {
				eventChannel <- mpdapi.ON_OPTIONS_CHANGED
			}
		}
	}()
	for {
		event := <-eventChannel
		go func() {
			if event == mpdapi.ON_PLAYER_CHANGED || event == mpdapi.ON_CONNECT {
				checkStatus()
			}
			if event == mpdapi.ON_DISCONNECT {
				playing = false
			}
		}()
		go func() {
			processables, ok := eventTypeToMpdProcessableMap[event]
			if !ok {
				return
			}
			for _, mpdProcessable := range processables {
				payload, err := mpdProcessable.process(context.Background())
				if err != nil {
					return
				}
				payloadType := mpdProcessable.getPayloadType()
				wsEvent := WsResponse{
					PayloadType: payloadType,
					Payload:     payload,
				}
				for c := range hub.clients {
					//if _, ok = c.topics[event]; ok {
					c.send <- wsEvent
					//}
				}
			}
		}()
	}
}
