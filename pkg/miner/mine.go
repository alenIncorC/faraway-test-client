package miner

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"faraway-tcp-client/pkg/random"
	humanize "github.com/dustin/go-humanize"
)

func Mine(prefix []byte, difficulty []byte) []byte {
	start := time.Now()
	complexity, err := strconv.Atoi(string(difficulty))
	if err != nil {
		complexity = 2
	}
	numberOfGOR := runtime.NumCPU()
	hashes := make([]int, numberOfGOR)
	solutionChan := make(chan []byte)

	for i := 0; i < numberOfGOR; i++ {
		go func(index int, cmplx int) {
			offset := len(prefix)
			strWithPrefix := make([]byte, 20+offset)
			copy(strWithPrefix[:offset], prefix)
			seed := uint64(index)
			for {
				hashes[index]++
				seed = random.RandomString(strWithPrefix, offset, seed)
				if Hash(strWithPrefix, cmplx) {
					solutionChan <- strWithPrefix
					break
				}
			}
		}(i, complexity)
	}

	solution := <-solutionChan

	hashesSum := 0

	for i := 0; i < numberOfGOR; i++ {
		fmt.Printf("goroutine num: %d,number of hashes %d\n", i, hashes[i])
		hashesSum += hashes[i]
	}

	end := time.Now()
	fmt.Println(string(solution))
	fmt.Printf("totalNumber of hashes: %d\n", hashesSum)
	fmt.Printf("time spent: %g\n", end.Sub(start).Seconds())
	fmt.Println("processed:", humanize.Comma(int64(hashesSum)))
	fmt.Println("hashesh per sec:", humanize.Comma(int64(float64(hashesSum)/end.Sub(start).Seconds())))
	return solution
}

func Hash(str []byte, complexity int) bool {
	hash := sha256.Sum256(str)
	for i := 0; i < complexity; i++ {
		if hash[i] > 0 {
			return false
		}
	}

	return true
}
