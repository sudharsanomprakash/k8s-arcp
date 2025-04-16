package internal

import (
	"fmt"
	"os"
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metrics struct {
	CPUUsage float64
	MemUsage uint32
	Timestamp int64
}

func Collect() (*Metrics, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &Metrics{
		CPUUsage:   cpuPercent[0],
		MemUsage:   uint32(vmStat.Used / 1024 / 1024),
		Timestamp:  time.Now().Unix(),
	}, nil
}
