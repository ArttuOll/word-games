package wordmaster

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
)

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	guess := strings.Join(r.Form["guess"], "")

	fmt.Println("Received guess: ", guess)

	solution, err := r.Cookie("solution")
	fmt.Println("Solution is: ", solution.Value)

	colors := getColorsForWord(strings.ToLower(guess), strings.ToLower(solution.Value))

	fmt.Printf("Colors are: %v\n", colors)

	// name := r.FormValue("name")
	// component := HelloPost(name)
	// err = component.Render(r.Context(), w)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	// }
}

type Color = string

const (
	White  Color = "white"
	Yellow Color = "yellow"
	Green  Color = "green"
	Gray   Color = "gray"
)

func getColorsForWord(guess, solution string) []Color {
	result := []Color{White, White, White, White, White}
	characterPool := strings.Split(solution, "")

	for i, c := range guess {
		if c == rune(solution[i]) {
			result[i] = Green
			characterPool[i] = ""
		}
	}

	for i, c := range guess {
		if result[i] == Green {
			continue
		}

		foundIndex := slices.Index(characterPool, string(c))
		if foundIndex != -1 {
			result[i] = Yellow
			characterPool[foundIndex] = ""
		} else {
			result[i] = Gray
		}
	}

	return result
}
