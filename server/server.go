package server

import (
	"net/http"

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
	Web.Any("/cpu-mem", CpuMem)
	Web.Any("/cpu-infos", CpuInfos)
	Web.Any("/procs", procs)
	Web.Any("/warns", warns)
	Web.Any("/winfos", winfos)
	Web.LoadHTMLGlob("view/index.html")
	Web.Static("/js", "./view/js")
	Web.Static("/css", "./view/css")
	Web.Static("/img", "./view/img")
	Web.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	err := Web.Run(":8080")
	if err != nil {
		hbtp.Debugf("gin run err:%v", err)
	}
	if YwtEgn != nil {
		YwtEgn.Stop()
	}
}
