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
	x := nextInt()
	y := nextInt()

	fmt.Println(Gcd(x, y))
}

func Gcd(x int, y int) int {
	for y > 0 {
		r := x % y
		x = y
		y = r
	}
	return x
}
