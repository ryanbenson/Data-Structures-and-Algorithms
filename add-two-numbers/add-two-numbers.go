package addtwonumbers

type listNode struct {
	num  int
	next *listNode
}

// Adds two list nodes togehter that are linked lists with an int in it
func addTwoNumbers(list1 *listNode, list2 *listNode, carry int) *listNode {
	res := &listNode{
		num: 1,
		next: nil,
	}
	return res
}