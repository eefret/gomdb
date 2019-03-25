package gomdb

import (
	"os"
	"testing"
)

var apiKey = os.Getenv("OMDB_API_KEY")

func TestNoKey(t *testing.T) {
	api := Init("")

	_, err := api.Search(&QueryData{Title: "Her"})
	if err == nil {
		t.Errorf("Expected to fail")
	}
	if err != nil {
		expectedErrorMsg := "Status Code 401 received from IMDB"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Expected- %s, Got- %s", expectedErrorMsg, err)
		}
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		query *QueryData
		title string
		year  string
	}{
		{&QueryData{Title: "Fight Club", Year: "1999", SearchType: MovieSearch},
			"Fight Club",
			"1999",
		},
		{&QueryData{Title: "Her"},
			"Her",
			"2013",
		},
		{&QueryData{Title: "Macbeth", Year: "2015"},
			"Macbeth",
			"2015",
		},
	}

	api := Init(apiKey)

	for i, item := range tests {
		resp, err := api.Search(item.query)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if resp.Search[0].Title != item.title {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.title, resp.Search[0].Title)
			continue
		}
		if resp.Search[0].Year != item.year {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.year, resp.Search[0].Year)
			continue
		}
	}
}

func TestFailSearch(t *testing.T) {
	tests := []struct {
		query *QueryData
	}{
		{&QueryData{Title: "Game of Thrones", Year: "2001"}},
		{&QueryData{Title: "Dexter", SearchType: EpisodeSearch}},
	}

	api := Init(apiKey)

	for i, item := range tests {
		_, err := api.Search(item.query)
		if err == nil {
			t.Errorf("Test[%d]: Got nil error", i)
			continue
		}
		// Checking for strings is bad. But the API might change.
		if err.Error() != "Movie not found!" {
			t.Errorf("Test[%d]: Unexpected value- %s", i, err)
			continue
		}
	}
}

func TestInvalidCategory(t *testing.T) {
	tests := []struct {
		query *QueryData
	}{
		{&QueryData{Title: "Game of Thrones", Year: "2001", SearchType: "bad"}},
		{&QueryData{Title: "Dexter", SearchType: "bad"}},
	}

	api := Init(apiKey)

	for i, item := range tests {
		_, err := api.Search(item.query)
		if err == nil {
			t.Errorf("Test[%d]: Got nil error", i)
			continue
		}
		// Checking for strings is bad. But the error type is formatted
		if err.Error() != "Invalid search category- bad" {
			t.Errorf("Test[%d]: Unexpected value- %s", i, err)
			continue
		}
	}
}

func TestMediaByTitle(t *testing.T) {
	tests := []struct {
		query *QueryData
		title string
		year  string
	}{
		{&QueryData{Title: "Fight Club", Year: "1999", SearchType: MovieSearch},
			"Fight Club",
			"1999",
		},
		{&QueryData{Title: "Her"},
			"Her",
			"2013",
		},
		{&QueryData{Title: "Macbeth", Year: "2015"},
			"Macbeth",
			"2015",
		},
		{
			&QueryData{Title: "Rick and Morty", Season: "1", SearchType: SeriesSearch},
			"Rick and Morty",
			"2013–",
		},
		{
			&QueryData{Title: "Rick and Morty", Season: "1", Episode: "8", SearchType: EpisodeSearch},
			"Rixty Minutes",
			"2014",
		},
	}

	api := Init(apiKey)

	for i, item := range tests {
		resp, err := api.MovieByTitle(item.query)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if resp.Title != item.title {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.title, resp.Title)
			continue
		}
		if resp.Year != item.year {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.year, resp.Year)
			continue
		}
	}
}

func TestMediaByImdbID(t *testing.T) {
	tests := []struct {
		query *QueryData
		title string
		year  string
	}{
		{&QueryData{ImdbId: "tt0137523", SearchType: MovieSearch},
			"Fight Club",
			"1999",
		},
		{&QueryData{ImdbId: "tt1798709", SearchType: MovieSearch},
			"Her",
			"2013",
		},
		{&QueryData{ImdbId: "tt2884018", SearchType: MovieSearch},
			"Macbeth",
			"2015",
		},
		{&QueryData{ImdbId: "tt3952222", Season: "1", SearchType: SeriesSearch},
			"Killjoys",
			"2015–",
		},
		{&QueryData{ImdbId: "tt0944947", Season: "1", Episode: "1", SearchType: EpisodeSearch},
			"Winter Is Coming",
			"2011",
		},
	}

	api := Init(apiKey)

	for i, item := range tests {
		resp, err := api.MovieByImdbID(item.query)
		if err != nil {
			t.Errorf("Test[%d]: %s", i, err)
			continue
		}
		if resp.Title != item.title {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.title, resp.Title)
			continue
		}
		if resp.Year != item.year {
			t.Errorf("Test[%d]: Expected- %s, Got- %s", i, item.year, resp.Year)
			continue
		}
	}
}
