package service

import (
	"cos_poit_1/internal/utils"
	"log"
	"math"
)

func AmplModulation(nameOne, nameTwo string, sec int) (result []int32) {
	waveOne := NewWave(nameOne)
	waveTwo := NewWave(nameTwo)

	var xnOne []int32
	var xnTwo []int32

	for n := 0; n < utils.Rate*sec; n++ {
		xnOne = append(xnOne, waveOne.CalcWave(0.8, 440, 1, 0.5, n))
		xnTwo = append(xnTwo, waveTwo.CalcWave(0.2, 10, 1, 0.5, n))
	}

	var i float64
	for j, _ := range xnOne {
		i = float64(xnTwo[j])/math.Pow(2, utils.Bits-1) + 1.5
		result = append(result, int32(float64(xnOne[j])*i))
	}

	log.Println(result[0:40])

	return
}

func FreqModulation(nameOne, nameTwo string, sec int) (result []int32) {
	waveOne := NewWave(nameOne)
	waveTwo := NewWave(nameTwo)

	var xnOne []int32
	var xnTwo []int32

	for n := 0; n < utils.Rate*sec; n++ {
		xnOne = append(xnOne, waveOne.CalcWave(0.8, 440, 1, 0.5, n))
		xnTwo = append(xnTwo, waveTwo.CalcWave(0.2, 440, 1, 0.5, n))
	}

	i := 0.0

	for j, _ := range xnOne {
		result = append(result, waveOne.CalcWaveP(0.8, i, 0.5))
		i += 2 * math.Pi * 440 * (1 + float64(xnTwo[j])/math.Pow(2, utils.Bits-1))
	}

	log.Println(result[0:40])

	return
}
