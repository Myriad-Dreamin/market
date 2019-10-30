package mock

import (
	"bytes"
	"encoding/json"
	"io"
)

type JSON Stream

type J = JSON

func (s JSON) Serialize() (io.Reader, error) {
	var buf = bytes.NewBuffer(nil)
	e := json.NewEncoder(buf)
	return buf, e.Encode(s.Content)
}

func (s JSON) ContentType() string {
	return ContentTypeJSON
}
