package server

import (
	"github.com/gin-gonic/gin"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/go-sdk/ywtree"
)

var (
	Web    = gin.New()
	Mgr    *Manager
	YwtEgn *ywtree.Engine
)

func Run() {
	hbtp.Debug = true
	Mgr = NewManager()
	YwtEgn = ywtree.NewEngine(Mgr, &ywtree.Config{
		Host: "localhost:7000",
		Org:  "mgr",
		Name: "cpu-ui",
	})
	go runWeb()
	err := YwtEgn.Run()
	if err != nil {
		return
	}
	Mgr.cncl()
}

func runWeb() {
	Web.GET("/cpu-mem", CpuMem)
	err := Web.Run(":8080")
	if err != nil {
		hbtp.Debugf("gin run err:%v", err)
	}
	if YwtEgn != nil {
		YwtEgn.Stop()
	}
}
