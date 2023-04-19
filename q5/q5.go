package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {

	if fromScale == toScale {

		if fromScale == "C" || fromScale == "F" || fromScale == "K" {

			return temp, nil

		} else {

			return 0, fmt.Errorf("escala inválida")

		}
	}

	switch fromScale {

	case "C":

		switch toScale {

		case "F":

			temp = temp*9/5 + 32
			return temp, nil

		case "K":

			temp = temp + 273.15
			return temp, nil

		default:

			return 0, fmt.Errorf("escala inválida")

		}

	case "F":

		switch toScale {

		case "C":

			temp = (temp - 32) * 5 / 9
			return temp, nil

		case "K":

			temp = (temp-32)*5/9 + 273.15
			return temp, nil

		default:

			return 0, fmt.Errorf("escala inválida")

		}

	case "K":

		switch toScale {

		case "C":

			temp = temp - 273.15
			return temp, nil

		case "F":

			temp = (temp-273.15)*9/5 + 32
			return temp, nil

		default:

			return 0, fmt.Errorf("escala inválida")

		}

	default:

		return 0, fmt.Errorf("escala inválida")

	}
}