package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	name string
	time int
}

type Queue []Data

func (q *Queue) Enqueue(data Data) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() (Data, error) {
	if q.Empty() {
		return Data{name: "", time: 0}, fmt.Errorf("queue is empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, nil
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func NewQueue() *Queue {
	q := new(Queue)
	return q
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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	quantum := nextInt()
	sumTime := 0
	q := NewQueue()

	for i := 0; i < n; i++ {
		name := next(); time := nextInt()
		data := Data{name, time}
		q.Enqueue(data)
	}

	for ; !q.Empty(); {
		data, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if data.time > quantum {
			data.time -= quantum
			sumTime += quantum
			q.Enqueue(data)
		} else {
			sumTime += data.time
			fmt.Println(data.name, sumTime)
		}
	}
}
