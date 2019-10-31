package crud_dao

func (model CRUDModel) Where1(template string) func(id interface{}) (interface{}, error) {
	return func(id interface{}) (object interface{}, err error) {
		object = model.i.ObjectFactory()
		rdb := model.i.GetDB().Where(template, id).Find(object)
		err = rdb.Error
		if rdb.RecordNotFound() {
			object = nil
			err = nil
		}
		return
	}
}
