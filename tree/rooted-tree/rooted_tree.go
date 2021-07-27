package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Id     int
	Parent int
	Depth  int
	Type string
	Children []int
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

func setType(node *Node) {
	switch {
	case node.Parent == -1:
		node.Type = "root"
	case node.Parent != -1 && len(node.Children) == 0:
		node.Type = "leaf"
	default:
		node.Type = "internal node"
	}
}

func depthCheck(s []Node, parentId int) {
	setType(&s[parentId])
	if s[parentId].Type != "root" {
		s[parentId].Depth++
	}
	if s[parentId].Type == "leaf" {
		return
	}
	for _, v := range s[parentId].Children {
		s[v].Depth += s[parentId].Depth
		depthCheck(s, v)
	}
}


func formatSlice(s []int) string {
	str := ""
	for i, v := range s {
		if i == 0 {
			str += strconv.Itoa(v)
		} else {
			str += ", " + strconv.Itoa(v)
		}
	}
	return str
}

func dump(s []Node) {
	for _, v := range s {
		fmt.Printf("node %d: parent = %d, depth = %d, %s, [%s]\n", v.Id, v.Parent, v.Depth, v.Type, formatSlice(v.Children))
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	s := make([]Node, n)
	for i := 0; i < n; i++ {
		s[i] = Node{Parent: -1, Depth: 0}
	}
	for range s {
		id := nextInt()
		s[id].Id = id
		k := nextInt()
		for j := 0; j < k; j++ {
			child := nextInt()
			s[id].Children = append(s[id].Children, child)
			s[child].Parent = id
		}
	}
	var rootId int
	for _, v := range s {
		setType(&v)
		if v.Type == "root" {
			rootId = v.Id
			break
		}
	}
	depthCheck(s, rootId)

	dump(s)
}