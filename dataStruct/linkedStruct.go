package dataStruct

type LinkedNode2 struct {
	Data int
	Next *LinkedNode2
}

func (node *LinkedNode2) New() *LinkedNode2 {
	return &LinkedNode2{0,nil}
}
func AddNode(node *LinkedNode2, data int) bool {
	head := node
	last := node
	for head != nil {
		last = head
		head = head.Next
	}
	head = &LinkedNode2{Data:data}
	last.Next = head
	node = last
	return true
}