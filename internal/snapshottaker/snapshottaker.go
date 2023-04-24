package snapshottaker

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
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

	printSeperator("Baseline Processes")
	st.reporter.Report(st.BaselineProcesses...)

	printSeperator("New Processes")
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
			st.NewProcesses = append(st.NewProcesses, item)

			st.reporter.Report(item)
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

func printSeperator(seperator string) {
	padding := 17

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", strings.Repeat("*", padding), seperator, strings.Repeat("*", padding))
	tw.Flush()
}
