package server

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/go-core/bean"
)

func CpuMem(c *gin.Context) {
	Mgr.blk.RLock()
	box := Mgr.box
	Mgr.blk.RUnlock()
	c.IndentedJSON(200, box)
}
func CpuInfos(c *gin.Context) {
	Mgr.blk.RLock()
	pth := Mgr.cpuDev
	Mgr.blk.RUnlock()
	tms := time.Now()
	req, err := YwtEgn.HbtpGrpcRequest(pth, 1, "CpuInfos")
	if err != nil {
		c.String(500, "req err:%v", err)
		return
	}
	defer req.Close()
	err = req.Do(Mgr.Ctx, nil)
	if err != nil {
		c.String(500, "req.Do err:%v", err)
		return
	}
	if req.ResCode() != hbtp.ResStatusOk {
		c.String(500, "res code(%d) err:%s", req.ResCode(), string(req.ResBodyBytes()))
		return
	}
	hbtp.Debugf("req CpuInfos times:%.4fs", time.Since(tms).Seconds())
	c.Data(200, "application/json", req.ResBodyBytes())
}
func procs(c *gin.Context) {
	Mgr.blk.RLock()
	pth := Mgr.cpuDev
	Mgr.blk.RUnlock()
	tms := time.Now()
	req, err := YwtEgn.HbtpGrpcRequest(pth, 1, "Process")
	if err != nil {
		c.String(500, "req err:%v", err)
		return
	}
	defer req.Close()
	err = req.Do(Mgr.Ctx, nil)
	if err != nil {
		c.String(500, "req.Do err:%v", err)
		return
	}
	if req.ResCode() != hbtp.ResStatusOk {
		c.String(500, "res code(%d) err:%s", req.ResCode(), string(req.ResBodyBytes()))
		return
	}
	hbtp.Debugf("req Process times:%.4fs", time.Since(tms).Seconds())
	c.Data(200, "application/json", req.ResBodyBytes())
}

func getWarnln(args url.Values) (string, error) {
	req, err := YwtEgn.HbtpGrpcRequest(bean.NewCliGroupPath("mgr", "cpu-report"), 1, "GetWarnLen", args)
	if err != nil {
		return "", err
	}
	defer req.Close()
	err = req.Do(Mgr.Ctx, nil)
	if err != nil {
		return "", err
	}
	if req.ResCode() != hbtp.ResStatusOk {
		return "", fmt.Errorf("res code(%d) err:%s", req.ResCode(), string(req.ResBodyBytes()))
	}
	return string(req.ResBodyBytes()), nil
}
func warnln(c *gin.Context) {
	ns1, err := getWarnln(nil)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	pars := url.Values{}
	pars.Set("day", "-1")
	ns2, err := getWarnln(pars)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	pars = url.Values{}
	pars.Set("day", "1")
	ns3, err := getWarnln(pars)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	pars = url.Values{}
	pars.Set("day", "2")
	ns4, err := getWarnln(pars)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, hbtp.Map{
		"count":         ns1,
		"count_today":   ns2,
		"count_lastday": ns3,
		"count_3day":    ns4,
	})
}
func warns(c *gin.Context) {
	req, err := YwtEgn.HbtpGrpcRequest(bean.NewCliGroupPath("mgr", "cpu-report"), 1, "GetWarns")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	defer req.Close()
	err = req.Do(Mgr.Ctx, nil)
	if err != nil {
		c.String(500, "req.Do err:%v", err)
		return
	}
	if req.ResCode() != hbtp.ResStatusOk {
		c.String(500, "res code(%d) err:%s", req.ResCode(), string(req.ResBodyBytes()))
		return
	}
	c.Data(200, "application/json", req.ResBodyBytes())
}
func winfoln(c *gin.Context) {
	req, err := YwtEgn.HbtpGrpcRequest(bean.NewCliGroupPath("mgr", "cpu-report"), 1, "GetInfoLen")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	defer req.Close()
	err = req.Do(Mgr.Ctx, nil)
	if err != nil {
		c.String(500, "req.Do err:%v", err)
		return
	}
	if req.ResCode() != hbtp.ResStatusOk {
		c.String(500, "res code(%d) err:%s", req.ResCode(), string(req.ResBodyBytes()))
		return
	}
	c.JSON(200, hbtp.Map{
		"count": string(req.ResBodyBytes()),
	})
}
func winfos(c *gin.Context) {
	req, err := YwtEgn.HbtpGrpcRequest(bean.NewCliGroupPath("mgr", "cpu-report"), 1, "GetInfos")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	defer req.Close()
	err = req.Do(Mgr.Ctx, nil)
	if err != nil {
		c.String(500, "req.Do err:%v", err)
		return
	}
	if req.ResCode() != hbtp.ResStatusOk {
		c.String(500, "res code(%d) err:%s", req.ResCode(), string(req.ResBodyBytes()))
		return
	}
	c.Data(200, "application/json", req.ResBodyBytes())
}
