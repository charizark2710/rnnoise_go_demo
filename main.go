package main

import (
	"os"
	C_wrapper "rnnoise_go_demo/c_wrapper"
)

func main() {
	f, err := os.Create(os.Args[2:][0])
	if err != nil {
		panic(err)
	}
	f.Close()
	C_wrapper.ExecuteRnnoise(os.Args[1:][0], os.Args[2:][0])
}
