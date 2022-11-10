package service

import (
	"cos_poit_1/internal/utils"
	"math"
	"math/rand"
)

type Noise struct {
}

func NewNoiseWave() *Noise {
	return &Noise{}
}

func (s *Noise) CalcWave(ampl, freq, phase, duty float64, n int) (y int32) {

	y = int32(ampl * math.Pow(2, utils.Bits-1) * math.Sin(freq*(float64(n)/utils.Rate+1)*2*math.Pi*float64(rand.Intn(20)-10)+phase))

	return
}

func (s *Noise) CalcWaveP(ampl, phase, duty float64) (y int32) {

	y = int32(ampl * math.Pow(2, utils.Bits-1) * math.Sin(phase))

	return
}
