package service

import (
	"cos_poit_1/internal/utils"
	"math"
)

type Tri struct {
}

func NewTriWave() *Sin {
	return &Sin{}
}

func (s *Tri) CalcWave(ampl, freq, phase, duty float64, n int) (y int32) {

	y = int32(2 * ampl / math.Pi * math.Pow(2, utils.Bits-1) * math.Asin(math.Sin(freq*(float64(n)/utils.Rate+1)*2*math.Pi+phase)))

	return
}

func (s *Tri) CalcWaveP(ampl, phase, duty float64) (y int32) {

	y = int32(2 * ampl / math.Pi * math.Pow(2, utils.Bits-1) * math.Asin(math.Sin(phase)))

	return
}
