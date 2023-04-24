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
		fmt.Fprintf(writer, "Type\tProcess Name\tPID\tPPID\t\n")
	}

	nameWidth := 34 // Minimum width for the process name
	typeWidth := 8  // Minimum width for the type

	for _, process := range processes {
		process.Name = padText(process.Name, nameWidth)
		processType := string(process.Type)

		processType = padText(processType, typeWidth)
		fmt.Fprintf(writer, "%s\t%s\t%d\t%d\t\n", processType, process.Name, process.PID, process.PPID)
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
