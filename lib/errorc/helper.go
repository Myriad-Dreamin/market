package errorc

import (
	"github.com/Myriad-Dreamin/market/types"
	"reflect"
)

type Code = int

func MaybeSelectError(anyObj interface{}, err error) (Code, string) {
	if err != nil {
		return types.CodeSelectError, err.Error()
	}
	if reflect.ValueOf(anyObj).IsNil() {
		return types.CodeNotFound, "not found"
	}
	return types.CodeOK, ""
}

type UpdateFieldsable interface {
	UpdateFields(fields []string) (int64, error)
}

func UpdateFields(obj UpdateFieldsable, fields []string) (Code, string) {
	_, err := obj.UpdateFields(fields)
	if err != nil {
		return types.CodeUpdateError, err.Error()
	}
	return types.CodeOK, ""
}

