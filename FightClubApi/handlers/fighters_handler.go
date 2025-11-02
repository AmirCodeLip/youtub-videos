package handlers

import (
	"encoding/json"
	"fight_club/entities"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samber/lo"
)

type fightersHandler struct {
	fighters []entities.User
}

func NewFightersHandler() *fightersHandler {
	fh := fightersHandler{}
	// ------------read all users from json --------------------
	fightersStr, err1 := os.ReadFile("./fighters.json")
	if err1 != nil {
		panic(err1)
	}
	// read fighters and parse them.
	fighters := []entities.User{}
	json.Unmarshal(fightersStr, &fighters)
	fh.fighters = fighters
	return &fh
}

func (h *fightersHandler) GetFighters(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(h.fighters); err != nil {
		fmt.Println(err)
	}
}

func (h *fightersHandler) GetFighter(w http.ResponseWriter, r *http.Request) {
	// get id from url
	vars := mux.Vars(r)
	id := vars["id"]
	numId, err2 := strconv.ParseInt(id, 10, 64)
	if err2 != nil {
		panic(err2)
	}
	// find fighter by id
	fighter, isNotNull := lo.Find(h.fighters, func(user entities.User) bool {
		return user.Id == int(numId)
	})
	/// if fighter found return it
	if isNotNull {
		/// try to encode fighter to json and return it
		if err := json.NewEncoder(w).Encode(fighter); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("the fighter with id %d is not exist", numId)
	}
}
