package base_service

import "github.com/gin-gonic/gin"

type CRUDEntity interface {
	Create() (int64, error)
	UpdateFields(fields []string) (int64, error)
	Delete() (int64, error)
}

type dHookObject interface {
	DeleteHook(c *gin.Context, obj CRUDEntity) bool
}

type FObject interface {
	CreateEntity(id uint) CRUDEntity
	GetEntity(id uint) (CRUDEntity, error)
}

type DObject interface {
	dHookObject
}
