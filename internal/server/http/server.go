package http

import (
	"github.com/bilibili/kratos/pkg/conf/paladin"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/hjd919/server/internal/service"
)

var (
	svc *service.Service
)

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
