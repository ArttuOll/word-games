package wordmaster

import (
	"log"
	"net/http"
	"word-games/cmd/web/wordmaster/gamestate"
	"word-games/internal/database"
)

func WordMasterHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		solution, err := db.Queries().GetRandomWord(r.Context())
		if err != nil {
			http.Error(w, "Error fetching puzzle solution", http.StatusInternalServerError)
			log.Printf("Error fetching puzzle solution: %e", err)
		} else {
			gamestate.UpdateGameState(1, solution.Name, w)
		}

		component := WordMaster()
		err = component.Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf("Error rendering in WordMasterfHandler: %e", err)
		}
	}
}
