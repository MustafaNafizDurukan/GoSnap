package collect

import (
	"fmt"
	"time"

	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
	"github.com/shirou/gopsutil/process"
)

func Processes() ([]*types.ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	processList := make([]*types.ProcessInfo, 0)
	for _, proc := range processes {
		cpuPercent, err := proc.CPUPercent()
		if err != nil {
			fmt.Printf("CPUPercent could not have taken: %v \r\n", err)
			cpuPercent = 0
		}
		memPercent, err := proc.MemoryPercent()
		if err != nil {
			fmt.Printf("MemoryPercent could not have taken: %v \r\n", err)
			memPercent = 0
		}
		ioCounters, err := proc.IOCounters()
		if err != nil {
			fmt.Printf("IOCounters could not have taken: %v \r\n", err)
			ioCounters = &process.IOCountersStat{}
		}
		createTime, err := proc.CreateTime()
		if err != nil {
			fmt.Printf("CreateTime could not have taken: %v \r\n", err)
			createTime = 0
		}
		name, err := proc.Name()
		if err != nil {
			fmt.Printf("Name could not have taken: %v \r\n", err)
			name = ""
		}
		username, err := proc.Username()
		if err != nil {
			fmt.Printf("Username could not have taken: %v \r\n", err)
			username = ""
		}
		cmdline, err := proc.Cmdline()
		if err != nil {
			fmt.Printf("Cmdline could not have taken: %v \r\n", err)
			cmdline = ""
		}
		numThreads, err := proc.NumThreads()
		if err != nil {
			fmt.Printf("NumThreads could not have taken: %v \r\n", err)
			numThreads = 0
		}
		memInfo, err := proc.MemoryInfo()
		if err != nil {
			fmt.Printf("MemoryInfo could not have taken: %v \r\n", err)
			memInfo = &process.MemoryInfoStat{}
		}

		processList = append(processList, &types.ProcessInfo{
			Process:         proc,
			Pid:             proc.Pid,
			Name:            name,
			Cmdline:         cmdline,
			Username:        username,
			CPU:             cpuPercent,
			Memory:          float64(memPercent),
			NumThreads:      numThreads,
			CreateTime:      time.Unix(int64(createTime/1000), 0),
			IOCountersRead:  ioCounters.ReadBytes,
			IOCountersWrite: ioCounters.WriteBytes,
			MemoryInfo:      memInfo,
		})

	}
	return processList, nil
}
