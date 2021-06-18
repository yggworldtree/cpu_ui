package comm

import (
	"github.com/yggworldtree/go-core/bean"
)

var (
	MsgPthCpuMem  = bean.NewTopicPath("system", "cpu-mem")
	MsgPthCpuWarn = bean.NewTopicPath("system", "cpu-warn")
	MsgPthMemWarn = bean.NewTopicPath("system", "mem-warn")
)

type MsgCpuInfo struct {
	Percents []float64 `json:"percents"`
	Average  float64   `json:"average"`

	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physicalId"`
	CoreID     string   `json:"coreId"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"modelName"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cacheSize"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
}
type MsgMemInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
	Sin         uint64  `json:"sin"`
	Sout        uint64  `json:"sout"`
	PgIn        uint64  `json:"pgIn"`
	PgOut       uint64  `json:"pgOut"`
	PgFault     uint64  `json:"pgFault"`

	// Linux specific numbers
	// https://www.kernel.org/doc/Documentation/cgroup-v2.txt
	PgMajFault uint64 `json:"pgMajFault"`
}

type MsgBox struct {
	Name       string     `json:"name"`
	Cpu        MsgCpuInfo `json:"cpu"`
	SwapMem    MsgMemInfo `json:"swapMem"`
	VirtualMem MsgMemInfo `json:"virtualMem"`
}
