package main

type node struct {
	val  int
	prev *node
	next *node
}

type FrontMiddleBackQueue struct {
	head *node
	tail *node
	cnt  int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.cnt++
	if this.head == nil {
		this.head = &node{val: val}
		this.tail = this.head
		return
	}

	h := this.head
	this.head = &node{
		val:  val,
		next: this.head,
	}
	h.prev = this.head
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.cnt == 0 {
		this.PushFront(val)
		return
	}

	tar := this.head
	for i := 1; i < this.cnt/2; i++ {
		tar = tar.next
	}

	this.cnt++

	n := &node{
		val:  val,
		next: tar.next,
		prev: tar,
	}
	tar.next = n
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.cnt++
	if this.tail == nil {
		this.head = &node{val: val}
		this.tail = this.head
		return
	}

	h := this.tail
	this.tail = &node{
		val:  val,
		prev: this.tail,
	}
	h.next = this.tail
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.head == nil {
		return -1
	}

	this.cnt--

	old := this.head

	this.head = this.head.prev
	if this.head != nil {
		this.head.prev = nil
	}

	if this.cnt == 0 {
		this.tail = nil
	}

	old.next = nil
	return old.val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.cnt == 0 {
		return -1
	}

	if this.cnt == 1 {
		return this.PopFront()
	}

	tar := this.head
	for i := 1; i < this.cnt/2; i++ {
		tar = tar.next
	}

	if this.cnt%2 != 0 {
		tar = tar.next
	}
	this.cnt--

	if tar.prev != nil {
		tar.prev.next = tar.next
	}

	if tar.next != nil {
		tar.next.prev = tar.prev
	}

	tar.next = nil
	tar.prev = nil
	return tar.val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.tail == nil {
		return -1
	}
	this.cnt--

	old := this.tail

	this.tail = this.tail.prev
	if this.tail != nil {
		this.tail.next = nil
	}

	if this.cnt == 0 {
		this.head = nil
	}

	old.prev = nil
	return old.val
}

func main() {

	obj := Constructor()
	obj.PushFront(1)
	obj.PushBack(2)
	obj.PushMiddle(3)
	obj.PushMiddle(4)

	println(obj.PopFront())  // return 1 -> [4, 3, 2]
	println(obj.PopMiddle()) // return 3 -> [4, 2]
	println(obj.PopMiddle()) // return 4 -> [2]
	println(obj.PopBack())   // return 2 -> []
	println(obj.PopFront())
}
