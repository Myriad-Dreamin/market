package mock

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type ResponseI interface {
	http.ResponseWriter
	Body() *bytes.Buffer
	JSON(x interface{}) error
	XML(x interface{}) error
	Code() int
	Flushed() bool
	Flush()
	Result() *http.Response
}

type Response struct {
	*httptest.ResponseRecorder
}

func (r *Response) Body() *bytes.Buffer {
	return r.ResponseRecorder.Body
}

func (r *Response) JSON(v interface{}) error {
	return json.NewDecoder(r.ResponseRecorder.Body).Decode(v)
}

func (r *Response) XML(v interface{}) error {
	return xml.NewDecoder(r.ResponseRecorder.Body).Decode(v)
}

func (r *Response) Code() int {
	return r.ResponseRecorder.Code
}

func (r *Response) Flushed() bool {
	return r.ResponseRecorder.Flushed
}

func (r *Response) Flush() {
	r.ResponseRecorder.Flush()
	return
}

func (r *Response) Header() http.Header {
	return r.ResponseRecorder.Header()
}

func (r *Response) Result() *http.Response {
	return r.ResponseRecorder.Result()
}

func (r *Response) String() string {
	fmt.Println(r.Body())
	body := r.Body()
	var bs string
	var bb []byte
	if body != nil {
		bs, bb = body.String(), body.Bytes()
	}
	return fmt.Sprintf(
		`Code: %v, Flushed: %v,
Header: %v,
Result: %v,
Bodybytes: %v
BodyString: %v
`, r.Code(), r.Flushed(), r.Header(), r.Result(), bb, bs)
}

func NewResponse() *Response {
	resp := &Response{ResponseRecorder: new(httptest.ResponseRecorder)}
	resp.ResponseRecorder.Body = bytes.NewBuffer(nil)
	return resp
}

