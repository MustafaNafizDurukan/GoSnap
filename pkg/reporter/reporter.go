package reporter

import "github.com/mustafanafizdurukan/GoSnap/pkg/types"

type Reporter struct {
	reporters []typeReportFunc
}

type typeReportFunc func(...*types.ProcessInfo)

func New(reporterFuncs ...typeReportFunc) *Reporter {
	r := Reporter{
		reporters: reporterFuncs,
	}

	r.reporters = append(r.reporters, print)

	return &r
}

func (r *Reporter) Report(processes ...*types.ProcessInfo) {
	for _, r := range r.reporters {
		r(processes...)
	}
}
