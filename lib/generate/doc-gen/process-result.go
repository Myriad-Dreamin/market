package doc_gen

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/lib/parser"
	"sort"
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
	GetCategory() string
}

type ControllerProvider interface {
	GetControllers() []interface{}
	GetProvider() interface{}
}

type GinInfo struct{
	ControllerProvider ControllerProvider
	Result []GinResult
	Host string
	DocName string
	Description string
}

type mergedResultInterface interface {
	GetPath() string
	GetCategoryName() string
	GetTitle() string
	GetDescription() string
	GetResults() []processedResultInterface
}

type mergedResult struct {
	Path string
	Cate string
	Description string
	Results []processedResultInterface
}

func (m *mergedResult) GetPath() string {
	return m.Path
}

func (m *mergedResult) GetCategoryName() string {
	return m.Cate
}

func (m *mergedResult) GetDescription() string {
	return m.Description
}

func (m *mergedResult) GetTitle() string {
	return m.Description
}

func (m *mergedResult) GetResults() []processedResultInterface {
	return m.Results
}

type ginProcessedInfo struct{
	ProviderInfoInterface
	Categories []mergedResultInterface
	Host string
}

type parsedResult struct {
	Description string
	Title string
	Category string
	Success string
}

func (p parsedResult) GetDescription() string {
	return p.Description
}

func (p parsedResult) GetCategory() string {
	return p.Category
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
		case "category":
			res.Category = v
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
		y.parsedResult = parseDoc(parser.FuncDescription(y.GetHandlerFunc()))
		return y
	default:
		return nil
	}
}

func processResultResults(controllerInfo map[string]ControllerInfoInterface, resI interface{}) []mergedResultInterface {
	switch res := resI.(type) {
	case []GinResult:
		var x = make(map[string]*mergedResult)
		for r := range res {
			rr := res[r]
			pc := processResultResult(rr)
			var mr = x[rr.GetPath()]
			if mr == nil {
				mr = new(mergedResult)
				mr.Path = pc.GetPath()
				if ci, ok := controllerInfo[mr.Path]; ok {
					mr.Cate = ci.GetCategory()
					mr.Description = ci.GetDescription()
				}
				if len(mr.Cate) == 0 {
					mr.Cate = pc.GetTitle()
				}
			}
			mr.Results = append(mr.Results, pc)
			mr.Description += fmt.Sprintf("\n%v: %v", pc.GetMethod(), pc.GetDescription())

			ct := pc.GetCategory()
			if len(mr.Cate) == 0 && len(ct) != 0 {
				mr.Cate = ct
			}

			x[rr.GetPath()] = mr
		}
		var keys []string
		var y = make([]mergedResultInterface, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for i := range keys {
			y = append(y, x[keys[i]])
		}
		return y
	default:
		return nil
	}
}


type ControllerInfoInterface interface {
	GetCategory() string
	GetPath() string
	GetDescription() string
}


type ControllerInfo struct {
	Category string
	Path string
	Description string
}

func (c ControllerInfo) GetCategory() string {
	return c.Category
}

func (c ControllerInfo) GetPath() string {
	return c.Path
}

func (c ControllerInfo) GetDescription() string {
	return c.Description
}

type providerInfo struct {
	DocName string
	Description string
}

type ProviderInfoInterface interface {
	GetDocName() string
	GetDescription() string
}

func (c providerInfo) GetDocName() string {
	return c.DocName
}

func (c providerInfo) GetDescription() string {
	return c.Description
}

func parseController(i interface{}) ControllerInfoInterface {
	segs := strings.Split(parser.InterfaceDescription(i), "@")
	var res = new(ControllerInfo)
	for i := range segs {
		kv := strings.SplitN(strings.TrimSpace(segs[i]), " ", 2)
		if len(kv) == 1 {
			continue
		}
		k, v := strings.ToLower(strings.TrimSpace(kv[0])), strings.TrimSpace(kv[1])
		switch k {
		case "description":
			res.Description = v
		case "category":
			res.Category = v
		case "path":
			res.Path = v
		}
	}
	return res
}

func parseControllers(provider ControllerProvider) (mp map[string]ControllerInfoInterface) {
	mp = make(map[string]ControllerInfoInterface)
	if provider == nil {
		return
	}
	controllers := provider.GetControllers()
	for i := range controllers {
		ci := parseController(controllers[i])
		mp[ci.GetPath()] = ci
	}
	return
}

func parseProvider(i interface{}) ProviderInfoInterface {
	segs := strings.Split(parser.InterfaceDescription(i), "@")
	var res = new(providerInfo)
	for i := range segs {
		kv := strings.SplitN(strings.TrimSpace(segs[i]), " ", 2)
		if len(kv) == 1 {
			continue
		}
		k, v := strings.ToLower(strings.TrimSpace(kv[0])), strings.TrimSpace(kv[1])
		switch k {
		case "description":
			res.Description = v
		case "docname":
			res.DocName = v
		}
	}
	return res
}

func processResult(resI interface{}) *ginProcessedInfo {
	switch res := resI.(type) {
	case *GinInfo:
		var x = new(ginProcessedInfo)
		x.Host = res.Host
		if res.ControllerProvider != nil {
			x.ProviderInfoInterface = parseProvider(res.ControllerProvider.GetProvider())
		} else {
			x.ProviderInfoInterface = providerInfo{
				DocName:     res.DocName,
				Description: res.Description,
			}
		}
		x.Categories = processResultResults(
			parseControllers(res.ControllerProvider), res.Result)
		fmt.Println(x)
		return x
	default:
		return &ginProcessedInfo{}
	}
}
