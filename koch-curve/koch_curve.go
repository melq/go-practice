package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	X float64
	Y float64
}

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

func dump(s Point) {
	fmt.Printf("%.8f %.8f\n", s.X, s.Y)
}

func kochCurve(n int, p1, p2 Point) {
	if n == 0 {
		return
	}
	s := Point{(2 * p1.X + p2.X) / 3, (2 * p1.Y + p2.Y) / 3}
	t := Point{(p1.X + 2 * p2.X) / 3, (p1.Y + 2 * p2.Y) / 3}
	u := Point{
		(t.X - s.X) * math.Cos(math.Pi / 3) - (t.Y - s.Y) * math.Sin(math.Pi / 3) + s.X,
		(t.X - s.X) * math.Sin(math.Pi / 3) + (t.Y - s.Y) * math.Cos(math.Pi / 3) + s.Y}
	kochCurve(n - 1, p1, s)
	dump(s)
	kochCurve(n - 1, s, u)
	dump(u)
	kochCurve(n - 1, u, t)
	dump(t)
	kochCurve(n - 1, t, p2)
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	p1 := Point{0.0, 0.0}
	p2 := Point{100.0, 0.0}
	dump(p1)
	kochCurve(n, p1, p2)
	dump(p2)
}
