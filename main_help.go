package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/mustafanafizdurukan/GoSnap/pkg/collect"
	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

var (
	snapshots    = sync.Map{} //make(map[int][]*types.ProcessInfo)
	newProcesses = make([]*types.ProcessInfo, 0)
)

func takeSnapshots(givenSeconds int) {
	var err error

	// Take baseline snapshot
	baseline, err := collect.Processes()
	if err != nil {
		fmt.Println("Error taking baseline snapshot:", err)
		os.Exit(1)
	}
	snapshots.Store(0, baseline)

	// Loop until time limit is reached
	for i := 1; i <= givenSeconds; i++ {
		go takeSnapshot(i, baseline)

		// Sleep for 1 second before taking next snapshot
		time.Sleep(time.Second)
	}
}

func takeSnapshot(i int, baseline []*types.ProcessInfo) {
	// Take snapshot
	snapshot, err := collect.Processes()
	if err != nil {
		fmt.Println("Error taking snapshot:", err)
		return
	}

	// Add snapshot to snapshots map
	snapshots.Store(i, snapshot)

	// Compare snapshot to baseline snapshot
	for _, item := range snapshot {
		if !contains(baseline, item) {
			// If process is not in baseline snapshot, add to newProcesses map
			newProcesses = append(newProcesses, item)
		}
	}
}

func contains(baseline []*types.ProcessInfo, item *types.ProcessInfo) bool {
	for _, baselineItem := range baseline {
		if baselineItem.Process.Pid == item.Process.Pid &&
			baselineItem.CreateTime == item.CreateTime &&
			baselineItem.Cmdline == item.Cmdline {
			return true
		}
	}
	return false
}

func saveJson() {
	// Creating JSON files
	allProcessesFile, err := os.Create("all_processes.json")
	if err != nil {
		fmt.Println("Error creating all_processes.json file:", err)
		os.Exit(1)
	}
	defer allProcessesFile.Close()

	newProcessesFile, err := os.Create("new_processes.json")
	if err != nil {
		fmt.Println("Error creating new_processes.json file:", err)
		os.Exit(1)
	}
	defer newProcessesFile.Close()

	snapshotMap := make(map[int][]*types.ProcessInfo)
	snapshots.Range(func(key, value interface{}) bool {
		snapshotMap[key.(int)] = value.([]*types.ProcessInfo)
		return true
	})

	// Save new processes to  JSON files
	allProcessesJSON, err := json.MarshalIndent(snapshotMap, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling all processes:", err)
		os.Exit(1)
	}
	allProcessesFile.Write(allProcessesJSON)

	newProcessesJSON, err := json.MarshalIndent(newProcesses, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling new processes:", err)
		os.Exit(1)
	}
	newProcessesFile.Write(newProcessesJSON)
}
