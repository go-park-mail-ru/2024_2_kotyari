package grpc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/disk"
	"log"
	"time"
)

func updateDiskMetrics(total, free *prometheus.GaugeVec) {
	paths := []string{"/"}
	for {
		for _, path := range paths {
			diskStat, err := disk.Usage(path)
			if err != nil {
				log.Printf("Error fetching disk metrics for %s: %v", path, err)
				continue
			}

			total.WithLabelValues(path).Set(float64(diskStat.Total))
			free.WithLabelValues(path).Set(float64(diskStat.Free))
		}
		time.Sleep(10 * time.Second)
	}
}
