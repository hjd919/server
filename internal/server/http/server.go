package http

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/middleware/auth"
	"github.com/hjd919/server/internal/service"
)

var (
	svc *service.Service
)

func initRouter(e *bm.Engine) {
	e.POST("/record", svc.Record)

	jishua := e.Group("/jishua/brush_idfa")
	{
		jishua.GET("/getTask", svc.JishuaGetTask)
	}

	// 后台
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

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}
