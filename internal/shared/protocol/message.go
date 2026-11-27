package protocol

import (
	"encoding/json"
	"io"
)

type MessageType string

const (
	MessageTypeWelcome MessageType = "welcome"
	MessageTypeMove    MessageType = "move"
	MessageTypeState   MessageType = "state"
	MessageTypeError   MessageType = "error"
)

type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type MessagePayloadWelcome struct {
	PlayerID int `json:"player_id"`
}

type MessagePayloadMove struct{}

type ErrorCode string

type ErrorPayload struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

const (
	ErrorCodeServerFull   ErrorCode = "server_full"
	ErrorCodeInvalidMove  ErrorCode = "invalid_move"
	ErrorCodeBadMessage   ErrorCode = "bad_message"
	ErrorCodeUnknownError ErrorCode = "unknown_error"
)

func Write(w io.Writer, msgType MessageType, payload any) error {
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(Message{Type: msgType, Payload: raw})
}
