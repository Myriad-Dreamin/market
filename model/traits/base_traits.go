package traits

import "reflect"

type BaseTraits struct {
	TypeInfo reflect.Type
	SliceInfo reflect.Type
}

func (traits BaseTraits) ObjectFactory() interface{} {
	return reflect.New(traits.TypeInfo).Interface()
}

func (traits BaseTraits) SliceFactory() interface{} {
	return reflect.MakeSlice(traits.SliceInfo, 0, 0).Interface()
}


func NewBaseTraits(t interface{}) BaseTraits {
	traits := BaseTraits{}
	traits.TypeInfo = reflect.TypeOf(t)
	traits.SliceInfo = reflect.SliceOf(traits.TypeInfo)
	return traits
}