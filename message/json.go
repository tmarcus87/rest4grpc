package message

import (
	"bytes"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

type JsonMessage struct {
	json []byte
}

func NewJsonMessage(json []byte) *JsonMessage {
	return &JsonMessage{json: json}
}

func (m *JsonMessage) Apply(msg proto.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(m.json), msg)
}
