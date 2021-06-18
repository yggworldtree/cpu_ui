package server

import "github.com/gin-gonic/gin"

func CpuMem(c *gin.Context) {
	Mgr.blk.RLock()
	box := Mgr.box
	Mgr.blk.RUnlock()
	c.String(200, "name:%s, cpus:%.5f%%", box.Name, box.Cpu.Average)
}
