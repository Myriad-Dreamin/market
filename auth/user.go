package auth

import (
	"github.com/Myriad-Dreamin/market/lib/router"
	"github.com/casbin/casbin/v2"
"strconv"
)

type userEntity struct {}

func (userEntity) CreateObj(groupID uint) string {
	return "user:" + strconv.Itoa(int(groupID))
}

func (userEntity) Read() mgin.Requirement {
	return mgin.Requirement{Entity: "user", Action: "r/.*"}
}

func (userEntity) AddReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, UserEntity.CreateObj(aim), "r/.*")
}

func (userEntity) RemoveReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, UserEntity.CreateObj(aim), "r/.*")
}

func (userEntity) HasReadPolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, UserEntity.CreateObj(aim), "r/.*")
}

func (userEntity) Write() mgin.Requirement {
	return mgin.Requirement{Entity: "user", Action: "w/.*"}
}

func (userEntity) AddWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, UserEntity.CreateObj(aim), "w/.*")
}

func (userEntity) RemoveWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, UserEntity.CreateObj(aim), "w/.*")
}

func (userEntity) HasWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, UserEntity.CreateObj(aim), "w/.*")
}

func (userEntity) JustSimpleWrite() mgin.Requirement {
	return mgin.Requirement{Entity: "user", Action: "w/s"}
}

func (userEntity) AddJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.AddPolicy(subject, UserEntity.CreateObj(aim), "w/s")
}

func (userEntity) RemoveJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) (bool, error) {
	return e.RemovePolicy(subject, UserEntity.CreateObj(aim), "w/s")
}

func (userEntity) HasJustSimpleWritePolicy(e *casbin.SyncedEnforcer, subject string, aim uint) bool {
	return e.HasPolicy(subject, UserEntity.CreateObj(aim), "w/s")
}
