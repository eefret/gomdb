package gomdb

import "testing"

func TestImdbSearchMovies(t *testing.T) {
	resp, err := Search("Fight Club", "1999")
	if err != nil {
		t.Error("Failed Search Movies")
	}
	if resp.Search[0].Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}

func TestImdbGetMovieByTitle(t *testing.T) {
	resp, err := MovieByTitle("Fight Club", "1999")
	if err != nil {
		t.Error("Failed GetMovieByTitle")
	}
	if resp.Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}

func TestImdbGetMovieByImdbID(t *testing.T) {
	resp, err := MovieByImdbID("tt0137523")
	if err != nil {
		t.Error("Failed GetMovieByImdbID")
	}
	if resp.Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}
