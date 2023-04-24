package types

import (
	"time"

	"github.com/shirou/gopsutil/process"
)

type ProcessType string

const (
	BaselineProcess ProcessType = "Baseline"
	NewProcess      ProcessType = "New"
)

type ProcessInfo struct {
	Type            ProcessType             // Type of process. It can be whether baseline or new
	PID             int32                   // Process ID
	PPID            int32                   // Parent Process ID
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
