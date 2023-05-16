// Escreva uma função que receba um mapa onde as chaves são nomes de alunos e os
// valores são slices de notas. A função deve calcular a média das notas de cada aluno e
// retornar um novo mapa onde as chaves são os nomes dos alunos e os valores são as
// médias correspondentes.

package main

import "fmt"

func calculateAverage(nameGrade map[string][]float64) map[string]float64 {
	nameAverage := make(map[string]float64)

	for student, grades := range nameGrade {
		sum := 0.0
		for _, grade := range grades {
			sum += grade
		}
		average := sum / float64(len(grades))
		nameAverage[student] = average
	}
	return nameAverage
}

func main() {
	nameGrade := map[string][]float64{
		"João":   {6.5, 7, 9},
		"Maria":  {8, 9, 10},
		"Rafael": {10, 10, 10},
	}
	nameAvarege := calculateAverage(nameGrade)
	fmt.Print(nameAvarege)
}
