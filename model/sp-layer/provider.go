package splayer

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/types"
	"path"
)

var provider *Provider


type Provider struct {
	types.BaseModuler
	objectDB *ObjectDB
	needsDB *NeedsDB
	goodsDB *GoodsDB
	userDB *UserDB
	enforcer *Enforcer
}

func NewProvider(namespace string) *Provider {
	return &Provider{
		BaseModuler: types.BaseModuler{
			Namespace: namespace,
		},
	}
}

func (s *Provider) Register(name string, database interface{}) {
	if err := s.Provide(path.Join(s.Namespace, name), database); err != nil {
		panic(fmt.Errorf("unknown database %T", database))
	}

	switch ss := database.(type) {
	case *NeedsDB:
		s.needsDB = ss
	case *Enforcer:
		s.enforcer = ss
	case *GoodsDB:
		s.goodsDB = ss
	case *UserDB:
		s.userDB = ss
	case *ObjectDB:
		s.objectDB = ss
	default:
		if mm, ok := ss.(types.Moduler); ok {
			// todo:
			_ = mm
		}
	}
}

func SetProvider(p *Provider) (op *Provider) {
	op = provider
	provider = p
	return
}

/**
if err := s.BaseModuler.Replace(path.Join(s.Namespace, name), router); err != nil {
	panic(fmt.Errorf("unknown router %T", router))
}
 */
