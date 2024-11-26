package grpc

import (
	"log"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func getCPUUsage() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("Error fetching CPU metrics: %v", err)
		return 0.0
	}
	if len(percent) > 0 {
		return percent[0]
	}
	return 0.0
}

func getMemoryUsage() float64 {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("Error fetching Memory metrics: %v", err)
		return 0.0
	}
	return float64(vmStat.Used)
}
