package model

import (
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type StatFee = dblayer.StatFee
type StatFeeDB = dblayer.StatFeeDB
type StatFeeRequest = dblayer.StatFeeRequest
type StatFeeXYResult = dblayer.StatFeeXYResult
type StatFeeCountXYResult = dblayer.StatFeeCountXYResult

func NewStatFeeDB(m module.Module) (*StatFeeDB, error) {
	return dblayer.NewStatFeeDB(m)
}

func GetStatFeeDB(m module.Module) (*StatFeeDB, error) {
	return dblayer.GetStatFeeDB(m)
}
