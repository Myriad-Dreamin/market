package doc_gen

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/lib/mock"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_prettifyHeader(t *testing.T) {
	type args struct {
		h      http.Header
		indent int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test_base01", args{
			h: http.Header{
				"Content-Type": []string{"application/json"},
			},
			indent: 2,
		}, "        Content-Type: application/json"},
		{"test_base02", args{
			h: http.Header{
				"Content": []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
			indent: 2,
		}, "        Content: application/json\n        Content-Type: application/json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettifyHeader(tt.args.h, tt.args.indent); got != tt.want {
				t.Errorf("prettifyHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prettifyBody(t *testing.T) {
	type args struct {
		content     []byte
		contentType string
		indent      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test_base01", args{
			content: []byte(`{"a":1}`),
			contentType: mock.ContentTypeJSON,
			indent: 2,
		}, `        {
            "a": 1
        }`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prettifyBody(tt.args.content, tt.args.contentType, tt.args.indent); got != tt.want {
				t.Errorf("prettifyBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

// @category Accounts
// @titLE Get Accounts
//          @Description get accounts
//   @Success 200 {array} model.Account
func testHandler(ctx *gin.Context) {

}

// @category Accounts
// @titLE Put Accounts
//          @Description put accounts
//   @Success 200 {array} model.Account
func testHandler2(ctx *gin.Context) {

}


type controllerS struct {

}

// @titLE Get AccountsA
//          @Description get accounts
//   @Success 200 {array} model.Account
func (controllerS) TestHandler(ctx *gin.Context) {

}

// @titLE Get AccountsA
//          @Description get accounts
//   @Success 200 {array} model.Account
func (controllerS) testHandler2(ctx *gin.Context) {

}

// @category AccountsController
type controller interface {
	// @category Accounts
	// @titLE Get Accounts
	//          @Description get accounts
	//   @Success 200 {array} model.Account
	TestHandler(ctx *gin.Context)

	// @category Accounts
	// @titLE Put Accounts
	//          @Description put accounts
	//   @Success 200 {array} model.Account
	testHandler2(ctx *gin.Context)
}

func TestFromGinResults(t *testing.T) {
	type args struct {
		res     *GinInfo
		options []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test_base01", args{
			res: &GinInfo{
				Result:[]GinResult{mock.Results{
					RouteInfo: gin.RouteInfo{
						Method:      "PUT",
						Path:        "/v1/user",
						Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
						HandlerFunc: testHandler,
					},
					Recs:      []Records{mock.Records{
						RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
						RequestBody:    []byte(`{"1":1}`),
						ResponseCode:   200,
						ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
						ResponseBody:   []byte(`{"2":2}`),
					}},
				}},
				Host: "127.0.0.1",
				ApiDoc: "Minimum Document",

			},
		}, false},
		{"test_base02", args{
			res: &GinInfo{
				Result:[]GinResult{mock.Results{
					RouteInfo: gin.RouteInfo{
						Method:      "PUT",
						Path:        "/v1/user",
						Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
						HandlerFunc: testHandler2,
					},
					Recs:      []Records{mock.Records{
						RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
						RequestBody:    []byte(`{"1":1}`),
						ResponseCode:   200,
						ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
						ResponseBody:   []byte(`{"2":2}`),
					}},
				},
					mock.Results{
						RouteInfo: gin.RouteInfo{
							Method:      "GET",
							Path:        "/v1/user",
							Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
							HandlerFunc: testHandler,
						},
						Recs:      []Records{mock.Records{
							RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
							RequestBody:    []byte(`{"1":1}`),
							ResponseCode:   200,
							ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
							ResponseBody:   []byte(`{"2":2}`),
						}},
					},
				},
				Host: "127.0.0.1",
				ApiDoc: "Minimum Document",

			},
		}, false},
		{"test_base03", args{
			res: &GinInfo{
				Result:[]GinResult{mock.Results{
					RouteInfo: gin.RouteInfo{
						Method:      "PUT",
						Path:        "/v1/user",
						Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
						HandlerFunc: controllerS{}.testHandler2,
					},
					Recs:      []Records{mock.Records{
						RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
						RequestBody:    []byte(`{"1":1}`),
						ResponseCode:   200,
						ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
						ResponseBody:   []byte(`{"2":2}`),
					}},
				},
					mock.Results{
						RouteInfo: gin.RouteInfo{
							Method:      "GET",
							Path:        "/v1/user",
							Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
							HandlerFunc: controllerS{}.TestHandler,
						},
						Recs:      []Records{mock.Records{
							RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
							RequestBody:    []byte(`{"1":1}`),
							ResponseCode:   200,
							ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
							ResponseBody:   []byte(`{"2":2}`),
						}},
					},
				},
				Host: "127.0.0.1",
				ApiDoc: "Minimum Document",

			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromGinResults(tt.args.res, tt.args.options...); (err != nil) != tt.wantErr {
				t.Errorf("FromGinResults() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_processResultResult(t *testing.T) {
	type args struct {
		resI interface{}
	}
	var ri RawServiceI = RawService{}
	tests := []struct {
		name string
		args args
		want processedResultInterface
	}{
		{"test_base_just_func", args{
			resI:mock.Results{
				RouteInfo: gin.RouteInfo{
					Method:      "PUT",
					Path:        "/v1/user",
					Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
					HandlerFunc: testHandler,
				},
				Recs:      []Records{mock.Records{
					RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
					RequestBody:    []byte(`{"1":1}`),
					ResponseCode:   200,
					ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
					ResponseBody:   []byte(`{"2":2}`),
				}},
			},
		}, nil},
		{"test_base_func_with_receiver", args{
			resI:mock.Results{
				RouteInfo: gin.RouteInfo{
					Method:      "PUT",
					Path:        "/v1/user",
					Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
					HandlerFunc: RawService{}.TestHandler,
				},
				Recs:      []Records{mock.Records{
					RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
					RequestBody:    []byte(`{"1":1}`),
					ResponseCode:   200,
					ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
					ResponseBody:   []byte(`{"2":2}`),
				}},
			},
		}, nil},
		{"test_base_interface_func_with_receiver", args{
			resI:mock.Results{
				RouteInfo: gin.RouteInfo{
					Method:      "PUT",
					Path:        "/v1/user",
					Handler:     "ginhub.com/Myriad-Dreamin/x.y.z",
					HandlerFunc: ri.TestHandler,
				},
				Recs:      []Records{mock.Records{
					RequestHeader:  http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
					RequestBody:    []byte(`{"1":1}`),
					ResponseCode:   200,
					ResponseHeader: http.Header{"Content-Type": []string{mock.ContentTypeJSON}},
					ResponseBody:   []byte(`{"2":2}`),
				}},
			},
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := processResultResult(tt.args.resI); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("processResultResult() = %v, want %v", got, tt.want)
			//}
			fmt.Println(processResultResult(tt.args.resI))
			fmt.Println(processResultResult(tt.args.resI).GetDescription())
			fmt.Println(processResultResult(tt.args.resI).GetTitle())
		})
	}
}