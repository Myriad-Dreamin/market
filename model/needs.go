package model

import (
	splayer "github.com/Myriad-Dreamin/market/model/sp-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
)

type Needs = splayer.Needs
type NeedsDB = splayer.NeedsDB

func NewNeedsDB(m module.Module) (*NeedsDB, error) {
	return splayer.NewNeedsDB(m)
}

func GetNeedsDB(m module.Module) (*NeedsDB, error) {
	return splayer.GetNeedsDB(m)
}
