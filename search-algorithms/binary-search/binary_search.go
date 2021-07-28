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

func BinarySearch(s []int, key int) bool {
	left := 0; right := len(s)

	for left < right {
		mid := (left + right) / 2
		if s[mid] == key {
			return true
		} else if key < s[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
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
		if BinarySearch(s, t[i]) {
			cnt++
		}
	}

	fmt.Println(cnt)
}