package service

import (
	"cos_poit_1/internal/utils"
	"math"
)

var (
	Ak = []float64{0.7, 0.7, 0.7, 0.7, 0.7}
	Fk = []float64{441, 442, 443, 444, 445}
	Pk = []float64{math.Pi, math.Pi / 4, 0, 3 * math.Pi / 4, math.Pi / 2}
)

func PolyharmonicSignal(duty float64, nameOne, nameTwo string, sec int) (xn []int32) {

	waveOne := NewWave(nameOne)
	waveTwo := NewWave(nameTwo)

	var sum int32

	for n := 0; n < utils.Rate*sec; n++ {
		sum = 0
		for j := 0; j < 5; j++ {
			if j > 3 {
				sum += waveOne.CalcWave(Ak[j], Fk[j], Pk[j], duty, n)
			} else {
				sum += waveTwo.CalcWave(Ak[j], Fk[j], Pk[j], duty, n)
			}
		}
		xn = append(xn, sum)
	}
	return
}
