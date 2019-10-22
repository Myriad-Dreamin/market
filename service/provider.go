package service

import (
	"fmt"
)

type Provider struct {
	objectService *ObjectService
	needsService *NeedsService
	goodsService *GoodsService
	userService UserService
}

func (s *Provider) Register(service interface{}) {
	switch ss := service.(type) {
	case *NeedsService:
		s.needsService = ss
	case *GoodsService:
		s.goodsService = ss
	case UserService:
		s.userService = ss
	case *ObjectService:
		s.objectService = ss
	default:
		panic(fmt.Errorf("unknown service %T", service))
	}
}
