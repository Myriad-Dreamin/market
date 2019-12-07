package dblayer

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/types"
	"github.com/jinzhu/gorm"
)

func rollback(tx *gorm.DB) {
	tx.Rollback()
	if tx.Error != nil {
		fmt.Println("rollback error", tx.Error)
	}
}

func commitOrRollback(tx *gorm.DB) (int, string) {
	err := tx.Commit().Error
	if err != nil {
		rollback(tx)
		return types.CodeCommitTransactionError, err.Error()
	} else {
		return types.CodeOK, ""
	}
}

