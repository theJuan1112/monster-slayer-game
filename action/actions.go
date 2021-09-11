package action

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_VALUE
	maxAttackValue := PLAYER_ATTACK_MAX_VALUE

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_VALUE
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_VALUE
	}

	dgmValue := randomNumber(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dgmValue
	return dgmValue
}

func HealPlayer() int {

	healValue := randomNumber(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healthDif := PLAYER_HEALTH - currentPlayerHealth

	if healthDif >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDif
	}

}

func MonsterAttack() int {
	minAttackValue := MONSTER_ATTACK_MIN_VALUE
	maxAttackValue := MONSTER_ATTACK_MAX_VALUE

	dgmValue := randomNumber(minAttackValue, maxAttackValue)
	currentPlayerHealth -= dgmValue
	return dgmValue
}

func randomNumber(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}
