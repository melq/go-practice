package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, err := strconv.Atoi(next())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	n := nextInt()

	var s []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		s = append(s, tmp)
	}

	fmt.Println(MaxDifferenceValue(s))
}

func MaxDifferenceValue(s []int) int {
	n := len(s)
	max := s[1] - s[0]
	min := s[0]

	for i := 1; i < n; i++ {
		if s[i] - min > max {
			max = s[i] - min
		}
		if s[i] < min {
			min = s[i]
		}
	}
	return max
}