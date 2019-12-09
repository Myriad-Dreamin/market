package auth

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/casbin/casbin/v2"
	"strconv"
)

type objectEntity struct {}

func (objectEntity) CreateObj(groupID uint) string {
	return "object:" + strconv.Itoa(int(groupID))
}

func (objectEntity) Read() controller.Requirement {
	return controller.Requirement{Entity: "object", Action: "r/.*"}
}

func (objectEntity) AddReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ObjectEntity.CreateObj(aim), "r/.*")
}

func (objectEntity) RemoveReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ObjectEntity.CreateObj(aim), "r/.*")
}

func (objectEntity) HasReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ObjectEntity.CreateObj(aim), "r/.*")
}

func (objectEntity) Write() controller.Requirement {
	return controller.Requirement{Entity: "object", Action: "w/.*"}
}

func (objectEntity) AddWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ObjectEntity.CreateObj(aim), "w/.*")
}

func (objectEntity) RemoveWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ObjectEntity.CreateObj(aim), "w/.*")
}

func (objectEntity) HasWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ObjectEntity.CreateObj(aim), "w/.*")
}

func (objectEntity) JustSimpleWrite() controller.Requirement {
	return controller.Requirement{Entity: "object", Action: "w/s"}
}

func (objectEntity) AddJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, ObjectEntity.CreateObj(aim), "w/s")
}

func (objectEntity) RemoveJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, ObjectEntity.CreateObj(aim), "w/s")
}

func (objectEntity) HasJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, ObjectEntity.CreateObj(aim), "w/s")
}
