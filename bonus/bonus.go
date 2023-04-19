package bonus

import "fmt"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {

	if characterLevel < 0 || enemyLevel < 0 {

		return 0, fmt.Errorf("nível inválido")

	}

	damage := 0
	diferenca := 0

	if characterLevel > enemyLevel {

		damage = characterLevel * 10

	} else if characterLevel < enemyLevel {

		damage = characterLevel * 5

	} else {

		damage = characterLevel * 7

	}

	if characterLevel > 100 {

		damage *= 20

	}

	if enemyLevel > 100 {

		damage *= 2

	}

	if characterLevel > enemyLevel {

		diferenca = characterLevel - enemyLevel

	} else if enemyLevel > characterLevel {

		diferenca = enemyLevel - characterLevel

	}

	if diferenca > 20 {

		damage *= 5

	} else if diferenca < 20 {

		damage *= 2

	}

	return damage, nil

}
