package main

import (
	"encoding/json"
	"fight_club/entities"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func readCommand() string {
	var cmd string
	fmt.Print("Enter command: ")
	fmt.Scan(&cmd)
	return cmd
}

func main() {
	fullCmd := strings.Split(readCommand(), ",")
	// ------------read all users from json --------------------
	fightersStr, err1 := os.ReadFile("./fighters.json")
	if err1 != nil {
		panic(err1)
	}

	// read fighters and parse them.
	fighters := []entities.User{}
	json.Unmarshal(fightersStr, &fighters)
	cmd := fullCmd[0]
	switch cmd {
	case "all":
		fmt.Println(fighters)
	case "getById":
		numId, err2 := strconv.ParseInt(fullCmd[1], 10, 64)
		if err2 != nil {
			panic(err2)
		}
		fighter, isNotNull := lo.Find(fighters, func(user entities.User) bool {
			return user.Id == int(numId)
		})
		if isNotNull {
			fmt.Printf("the fighter with id %d is %s %s \n ", numId, fighter.Name, fighter.LastName)
		} else {
			fmt.Printf("the fighter with id %d is not exist", numId)

		}

	}
}
