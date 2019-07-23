package main

import (
	"bufio"
	"math/rand"
	"os"

	"github.com/pivotal-cf/guess/ui"
	"golang.org/x/sys/unix"
)

func main() {
	_, err := unix.IoctlGetTermios(int(os.Stdout.Fd()), unix.TCGETS)

	if err == nil {
		ui.Cli(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), rand.Intn)
	} else {
		ui.Webserver()
	}

}
