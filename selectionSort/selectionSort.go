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

func dump(s []int) {
	for i, v := range s {
		if i == 0 {
			fmt.Print(strconv.Itoa(v))
		} else {
			fmt.Print(" ")
			fmt.Print(strconv.Itoa(v))
		}
	}
	fmt.Print("\n")
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var s []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		s = append(s, tmp)
	}

	dump(SelectionSort(s))
}

func SelectionSort(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		minIndex := i
		j := i
		for ; j < n; j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			tmp := s[i]
			s[i] = s[minIndex]
			s[minIndex] = tmp
		}
	}

	return s
}