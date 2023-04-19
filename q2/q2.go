package q2

import (
	"fmt"
	"strings"
	"unicode"
)

func AverageLettersPerWord(text string) (float64, error) {

	soma := 0.0
	media := 0.0
	onlyLetters := ""

	for _, ranText := range text {

		if unicode.IsLetter(ranText) || string(ranText) == " " {

			onlyLetters += string(ranText)

		}

	}

	newText := strings.ReplaceAll(onlyLetters, "!", "")
	newText = strings.ReplaceAll(newText, ",", "")
	newText = strings.ReplaceAll(newText, ".", "")
	newText = strings.ReplaceAll(newText, "(", "")
	newText = strings.ReplaceAll(newText, ")", "")
	newText = strings.ReplaceAll(newText, "[", "")
	newText = strings.ReplaceAll(newText, "]", "")
	newText = strings.ReplaceAll(newText, "{", "")
	newText = strings.ReplaceAll(newText, "}", "")
	newText = strings.ReplaceAll(newText, "?", "")

	if len(strings.ReplaceAll(newText, " ", "")) == 0 {

		return media, fmt.Errorf("texto vazio")

	}

	sliceText := strings.Split(newText, " ")

	for _, ranSliceText := range sliceText {
		soma += float64(len(ranSliceText))
	}

	media = soma / float64(len(sliceText))

	return media, nil

}
