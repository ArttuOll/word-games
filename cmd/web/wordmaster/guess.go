package wordmaster

import (
	"fmt"
	"net/http"
	"strings"
)

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	guess := strings.Join(r.Form["guess"], "")

	fmt.Println("Received guess: ", guess)

	// name := r.FormValue("name")
	// component := HelloPost(name)
	// err = component.Render(r.Context(), w)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	// }
}
