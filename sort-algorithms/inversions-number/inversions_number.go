package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INF = 1000000001
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

func merge(s []int, l int, mid int, r int) int {
	nl := mid - l; nr := r - mid; cnt := 0
	a := make([]int, nl + 1); copy(a, s[l:mid])
	b := make([]int, nr + 1); copy(b, s[mid:r])
	a[nl] = INF; b[nr] = INF
	i, j := 0, 0
	for k := l; k < r; k++ {
		if a[i] <= b[j] {
			s[k] = a[i]
			i++
		} else {
			s[k] = b[j]
			j++
			cnt += nl - i
		}
	}
	return cnt
}

func MergeSort(s []int, l int, r int) int {
	cnt := 0
	if l + 1 < r {
		mid := (l + r) / 2
		cnt += MergeSort(s, l, mid)
		cnt += MergeSort(s, mid, r)
		cnt += merge(s, l, mid, r)
	}
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var s []int
	for i := 0; i < n; i++ {
		s = append(s, nextInt())
	}

	fmt.Println(MergeSort(s, 0, n))
}