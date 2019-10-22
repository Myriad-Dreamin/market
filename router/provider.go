package router

import (
	"fmt"
	"github.com/Myriad-Dreamin/ginx/types"
	"path"
)

type Provider struct {
	types.BaseModuler
	rootRouter *RootRouter
	objectRouter *ObjectRouter
	userRouter *UserRouter
	goodsRouter *GoodsRouter
	needsRouter *NeedsRouter
}

func NewProvider(namespace string) *Provider {
	return &Provider{
		BaseModuler: types.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *Provider) Register(name string, router interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), router); err != nil {
		panic(fmt.Errorf("unknown router %T", router))
	}
	switch ss := router.(type) {
	case *RootRouter:
		s.rootRouter = ss
	case *ObjectRouter:
		s.objectRouter = ss
	case *UserRouter:
		s.userRouter = ss
	default:
		panic(fmt.Errorf("unknown router %T", router))
	}
}

func (s *Provider) Replace(name string, router interface{}) {
	if err := s.BaseModuler.Replace(path.Join(s.Namespace, name), router); err != nil {
		panic(fmt.Errorf("unknown router %T", router))
	}
	switch ss := router.(type) {
	case *RootRouter:
		s.rootRouter = ss
	case *ObjectRouter:
		s.objectRouter = ss
	case *UserRouter:
		s.userRouter = ss
	default:
		panic(fmt.Errorf("unknown router %T", router))
	}
}

func (s *Provider) RootRouter() *RootRouter {
	return s.rootRouter
}
