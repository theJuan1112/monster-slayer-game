package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action        string
	PlayerAttack  int
	PlayerHeal    int
	MonsterAttack int
	PlayerHealth  int
	MonsterHealth int
}

func PrintGreeting() {

	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)

	asciiFigure.Print()
	fmt.Println("Starting New Game")
	fmt.Println("GLHF")

}

func ShowAvailableActions(isSpecialRound bool) {
	fmt.Println("Please Choose your Action")
	fmt.Println("----------------------------")
	fmt.Println("1) Attack Monster")
	fmt.Println("2) Heal")

	if isSpecialRound {
		fmt.Println("3) Special Attack")
	}
}

func GetWinner(winner string) {
	asciiFigure := figure.NewColorFigure("GAME OVER!", "", "red", true)

	asciiFigure.Print()
	fmt.Printf("%v WON!\n", winner)
}

func PrintRoundStats(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked Monster for %v damage\n", roundData.PlayerAttack)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed special attack to Monster for %v damage\n", roundData.PlayerAttack)
	} else {
		fmt.Printf("Player heal for %v damage\n", roundData.PlayerHeal)
	}

	fmt.Printf("Monster attacked Player for %v damage\n", roundData.MonsterAttack)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func WriteLogFile(rounds *[]RoundData) {
	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("FAILED TO GET PATH, EXITING")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")

	if err != nil {
		fmt.Println("Saving log file failed! Exiting.")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttack),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttack),
			"Player Heal":           fmt.Sprint(value.PlayerHeal),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writting to log file failed, Exiting.")
			continue
		}
	}
	file.Close()
	fmt.Println("Wrote Data to Log!")
}
