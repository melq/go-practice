package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INF = 1000000001
var m = 0
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
	fmt.Println(strings.TrimRight(fmt.Sprintf("%v", s)[1:], "]"))
}

func merge(s []int, l int, mid int, r int) {
	nl := mid - l; nr := r - mid
	a := make([]int, nl + 1); copy(a, s[l:mid])
	b := make([]int, nr + 1); copy(b, s[mid:r])
	a[nl] = INF; b[nr] = INF
	i, j := 0, 0
	for k := l; k < r; k++ {
		m++
		if a[i] <= b[j] {
			s[k] = a[i]
			i++
		} else {
			s[k] = b[j]
			j++
		}
	}
}

func mergeSort(s []int, l int, r int) {
	if l + 1 < r {
		mid := (l + r) / 2
		mergeSort(s, l, mid)
		mergeSort(s, mid, r)
		merge(s, l, mid, r)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var s []int
	for i := 0; i < n; i++ {
		s = append(s, nextInt())
	}
	mergeSort(s, 0, n)

	dump(s)
	fmt.Println(m)
}