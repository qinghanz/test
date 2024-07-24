package main

import (
	"fmt"
	"gotth-quick-start/cmd/mockdb"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	_ = mockdb.New()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", GetScoresHandler)
	r.Post("/add", AddScoreHandler)
	http.ListenAndServe(":8080", r)
}

// handles request to main welcome page
func GetScoresHandler(w http.ResponseWriter, r *http.Request) {
	db := mockdb.GetDB()
	matches := db.GetMatches()
	templ.Handler(ScoresPage(matches)).ServeHTTP(w, r)
}

// handles the request to add a new match score.
func AddScoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WE ARE HERE")
	var db = mockdb.GetDB()
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	player1 := r.FormValue("player1")
	score1, _ := strconv.Atoi(r.FormValue("score1"))
	player2 := r.FormValue("player2")
	score2, _ := strconv.Atoi(r.FormValue("score2"))

	match := mockdb.Match{
		Player1: player1,
		Player2: player2,
		Score1:  score1,
		Score2:  score2,
		Date:    time.Now(),
	}

	db.AddMatch(match)
	templ.Handler(ScoresPage(db.GetMatches())).ServeHTTP(w, r)
}
