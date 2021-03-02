package main

import (
	"bufio"
	"fmt"
	"math"
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
	n := nextInt()

	var a []int
	for i := 0; i < n; i++ {
		tmp := nextInt()
		a = append(a, tmp)
	}

	fmt.Println(CountPrime(a))
}

func CountPrime (a []int) (count int) {
	count = 0
	for _, v := range a {
		if IsPrime(v) {
			count++
		}
	}
	return
}

func IsPrime (n int) bool {
	if n == 2 {
		return true
	} else if n < 2 || n & 1 == 0 {
		return false
	}

	i := 3
	for ; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}