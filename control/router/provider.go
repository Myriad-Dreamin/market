package router

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"path"
)

type Provider struct {
	module.BaseModuler

	objectRouter *ObjectRouter

	rootRouter      *RootRouter
	authRouter      *AuthRouter
	statisticRouter *StatisticRouter
	userRouter      *UserRouter
	goodsRouter     *GoodsRouter
	needsRouter     *NeedsRouter
	constRouter     *ConstRouter
}

func NewProvider(namespace string) *Provider {
	return &Provider{
		BaseModuler: module.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *Provider) Register(name string, router interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), router); err != nil {
		panic(fmt.Errorf("unknown router %T", router))
	}
	switch ss := router.(type) {
	case *ConstRouter:
		s.constRouter = ss
	case *AuthRouter:
		s.authRouter = ss
	case *RootRouter:
		s.rootRouter = ss
	case *GoodsRouter:
		s.goodsRouter = ss
	case *NeedsRouter:
		s.needsRouter = ss
	case *UserRouter:
		s.userRouter = ss
	case *StatisticRouter:
		s.statisticRouter = ss
	default:
		panic(fmt.Errorf("unknown router %T", router))
	}
}

func (s *Provider) Replace(name string, router interface{}) {
	if err := s.BaseModuler.Replace(path.Join(s.Namespace, name), router); err != nil {
		panic(fmt.Errorf("unknown router %T", router))
	}
	switch ss := router.(type) {
	case *ConstRouter:
		s.constRouter = ss
	case *AuthRouter:
		s.authRouter = ss
	case *RootRouter:
		s.rootRouter = ss
	case *GoodsRouter:
		s.goodsRouter = ss
	case *NeedsRouter:
		s.needsRouter = ss
	case *UserRouter:
		s.userRouter = ss
	case *StatisticRouter:
		s.statisticRouter = ss
	default:
		panic(fmt.Errorf("unknown router %T", router))
	}
}

func (s *Provider) RootRouter() *RootRouter {
	return s.rootRouter
}
