// Command darban-webpanel will serve the web panel of Darban OVPN. For now it only reports its version.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

func run(args []string, out io.Writer) int {
	flags := flag.NewFlagSet("darban-webpanel", flag.ContinueOnError)
	flags.SetOutput(out)
	showVersion := flags.Bool("version", false, "print the version and exit")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if *showVersion {
		fmt.Fprintln(out, version)
		return 0
	}
	fmt.Fprintln(out, "darban-webpanel is not ready yet; only -version works.")
	return 1
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}
