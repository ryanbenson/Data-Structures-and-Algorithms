package addtwonumbers

type listNode struct {
	num  int
	next *listNode
}

// Adds two list nodes togehter that are linked lists with an int in it
func addTwoNumbers(list1 *listNode, list2 *listNode) *listNode {
	// our list
	head := &listNode{}

	// our current node which will move down as we go through the list
	currentNode := head

	// if there are any values to carry over from one list node to the next
	carry := 0

	// keep going until there are no items and no carry over values
	for list1 != nil || list2 != nil || carry != 0 {
		// start off with our current value as the carry from the previous
		sum := carry

		// if there's an item in either list, update the sum and move to the next item
		if list1 != nil {
			sum += list1.num
			list1 = list1.next
		}

		if list2 != nil {
			sum += list2.num
			list2 = list2.next
		}

		// reset our carry with any carry over (e.g. 11 / 10 = 1 gets carried over)
		carry = sum / 10
		// current node only wants the remainder of % 10 (e.g. 13 % 10 = 3)
		number := sum % 10

		// setup the next node
		currentNode.next = &listNode{ num: number }
		// move to the next one
		currentNode = currentNode.next
	}

	// return the start of the link list
	return head.next
}
