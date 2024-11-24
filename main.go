package main

import (
	aoc2023 "coding-challenge-runner/pkg/aoc/2023"
	"flag"
	"fmt"
	"time"
)

type Runnable interface {
	Run() error
	Name() string
}

func main() {
	name := flag.String("name", "", "Name of problem to run; runs all problems if omitted")

	challenges := []Runnable{
		&aoc2023.Day01{},
	}

	for _, p := range challenges {
		if *name != "" && *name != p.Name() {
			continue
		}
		start := time.Now()
		fmt.Printf("Running problem %s...", p.Name())
		p.Run()
		fmt.Println(time.Since(start).String())
	}
}
