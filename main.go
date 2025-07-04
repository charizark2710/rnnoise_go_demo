package main

import (
	"fmt"
	"os"
	audio "rnnoise_go_demo/audio"
	rnnoise "rnnoise_go_demo/rnnoise"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	input := os.Args[1:][0]
	outputRaw := strings.Split(input, ".")[0] + "_output.raw"
	f, err := os.Create(outputRaw)
	if err != nil {
		panic(err)
	}
	f.Close()
	input_48khz, err := audio.ConvertToRaw48kHz(input)
	finishTime_1 := time.Now()
	elapsedTime := finishTime_1.Sub(startTime)
	fmt.Printf("Finished ConvertToRaw48kHz in %v\n", elapsedTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	rnnoise.ExecuteRnnoise(input_48khz, outputRaw)
	finishTime_2 := time.Now()
	elapsedTime = finishTime_2.Sub(finishTime_1)
	fmt.Printf("Finished ExecuteRnnoise in %v\n", elapsedTime)

	audio.ConvertToWav(outputRaw)
	finishTime_3 := time.Now()
	elapsedTime = finishTime_3.Sub(finishTime_2)
	fmt.Printf("Finished ConvertToWav in %v\n", elapsedTime)

	finishTime := time.Now()
	elapsedTime = finishTime.Sub(startTime)
	fmt.Printf("Finished in %v\n", elapsedTime)
}
