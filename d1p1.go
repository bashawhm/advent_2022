package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func max(list []int) int {
	maxIndex := 0
	for i := 0; i < len(list); i++ {
		if list[i] > list[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func main() {
	buckets := make([]int, 1)
	scanner := bufio.NewScanner(os.Stdin)

	currBucket := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currBucket++
			buckets = append(buckets, 0)
		} else {
			num, _ := strconv.Atoi(line)
			buckets[currBucket] += num
		}
	}

	fmt.Println(buckets[max(buckets)])
}