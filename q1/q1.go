package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {

	historico := 0.0

	if currentPurchase <= 0 {
		return currentPurchase, fmt.Errorf("valor da compra inválido")
	}

	for _, ranPurchaseHistory := range purchaseHistory {
		historico += ranPurchaseHistory
	}

	if historico/float64(len(purchaseHistory)) > 1000 {
		currentPurchase = currentPurchase * 0.2
		return currentPurchase, nil
	}

	if len(purchaseHistory) == 0 {
		currentPurchase = currentPurchase * 0.1
		return currentPurchase, nil
	}

	if historico > 1000 {
		currentPurchase = currentPurchase * 0.1
		return currentPurchase, nil
	}

	if historico <= 1000 {
		currentPurchase = currentPurchase * 0.05
		return currentPurchase, nil
	}

	if historico <= 500 {
		currentPurchase = currentPurchase * 0.02
		return currentPurchase, nil
	}

	return 0, nil

}
