package day07

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	var sum int64 = 0
	var testVal int64

	for l := range input.Lines(f) {
		var operands []int64
		tokens := strings.Split(l, ":")
		testVal, _ = strconv.ParseInt(tokens[0], 10, 64)
		tokens = strings.Split(strings.Trim(tokens[1], " "), " ")
		for _, t := range tokens {
			n, _ := strconv.ParseInt(t, 10, 64)
			operands = append(operands, n)
		}

		if isPossible(testVal, operands) {
			sum += testVal
		}
	}

	return sum
}

func Part2(f *os.File) int64 {
	return 0
}

func isPossible(target int64, operands []int64) bool {
	numOperators := len(operands) - 1
	combos := getOpCombos(numOperators)
	jobChan := make(chan evalJob, len(combos))
	resChan := make(chan int64, len(combos))
	defer close(jobChan)
	parallelism := runtime.NumCPU()

	for i := 0; i < parallelism; i++ {
		go worker(jobChan, resChan)
	}

	for _, c := range combos {
		fmt.Printf("spawning worker for %+v with ops %+v\n", operands, c)
		jobChan <- evalJob{c, operands}
	}

	for i := 0; i < len(combos); i++ {
		if <-resChan == target {
			return true
		}
	}
	return false
}

func getOpCombos(n int) []string {
	ops := []string{"+", "*"}
	numCombos := int(math.Pow(float64(len(ops)), float64(n)))
	combos := make([]string, numCombos)
	fmtStr := "%0" + strconv.Itoa(n) + "b"

	var bitSet int64

	for {
		if bitSet == int64(numCombos) {
			break
		}

		currCombo := ""
		for _, b := range fmt.Sprintf(fmtStr, bitSet) {
			switch b {
			case '0':
				currCombo += "+"
			case '1':
				currCombo += "*"
			}
		}
		combos[bitSet] = currCombo
		bitSet++
	}

	return combos
}

type evalJob struct {
	operations string
	operands   []int64
}

func worker(jobs <-chan evalJob, results chan<- int64) {
	for j := range jobs {
		results <- eval(j)
	}
}

func eval(job evalJob) int64 {
	fmt.Printf("evaluating %+v\n", job)
	var res int64 = job.operands[0]
	for i := 0; i < len(job.operands)-1; i++ {
		switch job.operations[i] {
		case '+':
			res += job.operands[i+1]
		case '*':
			res *= job.operands[i+1]
		}
	}

	return res
}
