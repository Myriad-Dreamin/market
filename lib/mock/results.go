package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type ResultsI interface {
	GetMethod() string
	GetPath() string
	GetRecords() []RecordsI
}

type RecordsI interface {
	GetRequestHeader() http.Header
	GetRequestBody() []byte
	GetResponseHeader() http.Header
	GetResponseBody() []byte
}

type Records struct {
	RequestHeader  http.Header
	RequestBody    []byte
	ResponseHeader http.Header
	ResponseBody   []byte
}

func (r Records) GetRequestHeader() http.Header {
	return r.RequestHeader
}

func (r Records) GetRequestBody() []byte {
	return r.RequestBody
}

func (r Records) GetResponseHeader() http.Header {
	return r.ResponseHeader
}

func (r Records) GetResponseBody() []byte {
	return r.ResponseBody
}

type Results struct {
	gin.RouteInfo
	Recs []RecordsI
}

func (r Results) GetMethod() string {
	return r.Method
}

func (r Results) GetPath() string {
	return r.Path
}

func (r Results) GetRecords() []RecordsI {
	return r.Recs
}
