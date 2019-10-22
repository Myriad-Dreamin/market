package model

import (
	"github.com/Myriad-Dreamin/ginx/config"
	splayer "github.com/Myriad-Dreamin/ginx/model/sp-layer"
	"github.com/Myriad-Dreamin/ginx/types"
)

type Needs = splayer.Needs
type NeedsDB = splayer.NeedsDB

func NewNeedsDB(logger types.Logger, cfg *config.ServerConfig) (*NeedsDB, error) {
	return splayer.NewNeedsDB(logger, cfg)
}

func GetNeedsDB(logger types.Logger, cfg *config.ServerConfig) (*NeedsDB, error) {
	return splayer.GetNeedsDB(logger, cfg)
}