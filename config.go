package main

import (
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

var (
	Version     string
	ShortCommit string
	BuildDate   string
	Debug       bool
)

func ParseArgsAndFlags() {
	flag.Usage = overrideUsage()

	flag.BoolVar(&Debug, "debug", false, "Show verbose debug information")
	showVersion := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *showVersion {
		fmt.Printf("git-my-git version %s, commit %s, build %s\n",
			Version, ShortCommit, BuildDate)
		os.Exit(0)
	}

	if Debug {
		debugLog()
	}
}

func overrideUsage() func() {
	return func() {
		_, _ = fmt.Fprintf(
			os.Stdout,
			"Usage:"+
				"\n\t"+
				"cd $(git-my-git [flags] [QUERY])"+
				"\n\n"+
				"Flags:"+
				"\n")
		flag.PrintDefaults()
	}
}

func debugLog() {
	fmt.Println("Args:")
	fmt.Printf("  args=%s\n", aurora.Cyan(flag.Args()))
	fmt.Println()
	fmt.Println("Config:")
}
