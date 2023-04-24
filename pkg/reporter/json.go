package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mustafanafizdurukan/GoSnap/pkg/constants"
	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
)

// saveJson saves taken processes to json file
func saveJson(processes ...*types.ProcessInfo) {
	// Creating JSON files
	processesFile, err := os.Open(constants.FileJsonProcesses)
	if err != nil {
		os.Exit(1)
	}
	defer processesFile.Close()

	// Save new processes to JSON files
	allProcessesJSON, err := json.MarshalIndent(processes, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling all processes:", err)
		os.Exit(1)
	}
	processesFile.Write(allProcessesJSON)
}
