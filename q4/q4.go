package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {

	if basePrice <= 0 {

		return 0, fmt.Errorf("preço base inválido")

	}

	porcentagemEstado := 0.0
	porcentagemImposto := 0.0

	if state == "SP" {

		porcentagemEstado = 0.1

	} else if state == "RJ" {

		porcentagemEstado = 0.15

	} else if state == "MG" {

		porcentagemEstado = 0.2

	} else if state == "ES" {

		porcentagemEstado = 0.25

	} else {

		porcentagemEstado = 0.3

	}

	if taxCode == 1 {

		porcentagemImposto = 0.05

	} else if taxCode == 2 {

		porcentagemImposto = 0.1

	} else if taxCode == 3 {

		porcentagemImposto = 0.15

	} else {

		return 0, fmt.Errorf("imposto não encontrado")

	}

	finalPrice := basePrice + basePrice*porcentagemEstado + basePrice*porcentagemImposto

	return finalPrice, nil

}
