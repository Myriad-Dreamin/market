package model

import (
	"errors"
	"github.com/Myriad-Dreamin/market/config"
	"github.com/Myriad-Dreamin/market/model/db-layer"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func Register(rdb *gorm.DB, m module.Module) error {
	return dblayer.Register(rdb, m)
}

func booleanString(b bool) string {
	if b {
		return "True"
	} else {
		return "False"
	}
}

func concatQueryString(options string) string {
	if len(options) != 0 {
		return "&"
	} else {
		return "?"
	}
}

func parseConfig(cfg *config.DatabaseConfig) (string, string, error) {
	// user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	if len(cfg.ConnectionType) == 0 || len(cfg.User) == 0 || len(cfg.Password) == 0 || len(cfg.DatabaseName) == 0 {
		return "", "", errors.New("not enough params")
	}
	url := cfg.User + ":" + cfg.Password + "@"
	if len(cfg.Host) != 0 {
		url += "(" + cfg.Host + ")"
	}
	url += "/" + cfg.DatabaseName
	options := ""

	if len(cfg.Charset) != 0 {
		options += concatQueryString(options) + "charset=" + cfg.Charset
	}
	if cfg.ParseTime {
		options += concatQueryString(options) + "parseTime=" + booleanString(cfg.ParseTime)
	}
	if len(cfg.Location) != 0 {
		options += concatQueryString(options) + "loc=" + cfg.Location
	}
	return cfg.ConnectionType, url + options, nil
}

func OpenORM(cfg *config.ServerConfig) (*gorm.DB, error) {
	dialect, args, err := parseConfig(&cfg.DatabaseConfig)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dialect, args)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MockORM(_ *config.ServerConfig) (*gorm.DB, error) {

	return gorm.Open("sqlite3", ":memory:")
}
