package reporter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

func print(processes ...types.ProcessInfo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, process := range processes {
		fmt.Fprintf(w, "%s\t%d\tPARENT=>\t%d\t\n", process.Name, process.PID, process.PPID)
	}

	w.Flush()
}
