package model

import (
	"github.com/Myriad-Dreamin/market/config"
	splayer "github.com/Myriad-Dreamin/market/model/sp-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type Needs = splayer.Needs
type NeedsDB = splayer.NeedsDB

func NewNeedsDB(logger types.Logger, cfg *config.ServerConfig) (*NeedsDB, error) {
	return splayer.NewNeedsDB(logger, cfg)
}

func GetNeedsDB(logger types.Logger, cfg *config.ServerConfig) (*NeedsDB, error) {
	return splayer.GetNeedsDB(logger, cfg)
}
