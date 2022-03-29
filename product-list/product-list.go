package productlist

import (
	"errors"
)

type productList struct {
	list []int
}

func (p *productList) add(item int) {
	p.list = append(p.list, item)
}

func (p *productList) get(num int) (int, error) {
	listLen := len(p.list)
	if num > listLen {
		return 0, errors.New("Invalid range of products. Exceeds amount of products in the list")
	}

	if num == 0 {
		return 0, nil
	}

	var startIndex int = 0
	startIndex = listLen - num
	desiredProducts := p.list[startIndex:]
	totalProductVal := 1
	for _, productVal := range desiredProducts {
		totalProductVal *= productVal
	}
	return totalProductVal, nil
}