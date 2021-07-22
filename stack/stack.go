package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, err := strconv.Atoi(next())
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	//n := nextInt()

	s := NewStack()
	s.Push(1)
	s.Push(2)
	v, _ := s.Pop()
	fmt.Println(v)
	v, _ = s.Pop()
	fmt.Println(v)
	//var s []int
	//for i := 0; i < n; i++ {
	//	tmp := nextInt()
	//	s = append(s, tmp)
	//}
	//
	//fmt.Println(n, s)
}
