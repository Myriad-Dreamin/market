package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Myriad-Dreamin/market/mock"
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

type Mocker struct {
	*Server
	cancel        func()
	contextHelper *testing.T
}

type MockerContext struct {
	*Mocker
	*assert.Assertions
}

func Mock() (srv *Mocker) {
	srv = new(Mocker)
	srv.Server = newServer()
	if !(srv.InstantiateLogger() &&
		srv.MockDatabase()) {
		srv = nil
		return
	}
	defer func() {
		if err := recover(); err != nil {
			printStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		} else if srv == nil {
			srv.Terminate()
		}
	}()

	if !(srv.PrepareMiddleware() &&
		srv.PrepareService() &&
		srv.BuildRouter()) {
		srv = nil
		return
	}

	if err := srv.Module.Install(srv.RouterProvider); err != nil {
		fmt.Println("install router provider error", err)
	}
	if err := srv.Module.Install(srv.DatabaseProvider); err != nil {
		fmt.Println("install database provider error", err)
	}

	defer func() {
		if err := recover(); err != nil {
			printStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	srv.Router.Root.Build(srv.RouterEngine)
	srv.Module.Debug(srv.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if srv == nil {
			cancel()
		}
	}()

	for _, plg := range srv.plugins {
		go plg.Work(ctx)
	}

	if err := dblayer.GetRawInstance().Ping(); err != nil {
		srv.Logger.Debug("database died", "error", err)
		srv = nil
		return
	}

	srv.cancel = cancel
	return
}

func (mocker *Mocker) Context(t *testing.T) *MockerContext {
	return &MockerContext{
		Mocker:        &Mocker{
			Server:        mocker.Server,
			cancel:        nil,
			contextHelper: t,
		},
		Assertions: assert.New(t),
	}
}

func (mocker *Mocker) ReleaseMock() {
	if mocker.cancel != nil {
		mocker.cancel()
		mocker.cancel = nil
	}
	if err := os.Remove("./test.db"); err != nil {
		mocker.Logger.Error("clear db error", "error", err)
	}
}

type Request = http.Request

type ResponseI interface {
	Body() *bytes.Buffer
	Code() int
	Flushed() bool
	Flush()
	Header() http.Header
	Result() *http.Response
}

type Response struct {
	r *httptest.ResponseRecorder
}

func (r *Response) Body() *bytes.Buffer {
	return r.r.Body
}

func (r *Response) Code() int {
	return r.r.Code
}

func (r *Response) Flushed() bool {
	return r.r.Flushed
}

func (r *Response) Flush() {
	r.r.Flush()
	return
}

func (r *Response) Header() http.Header {
	return r.r.Header()
}

func (r *Response) Result() *http.Response {
	return r.r.Result()
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


func newResponse() *Response {
	return &Response{r: new(httptest.ResponseRecorder)}
}

func (srv *Server) mockServe(r *Request) (w *Response) {
	w = newResponse()
	w.r.Body = bytes.NewBuffer(nil)
	srv.RouterEngine.ServeHTTP(w.r, r)

	return
}

func (mocker *Mocker) report(err error) {
	if mocker.contextHelper != nil {
		mocker.contextHelper.Helper()
		mocker.contextHelper.Error(err)
	} else {
		mocker.Logger.Error("error occurs", "error", err)
	}
}

func (mocker *Mocker) Method(method, path string, params ...interface{}) ResponseI {
	var body io.Reader
	var contentType string
	if len(params) > 0 {
		switch p := params[0].(type) {
		case string, []byte:
			body = mock.NotStruct(p)
		case mock.Serializable:
			var err error
			body, err = p.Serialize()
			if err != nil {
				mocker.report(err)
				return nil
			}
			contentType = p.ContentType()
		case *url.Values:
			body = strings.NewReader(p.Encode())
		case io.Reader:
			body = p
		case *http.Request:
			return mocker.mockServe(p)
		case http.Request:
			return mocker.mockServe(&p)
		default:
			buf := bytes.NewBuffer(nil)
			body = buf
			if err := json.NewEncoder(buf).Encode(p); err != nil {
				mocker.Logger.Error("encode request to json error", "error", err)
			}
			contentType = "application/json"
		}
	}
	r, err := http.NewRequest(method, path, body)
	if err != nil {
		mocker.report(err)
		return nil
	}
	r.Header.Set("Content-Type", contentType)
	return mocker.mockServe(r)
}

func (mocker *Mocker) Get(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodGet, path, params...)
}

func (mocker *Mocker) Connect(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodConnect, path, params...)
}

func (mocker *Mocker) Delete(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodDelete, path, params...)
}

func (mocker *Mocker) Head(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodHead, path, params...)
}

func (mocker *Mocker) Options(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodOptions, path, params...)
}

func (mocker *Mocker) Patch(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodPatch, path, params...)
}

func (mocker *Mocker) Post(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodPost, path, params...)
}

func (mocker *Mocker) Put(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodPut, path, params...)
}

func (mocker *Mocker) Trace(path string, params ...interface{}) ResponseI {
	return mocker.Method(http.MethodTrace, path, params...)
}

