package pkg

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
	Size int
}

func (list *LinkedList) Add(item interface{}) {
	newNode := Node{data: item}

	if list.Size == 0 {
		list.head = &newNode
		list.tail = &newNode
		list.Size = 1
	} else {
		list.tail.next = &newNode
		newNode.prev = list.tail
		list.tail = &newNode
		list.Size += 1
	}
}

func (list *LinkedList) AddLast(item interface{}) {
	list.Add(item)
}

func (list *LinkedList) AddFirst(item interface{}) {
	if list.Size == 0 {
		list.Add(item)
	} else {
		newNode := Node{data: item}
		newNode.next = list.head
		list.head.prev = &newNode
		list.head = &newNode
		list.Size++
	}
}

func (list *LinkedList) AddAtIndex(item interface{}, index int) error {
	if index < 0 || index > list.Size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		list.AddFirst(item)
	} else if index == list.Size {
		list.AddLast(item)
	} else {
		curr := list.head

		for i := 0; i < index - 1; i++ {
			curr = curr.next
		}

		newNode := Node{data: item}

		newNode.next = curr.next
		newNode.prev = curr

		newNode.next.prev = &newNode
		curr.next = &newNode

		list.Size++
	}

	return nil
}

func (list *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.Size {
		return nil, fmt.Errorf("index out of bounds")
	}

	curr := list.head

	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.data, nil
}

func (list *LinkedList) Peek() (interface{}, error) {
	if list.Size == 0 {
		return nil, fmt.Errorf("cannot peek into empty list")
	}

	return list.Get(0)
}

func (list *LinkedList) RemoveFirst() (interface{}, error) {
	if list.Size == 0 {
		return nil, fmt.Errorf("cannot remove first element of empty list")
	}

	val := list.head.data

	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	}
	list.Size--

	return val, nil
}

func (list *LinkedList) RemoveLast() (interface{}, error) {
	if list.Size == 0 {
		return nil, fmt.Errorf("cannot remove last element of empty list")
	}
	
	if list.Size == 1 {
		return list.RemoveFirst()
	}

	val := list.tail.data

	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	}
	list.Size--

	return val, nil
}

func (list *LinkedList) RemoveAtIndex(index int) (interface{}, error) {
	if index < 0 || index >= list.Size {
		return nil, fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		return list.RemoveFirst()
	}

	if index == list.Size - 1 {
		return list.RemoveLast()
	}

	curr := list.head

	for i := 0; i < index; i++ {
		curr = curr.next
	}

	val := curr.data

	curr.next.prev = curr.prev
	curr.prev.next = curr.next

	list.Size--

	return val, nil
}