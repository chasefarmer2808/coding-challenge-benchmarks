package day07

import (
	"coding-challenge-runner/pkg/input"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	return solve(f, 2)
}

func Part2(f *os.File) int64 {
	return solve(f, 3)
}

func solve(f *os.File, base int) int64 {
	var sum int64 = 0
	var testVals []int64
	var operands [][]int64
	parallelism := runtime.NumCPU()

	for l := range input.Lines(f) {
		var ops []int64
		tokens := strings.Split(l, ":")
		testVal, _ := strconv.ParseInt(tokens[0], 10, 64)
		tokens = strings.Split(strings.Trim(tokens[1], " "), " ")
		for _, t := range tokens {
			n, _ := strconv.ParseInt(t, 10, 64)
			ops = append(ops, n)
		}
		testVals = append(testVals, testVal)
		operands = append(operands, ops)
	}
	jobChan := make(chan evalJob, len(testVals))
	defer close(jobChan)
	resChan := make(chan result, len(testVals))
	// fmt.Printf("parsed vals %+v and ops %+v\n", testVals, operands)
	for i := 0; i < parallelism; i++ {
		go worker(jobChan, resChan)
	}

	for i, testVal := range testVals {
		// fmt.Println("sending job")
		jobChan <- evalJob{testVal, operands[i], base}
	}

	for range testVals {
		// fmt.Println("waiting for res")
		res := <-resChan
		// fmt.Printf("got res %+v\n", res)
		if res.isPossible {
			sum += res.target
		}
	}
	close(resChan)

	return sum
}

type evalJob struct {
	target   int64
	operands []int64
	base     int
}

type result struct {
	isPossible bool
	target     int64
}

func worker(jobs <-chan evalJob, results chan<- result) {
	for j := range jobs {
		results <- result{isPossible(j.target, j.operands, j.base), j.target}
	}
}

func isPossible(target int64, operands []int64, base int) bool {
	// fmt.Printf("checking %d with ops %+v\n", target, operands)
	numOperators := len(operands) - 1
	numCombos := int(math.Pow(float64(base), float64(numOperators)))

	for i := 0; i < numCombos; i++ {
		operations := intToOperations(int64(i), numOperators, base)
		if eval(operands, operations) == target {
			return true
		}
	}

	return false
}

func intToOperations(n int64, size, base int) string {
	binStr := strconv.FormatInt(n, base)
	return strings.Repeat("0", size-len(binStr)) + binStr
}

func eval(operands []int64, operations string) int64 {
	// fmt.Printf("evaluating %+v with operations %+v\n", operands, operations)
	var res int64 = operands[0]
	for i := 0; i < len(operands)-1; i++ {
		switch operations[i] {
		case '0': // add
			res += operands[i+1]
		case '1': // mul
			res *= operands[i+1]
		case '2': // concat
			rs := strconv.FormatInt(res, 10)
			ns := strconv.FormatInt(operands[i+1], 10)
			res, _ = strconv.ParseInt(rs+ns, 10, 64)
		}
	}

	return res
}
