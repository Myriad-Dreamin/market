package dblayer

import (
	"github.com/Myriad-Dreamin/dorm"
	"github.com/Myriad-Dreamin/market/types"
)

type L struct {
	types.Logger
}

func (l *L) With(kvs ...interface{}) dorm.Logger {
	return &L{l.Logger.With(kvs)}
}

func adapt(logger types.Logger) dorm.Logger {
	return &L{logger}
}

