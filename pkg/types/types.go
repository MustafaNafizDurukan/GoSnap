package types

import (
	"time"

	"github.com/shirou/gopsutil/process"
)

type ProcessInfo struct {
	Process         *process.Process        // Pointer to the process struct
	Pid             int32                   // Process ID
	Name            string                  // Process name
	Cmdline         string                  // Command-line arguments
	Username        string                  // Username of the user who started the process
	CPU             float64                 // CPU usage (percentage) by the process
	Memory          float64                 // Memory usage (percentage) by the process
	NumThreads      int32                   // Number of threads spawned by the process
	CreateTime      time.Time               // Process create time
	IOCountersRead  uint64                  // Number of bytes read by the process
	IOCountersWrite uint64                  // Number of bytes written by the process
	MemoryInfo      *process.MemoryInfoStat // Memory information for the process
}
