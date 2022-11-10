package service

import (
	"cos_poit_1/internal/utils"
	"math"
)

type Pulse struct {
}

func NewPulseWave() *Pulse {
	return &Pulse{}
}

// Частота дескритизации, период и рисовать синусоиду 3 гц (example)
func (s *Pulse) CalcWave(ampl, freq, phase, duty float64, n int) (y int32) {
	res := math.Mod(math.Pi*freq*(float64(n)/utils.Rate)+phase, math.Pi) / math.Pi
	if res < duty {
		y = int32(ampl * math.Pow(2, utils.Bits-1))
	} else {
		y = int32(-ampl * math.Pow(2, utils.Bits-1))
	}
	return
}

func (s *Pulse) CalcWaveP(ampl, phase, duty float64) (y int32) {

	res := math.Mod(phase, math.Pi) / math.Pi
	if res < duty {
		y = int32(ampl * math.Pow(2, utils.Bits-1))
	} else {
		y = int32(-ampl * math.Pow(2, utils.Bits-1))
	}
	return
}
