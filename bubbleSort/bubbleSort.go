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

	fmt.Println(BubbleSort(s))
}

func BubbleSort(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if s[j] < s[j - 1] {
				tmp := s[j - 1]
				s[j - 1] = s[j]
				s[j] = tmp
			}
		}
	}
	return s
}