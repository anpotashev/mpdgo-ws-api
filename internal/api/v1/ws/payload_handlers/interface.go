package payload_handlers

import (
	"context"
	"encoding/json"
)

// WsPayloadHandleFunc is a websocket handler
type WsPayloadHandleFunc func(ctx context.Context, payload json.RawMessage) error
