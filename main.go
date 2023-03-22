package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func snapshotWithTime() ([]*ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	processList := make([]*ProcessInfo, 0)
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

		processList = append(processList, &ProcessInfo{
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

func contains(slice []*ProcessInfo, item *ProcessInfo) bool {
	for _, sliceItem := range slice {
		if sliceItem.Process.Pid == item.Process.Pid &&
			sliceItem.CreateTime == item.CreateTime &&
			sliceItem.Cmdline == item.Cmdline {
			return true
		}
	}
	return false
}

func main() {
	snapshots := make(map[int][]*ProcessInfo)
	newProcesses := make([]*ProcessInfo, 0)
	var err error

	// Take baseline snapshot
	snapshots[0], err = snapshotWithTime()
	if err != nil {
		fmt.Println("Error taking baseline snapshot:", err)
		os.Exit(1)
	}

	// Loop until time limit is reached
	timeLimit := 5 // in seconds
	for i := 1; i <= timeLimit; i++ {
		// Take snapshot
		snapshot, err := snapshotWithTime()
		if err != nil {
			fmt.Println("Error taking snapshot:", err)
			continue
		}

		// Add snapshot to snapshots map
		snapshots[i] = snapshot

		// Compare snapshot to baseline snapshot
		for _, item := range snapshot {
			if !contains(snapshots[0], item) {
				// If process is not in baseline snapshot, add to newProcesses map
				newProcesses = append(newProcesses, item)
			}
		}

		// Sleep for 1 second before taking next snapshot
		time.Sleep(time.Second)
	}

	// Print new processes map as JSON
	newProcessesJSON, err := json.MarshalIndent(newProcesses, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling new processes:", err)
		os.Exit(1)
	}
	fmt.Println(string(newProcessesJSON))
}
