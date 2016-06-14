package gomdb

import "testing"

func TestImdbSearchMovies(t *testing.T) {
	query := &QueryData{Title: "Fight Club", Year: "1999", SearchType: MovieSearch}
	resp, err := Search(query)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.Search[0].Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}

func TestImdbGetMovieByTitle(t *testing.T) {
	query := &QueryData{Title: "Fight Club", Year: "1999"}
	resp, err := MovieByTitle(query)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}

func TestImdbGetMovieByImdbID(t *testing.T) {
	resp, err := MovieByImdbID("tt0137523")
	if err != nil {
		t.Error(err)
		return
	}
	if resp.Title != "Fight Club" {
		t.Error("Wrong Movie")
	}
}
