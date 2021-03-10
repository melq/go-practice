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

func dump(s []string) {
	for i, v := range s {
		if i == 0 {
			fmt.Print(v)
		} else {
			fmt.Print(" ")
			fmt.Print(v)
		}
	}
	fmt.Print("\n")
}

func swapInt(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func swapString(s []string, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func equals(s1, s2 []string) bool {
	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var s []string
	for i := 0; i < n; i++ {
		tmp := next()
		s = append(s, tmp)
	}

	bubbleS := append([]string{}, s...)
	selectionS := append([]string{}, s...)

	BubbleSort(bubbleS)
	dump(bubbleS)
	fmt.Println("Stable")

	SelectionSort(selectionS)
	dump(selectionS)
	if equals(bubbleS, selectionS) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func BubbleSort(s []string) {
	n := len(s)
	var sInt []int
	for i := 0; i < n; i++ {
		tmp, _ := strconv.Atoi(s[i][1:])
		sInt = append(sInt, tmp)
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if sInt[j] < sInt[j - 1] {
				swapInt(sInt, j - 1, j)
				swapString(s, j - 1, j)
			}
		}
	}
}

func SelectionSort(s []string) {
	n := len(s)
	var sInt []int
	for i := 0; i < n; i++ {
		tmp, _ := strconv.Atoi(s[i][1:])
		sInt = append(sInt, tmp)
	}

	for i := 0; i < n; i++ {
		minIndex := i
		j := i
		for ; j < n; j++ {
			if sInt[j] < sInt[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			swapInt(sInt, i, minIndex)
			swapString(s, i, minIndex)
		}
	}
}