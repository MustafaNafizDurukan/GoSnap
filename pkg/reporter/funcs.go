package reporter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

var (
	isTablePrinted bool
	writer         = tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
)

func print(processes ...*types.ProcessInfo) {

	if !isTablePrinted {
		isTablePrinted = true
		fmt.Fprintf(writer, "Process Name\tPID\tPPID\t\n")
	}

	for _, process := range processes {
		fmt.Fprintf(writer, "%s\t%d\t%d\t\n", process.Name, process.PID, process.PPID)
	}

	writer.Flush()
}

func wrapText(text string, width int) string {
	if len(text) <= width {
		return text
	}

	var wrappedText string
	for len(text) > width {
		wrappedText += text[:width] + "\n"
		text = text[width:]
	}

	return wrappedText + text
}
