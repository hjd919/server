package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/middleware/auth"
)

func initRouter(e *bm.Engine) {
	e.POST("/record", svc.Record)

	authn := auth.New(&auth.Config{DisableCSRF: false})

	admin := e.Group("/admin", authn.User)
	{
		//获取学生列表
		admin.GET("/detail", svc.AdminDetail)
	}

	admin = e.Group("/admin")
	{
		//获取学生列表
		admin.GET("/login", svc.AdminLogin)
	}
}
