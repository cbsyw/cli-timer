package main

import (
	"flag"
	"fmt"
	"time"
	"os"
)

func main() {
	duration := flag.String("duration", "1", "set timer duration (e.g. 1s, 1m, 1h)")
	flag.Parse()

	d, err := time.ParseDuration(*duration)
	if err != nil {
		fmt.Println("invalid duration:", err)
		os.Exit(1)
	}

	fmt.Printf("timer set for %s. Counting down...\n", d)
	time.Sleep(d)

	fmt.Println("Time's up!")

}
