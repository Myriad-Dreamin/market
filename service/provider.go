//go:generate package-attach-to-path -generate_register_map
package service

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

type SubController interface {
	GetControllers() []interface{}
}

type subController struct {
	controllers []interface{}
}

func (s subController) GetControllers() []interface{} {
	return s.controllers
}

func JustProvide(controllers ...interface{}) SubController {
	return subController{controllers: controllers}
}

// @DocName Minimum-Market
// @Description this is the market backend powered by minimum
type Provider struct {
	statisticService StatisticService
	module.BaseModuler

	objectService ObjectService
	needsService  NeedsService
	goodsService  GoodsService
	userService   UserService

	subControllers []SubController
}

func NewProvider(namespace string) *Provider {
	return &Provider{
		BaseModuler: module.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *Provider) Register(name string, service interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), service); err != nil {
		panic(fmt.Errorf("unknown/registered service %T, err %v", service, err))
	}
	if ss, ok := service.(SubController); ok {
		s.subControllers = append(s.subControllers, ss)
	}

	switch ss := service.(type) {
	case StatisticService:
		s.statisticService = ss
		s.subControllers = append(s.subControllers, JustProvide(&ss))
	case NeedsService:
		s.needsService = ss
		s.subControllers = append(s.subControllers, JustProvide(&ss))
	case GoodsService:
		s.goodsService = ss
		s.subControllers = append(s.subControllers, JustProvide(&ss))
	case UserService:
		s.userService = ss
		s.subControllers = append(s.subControllers, JustProvide(&ss))
	case ObjectService:
		s.objectService = ss
		s.subControllers = append(s.subControllers, JustProvide(&ss))
	default:
		panic(fmt.Errorf("unknown service %T", service))
	}
}

// for documents
func (s *Provider) GetControllers() []interface{} {
	var controllers []interface{}
	for i := range s.subControllers {
		controllers = append(controllers, s.subControllers[i].GetControllers()...)
	}
	return controllers
}

func (s *Provider) GetProvider() interface{} {
	return s
}
