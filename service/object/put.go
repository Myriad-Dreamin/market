package objectservice

import (
	"github.com/Myriad-Dreamin/market/lib/controller"
	"github.com/Myriad-Dreamin/market/model"
)

type PutRequest struct {
}

func (srv *Service) fillPutFields(c controller.MContext, object *model.Object, req *PutRequest) (fields []string) {
	return nil
}
