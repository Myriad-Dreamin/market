package auth

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/casbin/casbin/v2"
	"strconv"
)

type goodsEntity struct {}

func (goodsEntity) CreateObj(groupID uint) string {
	return "goods:" + strconv.Itoa(int(groupID))
}

func (goodsEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "goods", Action: "r/.*"}
}

func (goodsEntity) AddReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, GoodsEntity.CreateObj(aim), "r/.*")
}

func (goodsEntity) RemoveReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, GoodsEntity.CreateObj(aim), "r/.*")
}

func (goodsEntity) HasReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, GoodsEntity.CreateObj(aim), "r/.*")
}

func (goodsEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "goods", Action: "w/.*"}
}

func (goodsEntity) AddWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, GoodsEntity.CreateObj(aim), "w/.*")
}

func (goodsEntity) RemoveWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, GoodsEntity.CreateObj(aim), "w/.*")
}

func (goodsEntity) HasWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, GoodsEntity.CreateObj(aim), "w/.*")
}

func (goodsEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "goods", Action: "w/s"}
}

func (goodsEntity) AddJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, GoodsEntity.CreateObj(aim), "w/s")
}

func (goodsEntity) RemoveJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, GoodsEntity.CreateObj(aim), "w/s")
}

func (goodsEntity) HasJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, GoodsEntity.CreateObj(aim), "w/s")
}
