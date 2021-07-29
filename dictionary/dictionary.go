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

	dict := make(map[string]bool, n)

	for i := 0; i < n; i++ {
		command := next()
		str := next()
		switch command {
		case "insert":
			dict[str] = true
		case "find":
			if _, exist := dict[str]; exist {
				fmt.Println("yes")
			}else {
				fmt.Println("no")
			}
		}
	}
}
