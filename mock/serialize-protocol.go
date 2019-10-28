package mock

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
	"io"
	"reflect"
)

type Stream struct {
	Content interface{}
}

func NotStruct(content interface{}) io.Reader {
	var buf = new(bytes.Buffer)
	switch ss := content.(type) {
	case string:
		_, err := buf.WriteString(ss)
		if err != nil {
			panic(err)
		}
	case []byte:
		_, err := buf.Write(ss)
		if err != nil {
			panic(err)
		}
	case reflect.Value:
		switch ss.Kind() {
		case reflect.Ptr, reflect.Interface:
			return NotStruct(ss.Elem())
		case reflect.Struct, reflect.Slice, reflect.Map:
			return nil
		}
	default:
		return NotStruct(reflect.ValueOf(content))
	}
	return buf
}

func (s Stream) notStruct() io.Reader {
	return NotStruct(s.Content)
}

type Serializable interface {
	Serialize() (io.Reader, error)
	ContentType() string
}

type MJ map[string]interface{}

func (s MJ) Serialize() (io.Reader, error) {
	return JSON{s}.Serialize()
}

func (s MJ) ContentType() string {
	return "application/json"
}

type MY map[string]interface{}

func (s MY) Serialize() (io.Reader, error) {
	return YAML{Stream: Stream{s}}.Serialize()
}

func (s MY) ContentType() string {
	return "application/yaml"
}

type MX map[string]interface{}

func (s MX) Serialize() (io.Reader, error) {
	return XML{s}.Serialize()
}

func (s MX) ContentType() string {
	return "application/xml"
}

type JSON Stream

type J = JSON

func (s JSON) Serialize() (io.Reader, error) {
	if r := (Stream)(s).notStruct(); r!= nil {
		return r, nil
	}
	var buf = bytes.NewBuffer(nil)
	e := json.NewEncoder(buf)
	return buf, e.Encode(s.Content)
}

func (s J) ContentType() string {
	return "application/json"
}

type YAML struct {
	Stream
}

type Y = YAML

func (s YAML) Serialize() (io.Reader, error) {
	if r := s.notStruct(); r!= nil {
		return r, nil
	}
	var buf = bytes.NewBuffer(nil)
	e := yaml.NewEncoder(buf)
	return buf, e.Encode(s.Content)
}

func (s Y) ContentType() string {
	return "application/yaml"
}

type XML Stream

type X = XML

func (s XML) Serialize() (io.Reader, error) {
	if r := (Stream)(s).notStruct(); r!= nil {
		return r, nil
	}
	var buf = bytes.NewBuffer(nil)
	e := xml.NewEncoder(buf)
	return buf, e.Encode(s.Content)
}

func (s X) ContentType() string {
	return "application/xml"
}

type ProtoMessage struct {
	Stream
}

func (s ProtoMessage) Serialize() (io.Reader, error) {
	if r := s.notStruct(); r!= nil {
		return r, nil
	}
	if msg, ok := s.Content.(proto.Message); ok {
		var buf = bytes.NewBuffer(nil)
		return buf, proto.MarshalText(buf, msg)
	} else {
		return nil, errors.New("bad message")
	}
}


