package main

import (
	"coding-challenge-runner/aoc/2023/day01"
	"fmt"
	"time"
)

type Runnable interface {
	Run() error
	Name() string
}

func main() {
	aoc2023 := []Runnable{
		&day01.Aoc2023Day01{},
	}

	for _, p := range aoc2023 {
		start := time.Now()
		fmt.Printf("Running problem %s...", p.Name())
		p.Run()
		fmt.Println(time.Since(start).String())
	}
}
