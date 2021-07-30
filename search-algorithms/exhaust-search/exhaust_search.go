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

func solve(s []int, i int, key int) bool {
	if key == 0 {
		return true
	}
	if i >= len(s) {
		return false
	}
	res := solve(s, i + 1, key) || solve(s, i + 1, key - s[i])
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	var s []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		s = append(s, tmp)
	}
	q := nextInt()
	for i := 0; i < q; i++ {
		if solve(s, 0, nextInt()) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
