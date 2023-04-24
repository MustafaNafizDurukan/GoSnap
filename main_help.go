package main

import (
	"fmt"
	"os"

	"github.com/mustafanafizdurukan/GoSnap/pkg/constants"
)

func welcome() {
	fmt.Println(`
	 ██████╗  ██████╗ ███████╗███╗   ██╗ █████╗ ██████╗ 
	██╔════╝ ██╔═══██╗██╔════╝████╗  ██║██╔══██╗██╔══██╗
	██║  ███╗██║   ██║███████╗██╔██╗ ██║███████║██████╔╝
	██║   ██║██║   ██║╚════██║██║╚██╗██║██╔══██║██╔═══╝ 
	╚██████╔╝╚██████╔╝███████║██║ ╚████║██║  ██║██║     
	 ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     
	         Take Process Snaphshots Easily!
 `)
}

func initialization() {
	// Creating JSON files
	_, err := os.Create(constants.FileJsonProcesses)
	if err != nil {
		fmt.Printf("could not create file %s: %v", constants.FileJsonProcesses, err)
		os.Exit(1)
	}
}
