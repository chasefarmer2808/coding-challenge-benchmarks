package day07

import (
	"coding-challenge-runner/pkg/input"
	"context"
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

		if isPossible(testVal, operands, 2) {
			sum += testVal
		}
	}

	return sum
}

func Part2(f *os.File) int64 {
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

		if isPossible(testVal, operands, 3) {
			sum += testVal
		}
	}

	return sum
}

func isPossible(target int64, operands []int64, base int) bool {
	numOperators := len(operands) - 1
	numCombos := int(math.Pow(float64(base), float64(numOperators)))
	jobChan := make(chan evalJob, numCombos)
	resChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer close(jobChan)
	parallelism := runtime.NumCPU()

	for i := 0; i < parallelism; i++ {
		go func() {
			worker(ctx, jobChan, resChan, target)
		}()
	}

	for i := 0; i < numCombos; i++ {
		operations := intToOperations(int64(i), numOperators, base)
		jobChan <- evalJob{operations, operands}
	}

	for i := 0; i < numCombos; i++ {
		if <-resChan {
			cancel()
			return true
		}
	}
	return false
}

func intToOperations(n int64, size, base int) string {
	binStr := strconv.FormatInt(n, base)
	return strings.Repeat("0", size-len(binStr)) + binStr
}

type evalJob struct {
	operations string
	operands   []int64
}

func worker(ctx context.Context, jobs <-chan evalJob, results chan<- bool, target int64) {
	for j := range jobs {
		select {
		case <-ctx.Done():
			return
		default:
			results <- eval(j) == target
		}
	}
}

func eval(job evalJob) int64 {
	var res int64 = job.operands[0]
	for i := 0; i < len(job.operands)-1; i++ {
		switch job.operations[i] {
		case '0': // add
			res += job.operands[i+1]
		case '1': // mul
			res *= job.operands[i+1]
		case '2': // concat
			rs := strconv.FormatInt(res, 10)
			ns := strconv.FormatInt(job.operands[i+1], 10)
			res, _ = strconv.ParseInt(rs+ns, 10, 64)
		}
	}

	return res
}
