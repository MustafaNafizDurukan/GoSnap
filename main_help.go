package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/mustafanafizdurukan/GoSnap/pkg/collect"
	"github.com/mustafanafizdurukan/GoSnap/pkg/counter"
	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

var (
	snapshots    = sync.Map{}
	newProcesses = make([]*types.ProcessInfo, 0)
)

// takeSnapshots take snapshots during given time and stores them into snapshots map.
func takeSnapshots(givenSeconds time.Duration) {
	var err error

	// Take baseline snapshot
	baseline, err := collect.Processes()
	if err != nil {
		fmt.Println("Error taking baseline snapshot:", err)
		os.Exit(1)
	}
	snapshots.Store(0, baseline)

	var i int
	counter.Start(givenSeconds, func() {
		takeSnapshot(i, baseline)
		i++
	})
}

// takeSnapshot takes snapshot of current processes
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

// contians checks whether given item is in baseline or not.
func contains(baseline []*types.ProcessInfo, item *types.ProcessInfo) bool {
	for _, baselineItem := range baseline {
		if baselineItem.PID == item.PID &&
			baselineItem.CreateTime == item.CreateTime &&
			baselineItem.Cmdline == item.Cmdline {
			return true
		}
	}
	return false
}

// saveJson saves taken processes to json file
func saveJson() {
	// Creating JSON files
	processesFile, err := os.Create("processes.json")
	if err != nil {
		os.Exit(1)
	}
	defer processesFile.Close()

	baselineAny, _ := snapshots.Load(0)
	baseline := baselineAny.([]*types.ProcessInfo)

	allProcesses := make(map[string][]*types.ProcessInfo)
	allProcesses["baseline_processes"] = baseline
	allProcesses["new_processes"] = newProcesses

	// Save new processes to  JSON files
	allProcessesJSON, err := json.MarshalIndent(allProcesses, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling all processes:", err)
		os.Exit(1)
	}
	processesFile.Write(allProcessesJSON)
}
