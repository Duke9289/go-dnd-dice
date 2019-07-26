package statrolls

import (
	"github.com/Duke9289/go-dnd-dice/diceroller"
)

func ThreeDSix() (stats []int) {
	for i := 0; i < 6; i++ {
		roll, _ := diceroller.ParseInputString("3d6")
		stats = append(stats, roll)
	}
	return
}

func FourDSixDropLowest() (stats []int) {
	for i := 0; i < 6; i++ {
		initialRoll, rolls := diceroller.ParseInputString("4d6")
		m := rolls[0]
		for i, roll := range rolls {
			if i == 0 || roll < m {
				m = roll
			}
		}
		stats = append(stats, initialRoll-m)
	}
	return
}
