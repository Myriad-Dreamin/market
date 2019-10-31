package rbac

import "github.com/casbin/casbin/v2/model"

var rbacModel model.Model

func init() {
	var err error
	rbacModel, err = model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && regexMatch(r.obj,p.obj) && regexMatch(r.act,p.act)
`)
	if err != nil {
		panic(err)
	}
}
