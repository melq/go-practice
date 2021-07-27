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

func LinearSearch(s []int, key int) bool {
	i := 0; n := len(s)
	s = append(s, key)
	for s[i] != key {
		i++
		if i == n {
			return false
		}
	}
	return true
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
	var t []int
	for i := 0; i < q; i++ {
		tmp := nextInt()
		t = append(t, tmp)
	}

	cnt := 0
	for i := 0; i < len(t); i++ {
		if LinearSearch(s, t[i]) {
			cnt++
		}
	}

	fmt.Println(cnt)
}
