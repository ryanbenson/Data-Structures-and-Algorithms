package walletcombo

import "fmt"

func getCombo(menu []struct {
	value int
	name  string
}, wallet int) [][]string {
	var allCombos [][]string

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
