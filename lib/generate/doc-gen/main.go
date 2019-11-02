//go:generate package-attach-to-path -generate_register_map
package doc_gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Myriad-Dreamin/market/lib/mock"
	"github.com/Myriad-Dreamin/market/lib/sugar"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"
)

const ContentTypeKey = "Content-Type"


/*
type ResultsI interface {
	GetMethod() string
	GetPath() string
	GetRecords() []RecordsI
}
*/
type Results = mock.ResultsI
type GinResult = mock.GinResultI
/*
type RecordsI interface {
	GetRequestHeader() http.Header
	GetRequestBody() []byte
	GetResponseHeader() http.Header
	GetResponseBody() []byte
}
*/
type Records = mock.RecordsI

type handler struct {
	io.Writer
}

func printf(w io.Writer, format string, args ...interface{}) {
	_, err := fmt.Fprintf(w, format, args...)
	if err != nil {
		panic(err)
	}
}

func (h handler) printf(format string, args ...interface{}) {
	printf(h, format, args...)
	return
}

func (h handler) linebreak() {
	printf(h, "\n")
}

const indentStyle = "    "
var selector, doter, indentS = []byte(": "), []byte(", "), indentStyle

func makeIndent(indent int) []byte {
	var indenter = make([]byte, indent * len(indentStyle))
	for i := 0; i < indent; i++ {
		copy(indenter[i*len(indentStyle):], indentStyle)
	}
	return indenter
}
func prettifyHeader(h http.Header, indent int) string {
	indenter := makeIndent(indent)
	var buf = bytes.NewBuffer(nil)
	var f = false
	var keys []string
	for k := range h {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		if f {
			buf.WriteByte('\n')
		} else {
			f = true
		}
		buf.Write(indenter)
		buf.WriteString(keys[i])
		buf.Write(selector)
		vs := h[keys[i]]

		if keys[i] == ContentTypeKey && (len(vs) == 0 || len(vs[0]) == 0) {
			buf.WriteString("text/plain")
		}

		for i := range vs {
			if i != 0 {
				buf.Write(doter)
			}
			buf.WriteString(vs[i])
		}
	}
	return buf.String()
}

func prettifyBody(content []byte, contentType string, indent int) string {
	indenter := makeIndent(indent)
	switch {
	case strings.HasPrefix(contentType, mock.ContentTypeJSON):
		var buf = bytes.NewBuffer(nil)
		if len(content) != 0 && content[0] == '{' {
			buf.Write(indenter)
		}
		err := json.Indent(buf, content, string(indenter), indentS)
		if err != nil {
			panic(err)
		}
		return buf.String()
	default:
		return string(indenter) + string(content)
	}
}


var resultTemplate = template.New("resultTemplate")


func FromGinResults(res *GinInfo, options... interface{}) (err error) {
	var parsedRes *ginProcessedInfo
	parsedRes, err = processResult(res)
	if err != nil {
		return
	}

	var outputPath = "result.md"

	for i := range options {
		switch option := options[i].(type) {
		case string:
			outputPath = option
		}
	}
	sugar.WithWriteFile(func(outputFile *os.File) {
		err = resultTemplate.Execute(outputFile, parsedRes)
	}, outputPath)
	return
}

func title(result GinResult) string {
	tp := strings.Split(result.GetHandler(), "/")
	return tp[len(tp) - 1]
}

func subtitle(index int, result GinResult) string {
	return result.GetHandler()
}

func contentType(h http.Header) string {
	content := h[ContentTypeKey]
	if len(content) == 0 || len(content[0]) == 0 {
		return "text/plain"
	}
	return strings.Join(content, string(doter))
}

func init() {
	resultTemplate = resultTemplate.Funcs(map[string]interface{} {
		"title": title,
		"subtitle": subtitle,
		"prettifyHeader": prettifyHeader,
		"prettifyBody": prettifyBody,
		"contentType": contentType,
	})
	var err error
	resultTemplate, err = resultTemplate.Parse(`FORMAT: 1A
HOST: {{.Host}}

# {{.GetDocName}}
{{.GetDescription}}
{{range $i, $cat := .Categories}}
## {{$cat.GetCategoryName}} [{{$cat.GetPath}}]
{{$cat.GetDescription}}

{{range $j, $v := $cat.GetResults}} + {{$v.GetMethod}}: {{$v.GetDescription}}
{{end}}
{{range $j, $v := $cat.GetResults}}{{$recs := $v.GetRecords}}{{if ne (len $recs) 0}}
### {{$v.GetTitle}} [{{$v.GetMethod}}]
{{range $k, $rec := $recs}}{{$reqType := contentType $rec.GetRequestHeader}}{{$resType := contentType $rec.GetResponseHeader}}
+ Request {{$rec.GetComment}}

    + Headers

{{prettifyHeader $rec.GetRequestHeader 3}}

    + Body

{{prettifyBody $rec.GetRequestBody $reqType 3}}

+ Response {{$rec.GetResponseCode}}

    + Headers

{{prettifyHeader $rec.GetResponseHeader 3}}

    + Body

{{prettifyBody $rec.GetResponseBody $resType 3}}
{{end}}{{end}}{{end}}{{end}}`)
	if err != nil {
		panic(err)
	}
	//sugar.WithOpenFile(func(template *os.File) {
	//	res, err := ioutil.ReadAll(template)
	//	if err != nil {
	//		panic(err)
	//	}
	//	resultTemplate, err = resultTemplate.Parse(string(res))
	//	if err != nil {
	//		panic(err)
	//	}
	//}, "doc.md.template", os.O_RDONLY, 0755)
}
