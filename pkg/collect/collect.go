package collect

import (
	"time"

	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
	"github.com/shirou/gopsutil/process"
)

func Processes() ([]*types.ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	processList := make([]*types.ProcessInfo, 0, 10)
	for _, proc := range processes {
		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		memPercent, err := proc.MemoryPercent()
		if err != nil {
			memPercent = 0
		}

		ioCounters, err := proc.IOCounters()
		if err != nil {
			ioCounters = &process.IOCountersStat{}
		}

		createTime, err := proc.CreateTime()
		if err != nil {
			createTime = 0
		}

		name, err := proc.Name()
		if err != nil {
			name = ""
		}

		username, err := proc.Username()
		if err != nil {
			username = ""
		}

		cmdline, err := proc.Cmdline()
		if err != nil {
			cmdline = ""
		}

		numThreads, err := proc.NumThreads()
		if err != nil {
			numThreads = 0
		}

		var ppid int32
		pp, err := proc.Parent()
		if err != nil {
			ppid = -1
		}
		if pp != nil {
			ppid = pp.Pid
		}

		processList = append(processList, &types.ProcessInfo{
			PID:             proc.Pid,
			PPID:            ppid,
			Name:            name,
			Cmdline:         cmdline,
			Username:        username,
			CPU:             cpuPercent,
			Memory:          float64(memPercent),
			NumThreads:      numThreads,
			CreateTime:      time.Unix(int64(createTime/1000), 0),
			IOCountersRead:  ioCounters.ReadBytes,
			IOCountersWrite: ioCounters.WriteBytes,
		})

	}
	return processList, nil
}
