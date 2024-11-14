package main

import (
	"flag"

	"github.com/subliker/ht-conf_conf-lang/internal/shell"
)

func main() {
	// parsing flags in all packages
	flag.Parse()

	// creating shell
	shell := shell.New()

	// run shell
	shell.Run()
}
