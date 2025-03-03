package main

import (
	"container/list"
	"fmt"
)

type number_t int32

type pair_t struct {
	num   int
	count int
}

func find_short_path(graph [][]int, begin int, end int) int {
	queue_next := new(list.List)
	array_used := make([]bool, len(graph))
	count := 0
	current := begin

	queue_next.PushBack(pair_t{begin, 0})
	array_used[begin] = true
	for {
		vert := queue_next.Front().Value.(pair_t)
		current = vert.num
		count = vert.count
		queue_next.Remove(queue_next.Front())
		for i := range graph[current] {
			if !array_used[graph[current][i]] {
				array_used[graph[current][i]] = true
				queue_next.PushBack(pair_t{graph[current][i], count + 1})
			}
		}
		if current == end {
			break
		}
	}
	return count
}

type Node struct {
	element int
	lst     *list.List
	left    *Node
	right   *Node
}

type Tree struct {
	root   *Node
	cached *Node
}

func (tree *Tree) AddElement(new_element int) {
	if tree.root == nil {
		tree.root = &Node{new_element, new(list.List), nil, nil}
	} else {
		tree.root.AddElement(new_element)
	}
}

func (tree *Tree) In(element int) bool {
	if tree.cached != nil {
		if element == tree.cached.element {
			return true
		} else {
			tree.cached = tree.root.In(element)
			if tree.cached != nil {
				return true
			}
		}
	}
	return false
}

func (tree *Tree) GetList(element int) *list.List {
	if tree.cached != nil {
		if element == tree.cached.element {
			return tree.cached.lst
		}
	}
	tree.cached = tree.root.In(element)
	if tree.cached != nil {
		return tree.cached.lst
	}
	return nil
}

func (head *Node) AddElement(new_element int) {
	current := head
	less := false
	var previous *Node = head

	for current != nil {
		previous = current
		if new_element < current.element {
			less = true
			current = current.left
		} else if new_element > current.element {
			less = false
			current = current.right
		} else {
			return
		}
	}
	if less {
		previous.left = new(Node)
		current = previous.left
	} else {
		previous.right = new(Node)
		current = previous.right
	}
	current.element = new_element
	current.left = nil
	current.right = nil
	current.lst = new(list.List)
}

func (head *Node) In(element int) *Node {
	current := head
	for current != nil && current.element != element {
		if current.element > element {
			current = current.left
		} else if current.element < element {
			current = current.right
		} else {
			return current
		}
	}
	return nil
}

func create_graph(islands []int) [][]int {
	var numMap *Tree = &Tree{
		root:   nil,
		cached: nil,
	}
	for i := range islands {
		tmpLst := numMap.GetList(islands[i])
		if tmpLst == nil {
			numMap.AddElement(islands[i])
		}
		numMap.GetList(islands[i]).PushBack(i)
	}
	graph := make([][]int, len(islands))
	if len(islands) > 1 {
		tmpLst := numMap.GetList(islands[0])
		graph[0] = make([]int, tmpLst.Len()+1)
		graph[0][0] = 1
		var index = 1
		for elem := tmpLst.Front(); elem != nil; elem = elem.Next() {
			graph[0][index] = elem.Value.(int)
			index++
		}
		tmpLst = numMap.GetList(islands[len(graph)-1])
		graph[len(graph)-1] = make([]int, tmpLst.Len()+1)
		graph[len(graph)-1][0] = len(graph) - 2
		index = 1
		for elem := tmpLst.Front(); elem != nil; elem = elem.Next() {
			graph[len(graph)-1][index] = elem.Value.(int)
			index++
		}
	}
	for i := 1; i < len(graph)-1; i++ {
		tmpLst := numMap.GetList(islands[i])
		graph[i] = make([]int, tmpLst.Len()+2)
		graph[i][0] = i - 1
		graph[i][1] = i + 1
		index := 2
		for elem := tmpLst.Front(); elem != nil; elem = elem.Next() {
			graph[i][index] = elem.Value.(int)
			index++
		}
	}
	return graph
}

func main() {
	var n = 0
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Printf("некорректный ввод количества элементов: %e", err)
	}
	arr := make([]int, n)
	if arr == nil {
		fmt.Printf("Не хватает памяти!")
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
	graph := create_graph(arr)
	for i := range graph {
		for j := range graph[i] {
			fmt.Printf("%d ", graph[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Result = %d", find_short_path(graph, 0, n-1))
}
