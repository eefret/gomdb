/*
Copyright 2014 Kaissersoft Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package imdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

//=======================================================================
//							Const
//=======================================================================
const baseURL string = "http://www.omdbapi.com/?"
const plot string = "full"
const tomatoes string = "true"

//=======================================================================
//							Global vars
//=======================================================================

//=======================================================================
//							Structs
//=======================================================================

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

//SearchMovies search for movies given a Title and year, Year is optional you can pass nil
func SearchMovies(title string, year string) (*SearchResponse, error) {
	resp, err := omdbAPIRequest(title, "", "", year)
	if err != nil {
		return nil, err
	}

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

//GetMovieByTitle returns a MovieResult given Title
func GetMovieByTitle(title string, year string) (*MovieResult, error) {
	resp, err := omdbAPIRequest("", "", title, year)
	if err != nil {
		return nil, err
	}

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

//GetMovieByImdbID returns a MovieResult given a ImdbID ex:"tt2015381"
func GetMovieByImdbID(id string) (*MovieResult, error) {
	resp, err := omdbAPIRequest("", id, "", "")
	if err != nil {
		return nil, err
	}

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

func omdbAPIRequest(s string, i string, t string, y string) (resp *http.Response, err error) {
	//s = Search Parameter, if this is != nil then its a searchMovies
	//i = Id Parameter, if this is != nil then its a getMovieByImdbID
	//t = Title Parameter, if this is != nil then its a getMovieByTitle
	//y = Year Parameter, Optional data for s and t search
	//var res http.Response
	var url string
	if s != "" {
		s = strings.Replace(s, " ", "%20", -1)
		url = fmt.Sprintf(baseURL+"s=%s&y=%s", s, y)
	} else if i != "" {
		url = fmt.Sprintf(baseURL+"i=%s&plot=%s&tomatoes=%s", i, plot, tomatoes)
	} else if t != "" {
		t = strings.Replace(t, " ", "%20", -1)
		url = fmt.Sprintf(baseURL+"t=%s&plot=%s&tomatoes=%s&y=%s", t, plot, tomatoes, y)
	} else {
		return nil, errors.New("Invalid Request")
	}
	//log.Print(url) //DEBUG
	res, err := http.Get(url)
	err = checkErrorStatus(res.StatusCode)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkErrorStatus(status int) error {
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
