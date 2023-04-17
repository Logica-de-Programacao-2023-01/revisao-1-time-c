package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {

	soma := 0.0
	media := 0.0

	apenasLetras := strings.ReplaceAll(text, "!", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, ",", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, ".", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, "(", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, ")", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, "[", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, "]", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, "{", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, "}", "")
	apenasLetras = strings.ReplaceAll(apenasLetras, "?", "")

	sliceText := strings.Split(apenasLetras, " ")

	if len(sliceText) == 0 {
		return media, fmt.Errorf("texto vazio")
	}

	for _, ranSliceText := range sliceText {
		soma += float64(len(ranSliceText))
	}

	media = soma / float64(len(sliceText))

	return media, nil

}
