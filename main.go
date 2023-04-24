package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mustafanafizdurukan/GoSnap/internal/snapshottaker"
)

func main() {
	var duration string
	var err error

	flag.StringVar(&duration, "d", "", "Duration to run the tool, e.g. 10s or 10m")
	flag.Parse()

	var givenTime time.Duration
	if duration != "" {
		givenTime, err = time.ParseDuration(duration)
		if err != nil {
			fmt.Println("Error parsing duration:", err)
			os.Exit(1)
		}
	}

	welcome()

	initialization()

	st := snapshottaker.New(givenTime)
	st.Start()
}
