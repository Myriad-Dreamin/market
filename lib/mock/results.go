package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type ResultsI interface {
	GetMethod() string
	GetPath() string
	GetHandlerFunc() interface{}
	GetRecords() []RecordsI
}

type GinResultI interface {
	ResultsI
	GetHandler() string
}

type RecordsI interface {
	GetRequestHeader() http.Header
	GetRequestBody() []byte
	GetResponseCode() int
	GetResponseHeader() http.Header
	GetResponseBody() []byte
}

type Records struct {
	RequestHeader  http.Header
	RequestBody    []byte
	ResponseCode int
	ResponseHeader http.Header
	ResponseBody   []byte
}

func (r Records) GetRequestHeader() http.Header {
	return r.RequestHeader
}

func (r Records) GetRequestBody() []byte {
	return r.RequestBody
}

func (r Records) GetResponseCode() int {
	return r.ResponseCode
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

func (r Results) GetHandlerFunc() interface{} {
	return r.HandlerFunc
}

func (r Results) GetHandler() string {
	return r.Handler
}

func (r Results) GetRecords() []RecordsI {
	return r.Recs
}
