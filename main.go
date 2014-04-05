package main

import (
	"./rockstar"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		printUsage()
	} else {
		username := flag.Arg(0)
		rockstar.ShowSummarization(username)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  rockstar [username]")
}
