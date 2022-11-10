package service

import (
	"cos_poit_1/internal/utils"
	"math"
)

type Pillow struct {
}

func NewPilWave() *Pillow {
	return &Pillow{}
}

func (s *Pillow) CalcWave(ampl, freq, phase, duty float64, n int) (y int32) {

	y = int32(-2 * ampl / math.Pi * math.Pow(2, utils.Bits-1) *
		math.Atan(math.Cos(freq*(float64(n)/utils.Rate+1)*math.Pi+phase)/
			math.Sin(freq*(float64(n)/utils.Rate+1)*math.Pi+phase)))

	return
}

func (s *Pillow) CalcWaveP(ampl, phase, duty float64) (y int32) {

	y = int32(-2 * ampl / math.Pi * math.Pow(2, utils.Bits-1) *
		math.Atan(math.Cos(phase)/math.Sin(phase)))
	return
}
