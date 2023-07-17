package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s", app.Domain)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1991-02-17")

	test1 := models.Movie{
		ID:          1,
		Title:       "tes1",
		Description: "desc2",
		ReleaseDate: rd,
		MPAARating:  "tes mpaa",
		RunTime:     110,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, test1)

	rd, _ = time.Parse("2006-01-02", "1985-11-01")

	test2 := models.Movie{
		ID:          2,
		Title:       "tes2",
		Description: "desc2",
		ReleaseDate: rd,
		MPAARating:  "asd",
		RunTime:     150,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, test2)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
