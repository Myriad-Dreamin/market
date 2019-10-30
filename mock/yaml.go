package mock

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io"
)

type YAML struct {
	Stream
}

type Y = YAML

func (s YAML) Serialize() (io.Reader, error) {
	var buf = bytes.NewBuffer(nil)
	e := yaml.NewEncoder(buf)
	return buf, e.Encode(s.Content)
}

func (s YAML) ContentType() string {
	return ContentTypeYAML
}

