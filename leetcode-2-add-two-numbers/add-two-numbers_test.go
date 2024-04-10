package addtwonumbers

import (
	"testing"
)

func makeListNode(list []int) *listNode {
	res := &listNode{
		num: list[0],
	}
	temp := res

	for i := 1; i < len(list); i++ {
		temp.next = &listNode{num: list[i]}
		temp = temp.next
	}

	return res
}

func TestAddTwoNumbers(t *testing.T) {
	list1 := makeListNode([]int{2, 4, 3})
	list2 := makeListNode([]int{9, 1, 4})
	results := addTwoNumbers(list1, list2)

	if results.num != 1 {
		t.Errorf("Sum of numbers is incorrect, got: %v, want: %v.", results, 1)
	}

	// move to our next item
	results = results.next
	if results.num != 6 {
		t.Errorf("Sum of numbers is incorrect, got: %v, want: %v.", results, 6)
	}

	// move to our next item
	results = results.next
	if results.num != 7 {
		t.Errorf("Sum of numbers is incorrect, got: %v, want: %v.", results, 7)
	}
}
