package imdb

import "testing"

func TestImdbSearchMovies(t *testing.T) {
	resp, err := SearchMovies("Fight Club", "1999")
	if err != nil {
		t.Error("Failed Search Movies")
	}
	if resp.Search[0].Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}

func TestImdbGetMovieByTitle(t *testing.T) {
	resp, err := GetMovieByTitle("Fight Club", "1999")
	if err != nil {
		t.Error("Failed GetMovieByTitle")
	}
	if resp.Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}

func TestImdbGetMovieByImdbID(t *testing.T) {
	resp, err := GetMovieByImdbID("tt0137523")
	if err != nil {
		t.Error("Failed GetMovieByImdbID")
	}
	if resp.Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}
