package mock

import (
	"bytes"
	"io"
	"reflect"
)

type Serializable interface {
	Serialize() (io.Reader, error)
	ContentType() string
}


func NotStruct(content interface{}) io.Reader {
	var buf = new(bytes.Buffer)
	switch ss := content.(type) {
	case string:
		_, err := buf.WriteString(ss)
		if err != nil {
			panic(err)
		}
	case []byte:
		_, err := buf.Write(ss)
		if err != nil {
			panic(err)
		}
	case reflect.Value:
		switch ss.Kind() {
		case reflect.Ptr, reflect.Interface:
			return NotStruct(ss.Elem())
		case reflect.Struct, reflect.Slice, reflect.Map:
			return nil
		}
	default:
		return NotStruct(reflect.ValueOf(content))
	}
	return buf
}

