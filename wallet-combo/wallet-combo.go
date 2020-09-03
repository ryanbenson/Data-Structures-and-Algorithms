package walletcombo

import "fmt"

func getCombo(menu []struct {
	value int
	name  string
}, wallet int) [][]string {
	var allCombos [][]string

	// start at the beginning of the menu
	// for each one, go through the entire menu, starting from 0
	// and keep collecting combinations until 0 has no more
	// then move to starting at 1, then 2...
	// once you're at the end of the menu, move the first loop up 1
	// keep going until you're at the end of the menu on the first loop

	for index, item := range menu {
		combo := []string{}
		currentComboCost := 0

		for j := index; j < len(menu); j++ {
			currentComboCost += item.value
			if currentComboCost <= wallet {
				combo = append(combo, item.name)
			}
			if currentComboCost >= wallet {
				continue
			}
		}

		allCombos = append(allCombos, combo)
	}

	fmt.Println(allCombos)

	return allCombos
}
