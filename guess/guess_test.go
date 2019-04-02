package guess_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/pivotal-cf/guess/guess"
)

func TestGuess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Guess Suite")
}

var _ = Describe("When providing a guessa", func() {
	Context("and provided with a correct guessa", func(){
		It ("should return as successful", func() {
			result := guess.IsGuessCorrect(2, 2)
			Expect(result).To(BeTrue())
		})
	})
	Context("and provided with an incorrect guessa", func(){
		It ("should return as unsuccessful", func() {
			result := guess.IsGuessCorrect(3, 2)
			Expect(result).To(BeFalse())
		})
	})
})



