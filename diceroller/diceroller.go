package diceroller

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ParseInputString(input string) (int, []int) {

	// TODO: validate format of input string
	// TODO: add support for operators

	diceNumbers := strings.Split(input, "d")
	noOfDie, _ := strconv.Atoi(diceNumbers[0])
	sizeOfDie, _ := strconv.Atoi(diceNumbers[1])
	return rollDie(sizeOfDie, noOfDie)
}

func rollDie(size int, noOfDie int) (int, []int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rolls []int
	var sum int
	for i := 0; i < noOfDie; i++ {
		roll := r1.Intn(size) + 1
		rolls = append(rolls, roll)
		sum = sum + roll
	}
	return sum, rolls
}
