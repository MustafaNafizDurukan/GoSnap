package reporter

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

var (
	isTablePrinted bool
	writer         = tabwriter.NewWriter(os.Stdout, 0, 6, 2, ' ', 0)
)

func print(processes ...*types.ProcessInfo) {
	if !isTablePrinted {
		isTablePrinted = true
		fmt.Fprintf(writer, "Process Name\tPID\tPPID\t\n")
	}

	columnWidth := 34 // Minimum width for the first column

	for _, process := range processes {
		paddedName := padText(process.Name, columnWidth)
		fmt.Fprintf(writer, "%s\t%d\t%d\t\n", paddedName, process.PID, process.PPID)
	}

	writer.Flush()
}

func padText(text string, width int) string {
	if len(text) >= width {
		return text
	}

	padding := strings.Repeat(" ", width-len(text))
	return text + padding
}
