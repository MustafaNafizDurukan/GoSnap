package reporter

import "github.com/mustafanafizdurukan/GoSnap/pkg/types"

type Reporter struct {
	reporters []func(...types.ProcessInfo)
}

func NewReporter() *Reporter {
	return &Reporter{
		reporters: []func(...types.ProcessInfo){
			print,
		},
	}
}

func (r *Reporter) Report(processes ...types.ProcessInfo) {
	for _, r := range r.reporters {
		r(processes...)
	}
}
