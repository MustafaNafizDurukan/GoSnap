package snapshottaker

import (
	"fmt"
	"os"
	"time"

	"github.com/mustafanafizdurukan/GoSnap/pkg/collect"
	"github.com/mustafanafizdurukan/GoSnap/pkg/counter"
	"github.com/mustafanafizdurukan/GoSnap/pkg/reporter"
	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

type SnapshotTaker struct {
	interval time.Duration
	reporter *reporter.Reporter

	BaselineProcesses []*types.ProcessInfo
	NewProcesses      []*types.ProcessInfo
}

func New(interval time.Duration) *SnapshotTaker {
	return &SnapshotTaker{
		interval: interval,
		reporter: reporter.New(),
	}
}

// Start take snapshots during given time and stores them into snapshots map.
func (st *SnapshotTaker) Start() {
	var err error

	// Take baseline snapshot
	st.BaselineProcesses, err = collect.Processes()
	if err != nil {
		fmt.Println("Error taking baseline snapshot:", err)
		os.Exit(1)
	}

	for i := range st.BaselineProcesses {
		st.BaselineProcesses[i].Type = types.BaselineProcess
	}
	st.reporter.Report(st.BaselineProcesses...)

	var i int
	counter.Start(st.interval, func() {
		st.takeSnapshot(i)
		i++
	})
}

// takeSnapshot takes snapshot of current processes
func (st *SnapshotTaker) takeSnapshot(i int) {
	// Take snapshot
	snapshot, err := collect.Processes()
	if err != nil {
		return
	}

	// Compare snapshot to baseline snapshot
	for _, item := range snapshot {
		if !contains(st.BaselineProcesses, item) && !contains(st.NewProcesses, item) {
			// If process is not in baseline snapshot, add to newProcesses map
			item.Type = types.NewProcess
			st.NewProcesses = append(st.NewProcesses, item)

			if !isEmpty(item) {
				st.reporter.Report(item)
			}
		}
	}
}

// contians checks whether given item is in baseline or not.
func contains(processes []*types.ProcessInfo, item *types.ProcessInfo) bool {
	for _, processItem := range processes {
		if processItem.PID == item.PID &&
			processItem.CreateTime == item.CreateTime &&
			processItem.Cmdline == item.Cmdline {
			return true
		}
	}
	return false
}

// isEmpty checks if given item is empty.
func isEmpty(item *types.ProcessInfo) bool {
	if item.Cmdline == "" &&
		item.PPID == -1 {
		return true
	}

	return false
}
