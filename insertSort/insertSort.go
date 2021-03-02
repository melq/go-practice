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
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var a []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		a = append(a, tmp)
	}

	fmt.Println(InsertSort(a))
}

func InsertSort (a []int) []int {
	n := len(a)

	for i := 0; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j + 1] = a[j]
			} else {
				break
			}
		}
		a[j + 1] = v
		//fmt.Println(a)
	}
	return a
}