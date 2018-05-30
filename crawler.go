package main

import (
	"fmt"
	"live"
	"os"
	"time"
)

func main() {
	for {
		hf, content, err := live.HasFootball()
		if err != nil {
			fmt.Fprintf(os.Stderr, "HasFootball error: %v\n", err)
			os.Exit(1)
		}
		if hf {
			fmt.Println("has football now")
			break
		} else {
			fmt.Println("no football now")
		}

		time.Sleep(30 * time.Second)
	}
}
