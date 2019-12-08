package model

import (
	"github.com/Myriad-Dreamin/market/config"
	dblayer "github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type StatFee = dblayer.StatFee
type StatFeeDB = dblayer.StatFeeDB
type StatFeeRequest = dblayer.StatFeeRequest
type StatFeeXYResult = dblayer.StatFeeXYResult
type StatFeeCountXYResult = dblayer.StatFeeCountXYResult

func NewStatFeeDB(logger types.Logger, cfg *config.ServerConfig) (*StatFeeDB, error) {
	return dblayer.NewStatFeeDB(logger, cfg)
}

func GetStatFeeDB(logger types.Logger, cfg *config.ServerConfig) (*StatFeeDB, error) {
	return dblayer.GetStatFeeDB(logger, cfg)
}
