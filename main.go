package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	fmt.Println("Number of Players:")
	var cntPlayers int
	_, err := fmt.Scanln(&cntPlayers)
	if err != nil {
		log.Fatal(err)
	}

	players := make([]Player, cntPlayers+1)
	players[0] = Player{Name: "Bank", Balance: 0}

	for i := 1; i <= cntPlayers; i++ {
		fmt.Printf("Names of Player %d:", i)
		_, err = fmt.Scanln(&players[i].Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		fmt.Printf("\n\n+%s+\n", strings.Repeat("-", 30))
		fmt.Printf("|%3s | %10s | %10s |\n", "ID", "Name", "Balance")
		fmt.Printf("+%s+\n", strings.Repeat("-", 30))
		for i := 0; i <= cntPlayers; i++ {
			fmt.Printf("|%3d | %10s | %10d |\n", i, players[i].Name, players[i].Balance)
		}
		fmt.Printf("+%s+\n", strings.Repeat("-", 30))
		fmt.Println("")
		var from, to, amt int
		fmt.Printf("From:")
		fmt.Scanln(&from)
		fmt.Printf("To:")
		fmt.Scanln(&to)
		fmt.Printf("Amt:")
		fmt.Scanln(&amt)
		if to <= cntPlayers && to >= 0 && from <= cntPlayers && from >= 0 {
			players[to].Balance += amt
			players[from].Balance -= amt
			fmt.Println("Transaction Done")
		}

	}

}

type Player struct {
	Name    string
	Balance int
}
