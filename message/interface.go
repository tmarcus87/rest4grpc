package message

import "github.com/gogo/protobuf/proto"

type Message interface {
	Apply(msg proto.Message) error
}
