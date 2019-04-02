package ui

import (
	"bufio"
	"fmt"
	"strconv"
	"io"
	"errors"
	"github.com/pivotal-cf/guess/guess"
)

func Cli(input io.Reader, output io.Writer, rng guess.Random) {
	scanner := myScanner{bufio.NewScanner(input)}
	scanner.split()
	var writer = myWriter{bufio.NewWriter(output)}
	writer.writeMessage("Generating number between 1 and 10...\n")
	target := guess.GenerateTarget(rng)
	ourGuess := -1
	guessedCorrectly := false
	for !guessedCorrectly {

		writer.writeMessage("Guess the number!:\n")

		ourGuess, _ = scanner.grabGuess()

		writer.writeMessage(fmt.Sprintf("You guessed %d \n", ourGuess))

		guessedCorrectly = guess.IsGuessCorrect(ourGuess, target)

		if !guessedCorrectly {
			writer.writeMessage("You guessed wrong.\n")
		} else {
			writer.writeMessage("wooooo\n")
		}
	}
}

type myWriter struct{
	internalWriter *bufio.Writer
}

func (w myWriter) writeString(s string) (int, error) {
	return w.internalWriter.WriteString(s)
}

func (w myWriter) flush() error {
	return w.internalWriter.Flush()
}

func (w *myWriter) writeMessage(msg string) {
	_, _ = w.writeString(msg)
	_ = w.flush()
}

type myScanner struct {
	internalScanner *bufio.Scanner
}
func (s myScanner) split(){
	s.internalScanner.Split(bufio.ScanLines)
}

func (s myScanner) scan() bool{
	return s.internalScanner.Scan()
}

func (s myScanner) text() string{
	return s.internalScanner.Text()
}

func (s myScanner) grabGuess() (int, error) {
	if s.scan() {
		response := s.text()
		guess, _ := strconv.Atoi(response)

		return guess, nil
	}

	return -1, errors.New("error grabbing guessa")
}