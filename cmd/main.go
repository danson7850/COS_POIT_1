package main

import (
	"cos_poit_1/internal/service"
	"cos_poit_1/internal/utils"
	"github.com/wcharczuk/go-chart"
	"log"
	"os"
)

func main() {

	//noise, triangle, pillow, sin, pulse

	//sec := 2
	//var xn []int32
	//
	//wave := service.NewWave("pulse")
	//
	//for n := 0; n < utils.Rate*sec; n++ {
	//	xn = append(xn, wave.CalcWave(0.8, 440, 1, 0.0, n))
	//}
	//
	//xn := service.PolyharmonicSignal(0.5, "triangle", "pillow", 1)

	//xn := service.AmplModulation("sin", "pulse", 2)
	xn := service.FreqModulation("sin", "triangle", 2)

	var rn []float64
	for n := 0; n < 44100; n++ {
		rn = append(rn, float64(n))
	}

	var yn []float64
	for _, j := range xn {
		yn = append(yn, float64(j))
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: rn[20000:24000],
				YValues: yn[20000:24000],
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	err := graph.Render(chart.PNG, f)
	log.Println("Chart error: ", err)

	utils.WriteWav(xn)
}
