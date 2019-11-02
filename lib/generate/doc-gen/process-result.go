package doc_gen

import (
	"errors"
	"github.com/Myriad-Dreamin/go-parse-package"
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

type GinInfo struct {
	ControllerProvider ControllerProvider
	Result             []GinResult
	Host               string
	DocName            string
	Description        string
}

type mergedResultInterface interface {
	GetPath() string
	GetCategoryName() string
	GetTitle() string
	GetDescription() string
	GetResults() []processedResultInterface
}

type mergedResult struct {
	Path        string
	Cate        string
	Description string
	Results     []processedResultInterface
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

type ginProcessedInfo struct {
	ProviderInfoInterface
	Categories []mergedResultInterface
	Host       string
}

type parsedResult struct {
	Description string
	Title       string
	Category    string
	Success     string
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

func processResultResult(resI interface{}) (processedResultInterface, error) {
	switch res := resI.(type) {
	case GinResult:
		desc, err := parser.FuncDescription(res.GetHandlerFunc())
		if err != nil {
			return nil, err
		}

		return &processedGinResult{
			GinResult:    res,
			parsedResult: parseDoc(desc),
		}, nil
	default:
		return nil, errors.New("bad result type")
	}
}

func processResultResults(controllerInfo map[string]ControllerInfoInterface, resI interface{}) ([]mergedResultInterface, error) {
	switch res := resI.(type) {
	case []GinResult:
		var x = make(map[string]*mergedResult)
		for r := range res {
			rr := res[r]
			pc, err := processResultResult(rr)
			if err != nil {
				return nil, err
			}
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
		return y, nil
	default:
		return nil, errors.New("bad results type")
	}
}

type ControllerInfoInterface interface {
	GetCategory() string
	GetPath() string
	GetDescription() string
}

type ControllerInfo struct {
	Category    string
	Path        string
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
	DocName     string
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

func parseController(i interface{}) (ControllerInfoInterface, error) {
	desc, err := parser.InterfaceDescription(i)
	if err != nil {
		return nil, err
	}
	segs := strings.Split(desc, "@")
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
	return res, nil
}

func parseControllers(provider ControllerProvider) (mp map[string]ControllerInfoInterface, err error) {
	mp = make(map[string]ControllerInfoInterface)
	if provider == nil {
		return
	}
	controllers := provider.GetControllers()
	var ci ControllerInfoInterface
	for i := range controllers {
		ci, err = parseController(controllers[i])
		if err != nil {
			mp = nil
			return
		}
		mp[ci.GetPath()] = ci
	}
	return
}

func parseProvider(i interface{}) (ProviderInfoInterface, error) {
	desc, err := parser.InterfaceDescription(i)
	if err != nil {
		return nil, err
	}
	segs := strings.Split(desc, "@")
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
	return res, nil
}

func processResult(resI interface{}) (*ginProcessedInfo, error) {
	var err error
	switch res := resI.(type) {
	case *GinInfo:
		var x = new(ginProcessedInfo)
		x.Host = res.Host
		if res.ControllerProvider != nil {
			x.ProviderInfoInterface, err = parseProvider(res.ControllerProvider.GetProvider())
			if err != nil {
				return nil, err
			}
		} else {
			x.ProviderInfoInterface = providerInfo{
				DocName:     res.DocName,
				Description: res.Description,
			}
		}
		var mp map[string]ControllerInfoInterface
		mp, err = parseControllers(res.ControllerProvider)
		if err != nil {
			return nil, err
		}
		x.Categories, err = processResultResults(
			mp, res.Result)
		return x, err
	default:
		return &ginProcessedInfo{}, errors.New("invalid result type")
	}
}
