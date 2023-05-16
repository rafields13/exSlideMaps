// Escreva uma função que receba um slice de inteiros e retorne um novo slice
// contendo apenas os valores únicos, ou seja, sem duplicatas. Utilize um mapa
// para auxiliar na remoção das duplicatas.

package main

import "fmt"

func removeDuplicatedMoreValues(values []int) []int {
	occurrences := make(map[int]int)
	var onlySingleValues []int

	for _, value := range values {
		occurrences[value]++
		if occurrences[value] == 1 {
			onlySingleValues = append(onlySingleValues, value)
		}
	}
	return onlySingleValues
}

func main() {
	values := []int{1, 1, 2, 2, 3, 4, 5}
	onlySingleValues := removeDuplicatedMoreValues(values)
	fmt.Print(onlySingleValues)
}
