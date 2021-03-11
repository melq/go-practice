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

	ShellSort(s)
}

func ShellSort(s []int) {
	n := len(s)
	count := 0
	G := []int{1}
	for i := 4; i < n; i = i * 3 + 1 {
		G = append(G, i)
	}
	m := len(G)

	for i := range G {
		if i < m / 2 {
			G[i], G[m - i - 1] = G[m - i - 1], G[i]
		}
		InsertSort(s, G[i], &count)
	}
	fmt.Println(m)
	dump(G)
	fmt.Println(count)
	for _, v := range s {
		fmt.Println(v)
	}
}

func InsertSort (s []int, g int, count *int) {
	n := len(s)

	for i := g; i < n; i++ {
		v := s[i]
		j := i - g
		for j >= 0 && s[j] > v {
			s[j + g] = s[j]
			*count++
			j -= g
		}
		s[j + g] = v
	}
}