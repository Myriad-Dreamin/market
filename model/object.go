package model

import (
	"github.com/Myriad-Dreamin/market/config"
	splayer "github.com/Myriad-Dreamin/market/model/sp-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type Object = splayer.Object
type ObjectDB = splayer.ObjectDB

func NewObjectDB(logger types.Logger, cfg *config.ServerConfig) (*ObjectDB, error) {
	return splayer.NewObjectDB(logger, cfg)
}

func GetObjectDB(logger types.Logger, cfg *config.ServerConfig) (*ObjectDB, error) {
	return splayer.GetObjectDB(logger, cfg)
}
