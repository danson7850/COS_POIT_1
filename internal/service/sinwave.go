package service

import (
	"cos_poit_1/internal/utils"
	"math"
)

type Sin struct {
}

func NewSinWave() *Sin {
	return &Sin{}
}

func (s *Sin) CalcWave(ampl, freq, phase, duty float64, n int) (y int32) {

	y = int32(ampl * math.Pow(2, utils.Bits-1) * math.Sin(freq*(float64(n)/utils.Rate+1)*2*math.Pi+phase))

	return
}

func (s *Sin) CalcWaveP(ampl, phase, duty float64) (y int32) {

	y = int32(ampl * math.Pow(2, utils.Bits-1) * math.Sin(phase))

	return
}
