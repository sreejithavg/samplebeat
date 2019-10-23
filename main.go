package main

import (
	"github.com/sreejithavg/dummybeat/cmd"
	"os"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/sreejithavg/dummybeat/include"
)

func main() {
	
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
