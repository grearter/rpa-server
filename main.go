package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grearter/rpa-server/conf"
	"github.com/grearter/rpa-server/controller/agent"
	"github.com/grearter/rpa-server/controller/robot"
	"github.com/grearter/rpa-server/controller/user"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := conf.InitConf(); err != nil {
		return
	}

	g := gin.Default()
	user.InitRoute(g)
	agent.InitRoute(g)
	robot.InitRoute(g)

	addr := fmt.Sprintf(":%d", conf.C.Server.HttpPort)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Error("listen err: %s, addr: %s", err.Error(), addr)
		return
	}

	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt, os.Interrupt)

	httpServer := http.Server{Addr: addr, Handler: g}

	sign := <-signCh

	go func() {
		if err := httpServer.Serve(ln); err != nil {
			logrus.Error("http.Serve err: %s", err.Error())
			os.Exit(2)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	_ = httpServer.Shutdown(ctx)
	cancel()

	logrus.Infof("receive signal:%s, EXITED", sign)
	return
}
