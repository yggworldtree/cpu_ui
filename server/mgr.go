package server

import (
	"context"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/cpu_ui/comm"
	"github.com/yggworldtree/go-core/bean"
	"sync"
)

type Manager struct {
	Ctx    context.Context
	cncl   context.CancelFunc
	reging bool

	blk    sync.RWMutex
	box    comm.MsgBox
	cpuDev *bean.CliGroupPath
}

func NewManager() *Manager {
	c := &Manager{}
	c.Ctx, c.cncl = context.WithCancel(context.Background())
	return c
}

func (c *Manager) StartReg() {
	defer func() {
		if err := recover(); err != nil {
			hbtp.Debugf("Manager StartReg recover:%v", err)
		}
	}()
	if c.reging {
		return
	}
	c.reging = true
	defer func() {
		c.reging = false
	}()
	for !hbtp.EndContext(c.Ctx) {
		err := YwtEgn.SubTopic([]*bean.TopicInfo{
			{Path: comm.MsgPthCpuMem.String(), Safed: false},
		})
		if err == nil {
			break
		}
		hbtp.Debugf("SubTopic %s err:%v", comm.MsgPthCpuMem.String(), err)
	}
}
