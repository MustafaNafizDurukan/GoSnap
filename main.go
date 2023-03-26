package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var duration string
	flag.StringVar(&duration, "d", "10s", "Duration to run the tool, e.g. 10s or 10m")
	flag.Parse()

	givenTime, err := time.ParseDuration(duration)
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		os.Exit(1)
	}

	takeSnapshots(int(givenTime.Seconds()))

	saveJson()
}
