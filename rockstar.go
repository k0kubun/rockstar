package main

import (
	"flag"
	"fmt"
	"github.com/k0kubun/rockstar/summarizer"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		printUsage()
	} else {
		username := flag.Arg(0)
		summarizer.ShowSummarization(username)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  rockstar [username]")
}
