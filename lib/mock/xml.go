package mock

import (
	"bytes"
	"encoding/xml"
	"io"
)

type XML Stream

type X = XML

func (s XML) Serialize() (io.Reader, error) {
	var buf = bytes.NewBuffer(nil)
	e := xml.NewEncoder(buf)
	return buf, e.Encode(s.Content)
}

func (s XML) ContentType() string {
	return ContentTypeXML
}
