package mock

import "io"

const (
	ContentTypeJSON              = "application/json"
	ContentTypeYAML              = "application/x-yaml"
	ContentTypeXML               = "application/xml"
	ContentTypeProtocolBufBinary = "application/grpc-web+proto"
)

type MJ map[string]interface{}

func (s MJ) Serialize() (io.Reader, error) {
	return JSON{s}.Serialize()
}

func (s MJ) ContentType() string {
	return ContentTypeJSON
}

type MY map[string]interface{}

func (s MY) Serialize() (io.Reader, error) {
	return YAML{Stream: Stream{s}}.Serialize()
}

func (s MY) ContentType() string {
	return ContentTypeYAML
}

type MX map[string]interface{}

func (s MX) Serialize() (io.Reader, error) {
	return XML{s}.Serialize()
}

func (s MX) ContentType() string {
	return ContentTypeXML
}
