package reporter

import (
	"sync"

	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

type Reporter struct {
	reporters     []typeReportFunc
	reporterMutex sync.Mutex
}

type typeReportFunc func(...*types.ProcessInfo)

func New(reporterFuncs ...typeReportFunc) *Reporter {
	r := Reporter{
		reporters: reporterFuncs,
	}

	r.reporters = append(r.reporters, print)
	r.reporters = append(r.reporters, saveDB)

	return &r
}

func (r *Reporter) Report(processes ...*types.ProcessInfo) {
	for _, report := range r.reporters {
		r.reporterMutex.Lock()
		report(processes...)
		r.reporterMutex.Unlock()
	}
}
