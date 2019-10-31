package mock

import (
	"bytes"
	"errors"
	"github.com/gogo/protobuf/proto"
	"io"
)

type ProtoMessage Stream

var (
	ErrNotProtoMessage = errors.New("bad message")
)

func (s ProtoMessage) Serialize() (io.Reader, error) {
	if msg, ok := s.Content.(proto.Message); ok {
		var buf = bytes.NewBuffer(nil)
		return buf, proto.MarshalText(buf, msg)
	} else {
		return nil, ErrNotProtoMessage
	}
}


func (s ProtoMessage) ContentType() string {
	return ContentTypeProtocolBufBinary
}
