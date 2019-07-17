package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/hjd919/server/internal/server/http"
	"github.com/hjd919/server/internal/service"
)

func main() {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}

	logConfig := loadLogConfig()
	log.Init(logConfig) // debug flag∏: log.dir={path}
	defer log.Close()

	log.Info("kratos-demo start")
	svc := service.New()
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}
			log.Info("kratos-demo exit")
			// svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func loadLogConfig() *log.Config {
	var (
		log struct {
			Log *log.Config
		}
	)
	if err := paladin.Get("log.toml").UnmarshalTOML(&log); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	return log.Log
}
