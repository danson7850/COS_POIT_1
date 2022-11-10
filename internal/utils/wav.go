package utils

import (
	"fmt"
	"github.com/cryptix/wav"
	"os"
	"time"
)

func WriteWav(wave []int32) {
	wavOut, err := os.Create("Out.wav")
	checkErr(err)
	defer wavOut.Close()

	meta := wav.File{
		Channels:        1,
		SampleRate:      Rate,
		SignificantBits: Bits,
	}

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	start := time.Now()

	for _, signal := range wave {
		err = writer.WriteInt32(signal)
		checkErr(err)
	}

	fmt.Printf("Simulation Done. Took:%v\n", time.Since(start))
}
