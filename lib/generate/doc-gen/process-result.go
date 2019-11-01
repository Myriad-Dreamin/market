package doc_gen

import (
	"fmt"
	"strings"
)

type processedGinResult struct {
	GinResult
	*parsedResult
}

func (p processedGinResult) GetTitle() string {
	titleV := p.parsedResult.GetTitle()
	if len(titleV) == 0 {
		return title(p.GinResult)
	}
	return titleV
}

type processedResultInterface interface {
	GinResult
	GetDescription() string
	GetTitle() string
}

type GinInfo struct{
	Result []GinResult
	Host string
	ApiDoc string
}

type ginProcessedInfo struct{
	Result []processedResultInterface
	Host string
	ApiDoc string
}

type parsedResult struct {
	Description string
	Title string
	Success string
}

func (p parsedResult) GetDescription() string {
	return p.Description
}

func (p parsedResult) GetTitle() string {
	return p.Title
}


func parseDoc(doc string) *parsedResult {
	segs := strings.Split(doc, "@")
	var res = new(parsedResult)
	for i := range segs {
		kv := strings.SplitN(strings.TrimSpace(segs[i]), " ", 2)
		if len(kv) == 1 {
			continue
		}
		k, v := strings.ToLower(strings.TrimSpace(kv[0])), strings.TrimSpace(kv[1])
		switch k {
		case "description":
			res.Description = v
		case "title":
			res.Title = v
		case "success":
			res.Success = v
		}
	}
	return res

	//var buf = bytes.NewBufferString(doc)
	//var res = new(parsedResult)
	//for b, err := buffer.ReadByte(); ; b, err = buffer.ReadByte() {
	//switch err {
	//case nil:
	//	switch b {
	//	case '@':
	//
	//	}
	//case io.EOF:
	//	return res
	//}
	//}
}

func processResultResult(resI interface{}) processedResultInterface {
	switch res := resI.(type) {
	case GinResult:
		y := new(processedGinResult)
		y.GinResult = res
		y.parsedResult = parseDoc(FuncDescription(y.GetHandlerFunc()))
		return y
	default:
		return nil
	}
}

func processResultResults(resI interface{}) []processedResultInterface {
	switch res := resI.(type) {
	case []GinResult:
		var x = make([]processedResultInterface, len(res))
		for r := range res {
			x[r] = processResultResult(res[r])
		}
		return x
	default:
		return nil
	}
}

func processResult(resI interface{}) *ginProcessedInfo {
	switch res := resI.(type) {
	case *GinInfo:
		var x = new(ginProcessedInfo)
		x.Host = res.Host
		x.ApiDoc = res.ApiDoc
		x.Result = processResultResults(res.Result)
		fmt.Println(x)
		return x
	default:
		return &ginProcessedInfo{}
	}
}
