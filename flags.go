package main

import (
	"strings"
)

type Flags struct {
	Output string
	Fs string
}

func parseArgs(args []string) (Flags, []string) {
	var remaining []string
	var flags Flags

	for i := 0; i < len(args); i++ {

		if strings.HasPrefix(args[i], "--") {

			sliceFlag := strings.SplitN(args[i], "=", 2)
			flagName := sliceFlag[0]
			flagVal := sliceFlag[1]

			switch flagName {
			case "--output":
				flags.Output = flagVal
			case "--fs":
				flags.Fs = flagVal
			}

		} else {
			remaining = append(remaining, args[i])
		}
	}
	return flags, remaining
}