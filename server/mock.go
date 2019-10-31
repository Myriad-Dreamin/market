package server

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	abstract_test "github.com/Myriad-Dreamin/market/lib/abstract-test"
	"github.com/Myriad-Dreamin/market/lib/mock"
	"github.com/Myriad-Dreamin/market/lib/tracer"
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	ginhelper "github.com/Myriad-Dreamin/market/service/gin-helper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

type ContextHelperInterface interface {
	Helper()
	Log(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Logf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Records struct {
	RequestHeader http.Header
	RequestBody []byte
	ResponseHeader http.Header
	ResponseBody []byte
}

type Results struct {
	gin.RouteInfo
	Recs []Records
}

type Mocker struct {
	*Server
	cancel             func()
	header             map[string]string
	routes             map[string]*Results
	contextHelper      ContextHelperInterface
	shouldPrintRequest bool
	assertNoError      bool
	collectResults bool
}

type MockerContext struct {
	*Mocker
	*assert.Assertions
}

func Mock() (srv *Mocker) {
	srv = new(Mocker)
	srv.Server = newServer()
	srv.header = make(map[string]string)
	if !(srv.InstantiateLogger() &&
		srv.MockDatabase()) {
		srv = nil
		return
	}
	defer func() {
		if err := recover(); err != nil {
			tracer.PrintStack()
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
			tracer.PrintStack()
			srv.Logger.Error("panic error", "error", err)
			srv.Terminate()
		}
	}()

	srv.RouterEngine.Use(mock.ContextRecorder())
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
	srv.contextHelper = &abstract_test.ContextHelper{log.New(os.Stdout, "mocker", log.Ldate|log.Ltime|log.Llongfile|log.LstdFlags)}


	routes := srv.RouterEngine.Routes()
	srv.routes = make(map[string]*Results)
	for _, route := range routes {
		srv.routes[route.Path + "@" + route.Method] = &Results{
			RouteInfo: route,
			Recs:      nil,
		}
	}

	return
}

func (mocker *Mocker) PrintRequest(p bool) {
	mocker.shouldPrintRequest = p
}

func (mocker *Mocker) Context(t *testing.T) *MockerContext {
	m := new(Mocker)
	*m = *mocker
	m.contextHelper = t
	return &MockerContext{
		Mocker:        m,
		Assertions: assert.New(t),
	}
}

func (mocker *Mocker) ReleaseMock() {
	if mocker.cancel != nil {
		mocker.cancel()
		mocker.cancel = nil
	}
}

type Request = http.Request

type ResponseI interface {
	Body() *bytes.Buffer
	JSON(x interface{}) error
	XML(x interface{}) error
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

func (r *Response) JSON(v interface{}) error {
	return json.NewDecoder(r.r.Body).Decode(v)
}

func (r *Response) XML(v interface{}) error {
	return xml.NewDecoder(r.r.Body).Decode(v)
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

type rc struct {
	*bytes.Buffer
}

func (rc) Close() error {
	return nil
}

func (mocker *Mocker) mockServe(r *Request) (w *Response) {

	w = newResponse()
	w.r.Body = bytes.NewBuffer(nil)
	var b []byte
	var err error
	if mocker.collectResults {
		if r.Body != nil {
			b, err = ioutil.ReadAll(r.Body)
			_ = r.Body.Close()
			if err != nil {
				mocker.contextHelper.Fatal("read failed", "error", err)
			}
			r.Body = rc{bytes.NewBuffer(b)}
		}
	}

	mocker.RouterEngine.ServeHTTP(w.r, r)

	if mocker.contextHelper != nil && mocker.assertNoError {
		if !mocker.NoErr(w) {
			mocker.contextHelper.Fatal("stopped by assertion")
		}
	}

	if mocker.shouldPrintRequest {
		fmt.Println("Method:", r.Method, "url:", r.URL, "http:",  r.Proto)
		fmt.Println("Request Header:", r.Header)
		fmt.Println("Response Header:", w.r.Header())
	}

	if mocker.collectResults {
		c := make([]byte, w.r.Body.Len())
		copy(c, w.r.Body.Bytes())
		pattern := w.r.Header().Get("Gin-Context-Matched-Path-Method")
		w.r.Header().Del("Gin-Context-Matched-Path-Method")
		if results, ok := mocker.routes[pattern]; ok {
			rec := Records{
				RequestBody:    b,
				ResponseBody:   c,
			}
			rec.RequestHeader = make(http.Header)
			for k,v := range r.Header {
				rec.RequestHeader[k] = v
			}
			rec.ResponseHeader = make(http.Header)
			for k,v := range w.r.Header() {
				rec.ResponseHeader[k] = v
			}
			results.Recs = append(results.Recs, rec)
		} else {
			mocker.contextHelper.Fatal("matched bad route", pattern)
		}
	}

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
	for k, v := range mocker.header {
		r.Header.Set(k, v)
	}
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

func (mocker *Mocker) SetHeader(k, v string) {
	mocker.header[k] = v
}

func (mocker *Mocker) UseToken(token string) {
	mocker.header[mocker.jwtMW.JWTHeaderKey] =
		mocker.jwtMW.JWTHeaderPrefixWithSplitChar + token
}

func (mocker *MockerContext) AssertNoError(noErr bool) *MockerContext {
	mocker.assertNoError = noErr
	return mocker
}

func (mocker *Mocker) CollectResults(collectResults bool) *Mocker {
	mocker.collectResults = collectResults
	return mocker
}


func (mocker *Mocker) NoErr(resp ResponseI) bool {
	if mocker.contextHelper == nil {
		panic("only used in test")
	}
	mocker.contextHelper.Helper()
	if resp.Code() != 200 {
		mocker.contextHelper.Error("resp has bad code ", resp.Code())
		return false
	}
	body := resp.Body()
	var obj ginhelper.ErrorSerializer
	if err := json.Unmarshal(body.Bytes(), &obj); err != nil {
		mocker.contextHelper.Error(err)
		return false
	}
	if len(obj.Error) != 0 || obj.Code != 0 {
		mocker.contextHelper.Errorf("Code, Error (%v, %v)", obj.Code, obj.Error)
	}
	return true
	//if gjson
}