package guess

type Random func(n int) int

func IsGuessCorrect(guess int, target int) bool {
	return guess == target
}

func GenerateTarget(rng Random) int {
	target := rng(10)
	target += 1
	return target
}

