package model

import (
	"github.com/Myriad-Dreamin/market/config"
	splayer "github.com/Myriad-Dreamin/market/model/sp-layer"
	"github.com/Myriad-Dreamin/market/types"
)

type User = splayer.User
type UserDB = splayer.UserDB

func NewUserDB(logger types.Logger, cfg *config.ServerConfig) (*UserDB, error) {
	return splayer.NewUserDB(logger, cfg)
}

func GetUserDB(logger types.Logger, cfg *config.ServerConfig) (*UserDB, error) {
	return splayer.GetUserDB(logger, cfg)
}