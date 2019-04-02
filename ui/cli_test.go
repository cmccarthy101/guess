package ui_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"strings"
	"testing"
)

func TestGuess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Guess Suite")
}

var allNumbers = "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n"

var _ = Describe("When using the command line interface", func(){
	Context("and provided with a correct guessa", func(){
		It("exits successfully", func(){
			input := strings.NewReader("2\n")
			outputBuffer := gbytes.NewBuffer()

			ui.Cli(input, outputBuffer, func (_ int) int {return 1})
			Expect(outputBuffer).Should(gbytes.Say("woooo"))
			outputBuffer.Close()
		})
	})

	Context("and provided with all numbers", func(){
		It("exits successfully", func(){
			input := strings.NewReader(allNumbers)
			outputBuffer := gbytes.NewBuffer()

			ui.Cli(input, outputBuffer, func (_ int) int {return 9})
			Expect(outputBuffer).Should(gbytes.Say("woooo"))
			outputBuffer.Close()
		})
	})
})

