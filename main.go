package main

import (
	"fmt"

	"github.com/theJuan1112/monster-slayer-game/action"
	"github.com/theJuan1112/monster-slayer-game/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // Player or Monster

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)

}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0 // Makes sure that the special attack is available every 3 rounds
	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	fmt.Println(userChoice)

	var playerdmg int
	var monsterdmg int
	var playerheal int

	if userChoice == "ATTACK" {
		playerdmg = action.AttackMonster(false)

	} else if userChoice == "HEAL" {
		playerheal = action.HealPlayer()
	} else {
		playerdmg = action.AttackMonster(true)
	}

	monsterdmg = action.MonsterAttack()

	playerHealth, monsterHealth := action.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:        userChoice,
		PlayerHealth:  playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttack:  playerdmg,
		PlayerHeal:    playerheal,
		MonsterAttack: monsterdmg,
	}

	interaction.PrintRoundStats(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {

	interaction.GetWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
