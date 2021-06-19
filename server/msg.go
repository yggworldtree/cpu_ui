package server

import (
	"encoding/json"

	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/cpu_ui/comm"
	"github.com/yggworldtree/go-core/messages"
	"github.com/yggworldtree/go-sdk/ywtree"
)

func (c *Manager) OnConnect(egn *ywtree.Engine) {
	go Mgr.startReg()
}
func (c *Manager) OnDisconnect(egn *ywtree.Engine) {

}
func (c *Manager) OnMessage(egn *ywtree.Engine, msg *ywtree.MessageTopic) *messages.ReplyInfo {
	pths := msg.Path.String()
	hbtp.Debugf("OnMessage:%s,from:%s", pths, msg.Sender.String())
	switch pths {
	case comm.MsgPthCpuMem.String():
		c.blk.Lock()
		c.cpuDev = msg.Sender
		err := json.Unmarshal(msg.Body, &c.box)
		c.blk.Unlock()
		if err != nil {
			hbtp.Debugf("OnMessage:%s json err:%v", pths, err)
		}
	}
	return nil
}
func (c *Manager) OnBroadcast(egn *ywtree.Engine, msg *messages.MessageBox) *messages.ReplyInfo {

	return nil
}
