package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {

	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}

	i := 0
	maior := numbers[0]
	menor := numbers[0]
	soma := 0.0
	media := 0.0

	for i = 0; i < len(numbers); i++ {

		if numbers[i] > maior {
			maior = numbers[i]
		}

		if numbers[i] < menor {
			menor = numbers[i]
		}

		soma += float64(numbers[i])
	}

	media = (soma - float64(maior) - float64(menor)) / float64(len(numbers)-2)

	return menor, maior, media, nil

}
