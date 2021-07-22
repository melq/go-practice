package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}

	v := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return v, nil
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func NewStack() *Stack {
	s := new(Stack)
	return s
}

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func numCheck(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func doCalc(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return 0
	}
}

func main() {
	s := NewStack()

	slice := strings.Split(nextLine(), " ")
	for _, v := range slice {
		if numCheck(v) {
			n, _ := strconv.Atoi(v)
			s.Push(n)
		} else {
			var a, b int
			var err error
			b, err = s.Pop()
			a, err = s.Pop()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			s.Push(doCalc(a, b, v))
		}
	}

	fmt.Println(strconv.Itoa((*s)[0]))
}
