package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	truncateCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "truncate":
		truncateCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'truncate'")
		fmt.Println("  tail:", truncateCmd.Args())
		// album.Truncate("localhost")
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
