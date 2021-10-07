package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	if !wmctrlIsOnPath() {
		fmt.Printf("you must have wmctrl on path to run this program\n")
		os.Exit(1)
	}

	var executableName string
	var processWindowName string

	flag.StringVar(&executableName, "name", "", "the executable name")
	flag.StringVar(&processWindowName, "windowName", "", "the process window name (use xprop to get the name)")
	flag.Parse()

	if executableName == "" {
		fmt.Printf("you must provide the process executable name (name)\n")
		os.Exit(1)
	}

	if processWindowName == "" {
		fmt.Printf("you must provide the process window name (windowName)\n")
		os.Exit(1)
	}

	processList, err := ps.Processes()
	if err != nil {
		fmt.Printf("could not enumerate the process, %v\n", err)
		os.Exit(1)
	}

	for _, process := range processList {
		if process.Executable() == executableName {
			err := exec.Command("wmctrl", "-a", processWindowName).Run()
			if err != nil {
				fmt.Printf("error switching to the process, %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}
	}

	err = exec.Command(executableName).Start()
	if err != nil {
		fmt.Printf("error starting the process, %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

// wmctrlIsOnPath check if wmctrl is installed
func wmctrlIsOnPath() bool {
	_, err := exec.LookPath("wmctrl")
	return err == nil
}
