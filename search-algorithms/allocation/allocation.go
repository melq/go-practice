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

func isOK(k int, size int, s []int) bool {
	trucks := 0
	sum := 0
	for _, v := range s {
		if sum + v > size {
			trucks++
			if trucks >= k {
				return false
			}
			sum = 0
		}
		sum += v
	}
	return true
}

func minTruckSize(k int, s []int) int {
	l := 0; r := 0
	for _, v := range s {
		if l < v {
			l = v
		}
		r += v
	}
	for l < r {
		mid := (l + r) / 2
		if isOK(k, mid, s) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()

	var s []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		s = append(s, tmp)
	}

	fmt.Println(minTruckSize(k, s))
}