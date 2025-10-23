package ws

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
type WsRequest struct {
	Type      requestType     `json:"@type"`
	RequestId *uuid.UUID      `json:"requestId,omitempty"`
	Payload   json.RawMessage `json:"payload,omitempty"`
}

type WsResponse struct {
	Type      responseType `json:"@type,omitempty"`
	RequestId *uuid.UUID   `json:"requestId,omitempty"`
	Error     *string      `json:"error,omitempty"`
	Payload   interface{}  `json:"payload,omitempty"`
}
