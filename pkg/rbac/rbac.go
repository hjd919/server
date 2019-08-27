package rbac

import "github.com/casbin/casbin"

func RbacStart() {
	// 获取Enforcer
	e := casbin.NewEnforcer("./config/model.conf", "./config/policy.csv")

	// 判断访问
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	if e.Enforce(sub, obj, act) == true {
		// permit alice to read data1
	} else {
		// deny the request, show an error
	}

	// 分配
	roles := e.GetRoles("alice")
}
