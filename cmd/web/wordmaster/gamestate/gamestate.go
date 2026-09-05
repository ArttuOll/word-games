package gamestate

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

type GameState struct {
	Solution string `json:"solution"`
	Attempt  int    `json:"attempt"`
}

func UpdateGameState(attempt int, solution string, w http.ResponseWriter) {
	gameState := GameState{
		Solution: solution,
		Attempt:  attempt,
	}

	data, err := json.Marshal(gameState)

	if err != nil {
		log.Printf("error marshaling game state struct to json: %v", err)
	}

	encoded := base64.RawURLEncoding.EncodeToString(data)

	http.SetCookie(w, &http.Cookie{
		Name:  "gamestate",
		Value: encoded,
	})
}

func ReadGameState(r *http.Request) GameState {
	cookie, err := r.Cookie("gamestate")
	if err != nil {
		log.Printf("error reading gamestate cookie: %v", err)
	}

	decoded, err := base64.RawURLEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Printf("error decoding cookie value: %v", err)
	}

	var gameState GameState
	err = json.Unmarshal(decoded, &gameState)
	if err != nil {
		log.Printf("error unmarshaling game state: %v", err)
	}

	return gameState
}
