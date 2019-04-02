package main

import (
	"bufio"
	"os"
	"math/rand"
	"github.com/pivotal-cf/guess/ui"
)

func main () {
	if len(os.Args) > 1 && os.Args[1] == "-http" {
		ui.Webserver()
	} else {
		ui.Cli(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), rand.Intn)
	}

}