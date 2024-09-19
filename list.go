package main

type listNode struct {
	data string
	next *listNode
	prev *listNode
}

type List struct {
	head       *listNode
	tail       *listNode
	capacity   int
	population int
}

func NewList(capacity int) *List {
	return &List{
		capacity: capacity,
	}
}

func (l *List) InsertHead(data string) string {
	var removedNodeData string

	if l.full() {
		removedNodeData = l.removeTail()
	}

	newNode := &listNode{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}

	l.population++
	return removedNodeData
}

func (l *List) Delete(data string) {
	if l.empty() {
		return
	}

	current := l.head
	for current != nil {
		if current.data == data {
			if current.prev == nil {
				l.head = current.next
			} else {
				current.prev.next = current.next
			}

			if current.next == nil {
				l.tail = current.prev
			} else {
				current.next.prev = current.prev
			}

			l.population--
			return
		}

		current = current.next
	}
}

func (l *List) removeTail() string {
	if l.empty() {
		return ""
	}

	removedNodeData := l.tail.data

	if l.tail.prev == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}

	l.population--
	return removedNodeData
}

func (l *List) peek() string {
	return l.head.data
}

func (l *List) empty() bool {
	return l.head == nil && l.tail == nil
}

func (l *List) full() bool {
	return l.population == l.capacity
}
