package service

import (
	"log"
)

type Waves interface {
	CalcWave(ampl, freq, phase, duty float64, n int) (y int32)
	CalcWaveP(ampl, phase, duty float64) (y int32)
}

func NewWave(signal string) (waveF Waves) {
	switch signal {
	case "sin":
		waveF = NewSinWave()
		log.Println(signal)
	case "triangle":
		waveF = NewTriWave()
		log.Println(signal)
	case "noise":
		waveF = NewNoiseWave()
		log.Println(signal)
	case "pillow":
		waveF = NewPilWave()
		log.Println(signal)
	case "pulse":
		waveF = NewPulseWave()
		log.Println(signal)
	default:
		waveF = NewSinWave()
		log.Println(signal)
	}
	return
}
