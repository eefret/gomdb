// Package gomdb is a golang implementation of the OMDB API.
package gomdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL  = "http://www.omdbapi.com/?"
	plot     = "full"
	tomatoes = "true"
)

//SearchResult is the type for the search results
type SearchResult struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
}

//SearchResponse is the struct of the response in a search
type SearchResponse struct {
	Search   []SearchResult
	Response string
	Error    string
}

//MovieResult is the result struct of an specific movie search
type MovieResult struct {
	Title             string
	Year              string
	Rated             string
	Released          string
	Runtime           string
	Genre             string
	Director          string
	Writer            string
	Actors            string
	Plot              string
	Language          string
	Country           string
	Awards            string
	Poster            string
	Metascore         string
	ImdbRating        string
	ImdbVotes         string
	ImdbID            string
	Type              string
	TomatoMeter       string
	TomatoImage       string
	TomatoRating      string
	TomatoReviews     string
	TomatoFresh       string
	TomatoRotten      string
	TomatoConsensus   string
	TomatoUserMeter   string
	TomatoUserRating  string
	TomatoUserReviews string
	TomatoURL         string
	DVD               string
	BoxOffice         string
	Production        string
	Website           string
	Response          string
	Error             string
}

//=======================================================================
//							Funcs
//=======================================================================

//Search search for movies given a Title and year, Year is optional you can pass nil
func Search(title string, year string) (*SearchResponse, error) {
	resp, err := requestAPI(title, "", "", year)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(SearchResponse)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}

	return r, nil
}

//MovieByTitle returns a MovieResult given Title
func MovieByTitle(title string, year string) (*MovieResult, error) {
	resp, err := requestAPI("", "", title, year)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}
	return r, nil
}

//MovieByImdbID returns a MovieResult given a ImdbID ex:"tt2015381"
func MovieByImdbID(id string) (*MovieResult, error) {
	resp, err := requestAPI("", id, "", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}
	return r, nil
}

func requestAPI(s string, i string, t string, y string) (resp *http.Response, err error) {
	//s = Search Parameter, if this is != nil then its a searchMovies
	//i = Id Parameter, if this is != nil then its a getMovieByImdbID
	//t = Title Parameter, if this is != nil then its a getMovieByTitle
	//y = Year Parameter, Optional data for s and t search
	//var res http.Response

	var URL *url.URL
	URL, err = url.Parse(baseURL)

	if err != nil {
		return nil, err
	}
	URL.Path += "/"
	parameters := url.Values{}
	if len(s) > 0 {
		parameters.Add("s", s)
		parameters.Add("y", y)
	} else if len(i) > 0 {
		parameters.Add("i", i)
		parameters.Add("plot", plot)
		parameters.Add("tomatoes", tomatoes)
	} else if len(t) > 0 {
		parameters.Add("t", t)
		parameters.Add("plot", plot)
		parameters.Add("tomatoes", tomatoes)
		parameters.Add("y", y)
	} else {
		return nil, errors.New("Invalid Request")
	}
	URL.RawQuery = parameters.Encode()
	res, err := http.Get(URL.String())
	err = checkErr(res.StatusCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkErr(status int) error {
	if status != 200 {
		return fmt.Errorf("Status Code %d received from IMDB", status)
	}
	return nil
}

//Stringer Interface for MovieResult
func (mr MovieResult) String() string {
	return fmt.Sprintf("#%s: %s (%s)", mr.ImdbID, mr.Title, mr.Year)
}

//Stringer Interface for SearchResult
func (sr SearchResult) String() string {
	return fmt.Sprintf("#%s: %s (%s) Type: %s", sr.ImdbID, sr.Title, sr.Year, sr.Type)
}
