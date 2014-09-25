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
	//"log"
	"net/http"
	"strings"
)

//=======================================================================
//							Const
//=======================================================================
const baseUrl string = "http://www.omdbapi.com/?"
const plot string = "full"
const tomatoes string = "true"

//=======================================================================
//							Global vars
//=======================================================================

//=======================================================================
//							Structs
//=======================================================================

//Type for the Search Response
type SearchResult struct {
	Title  string
	Year   string
	ImdbId string
	Type   string
}

//Type that respond Search
type SearchResponse struct {
	Search   []SearchResult
	Response string
	Error    string
}

//Type for searching a specific movie
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

//Search for movies given a Title and year, Year is optional you can pass nil
func SearchMovies(title string, year string) (*SearchResponse, error) {
	resp, err := omdbApiRequest(title, "", "", year)
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

//returns a MovieResult given Title
func GetMovieByTitle(title string, year string) (*MovieResult, error) {
	resp, err := omdbApiRequest("", "", title, year)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}
	return r, nil
}

// returns a MovieResult given a ImdbId ex:"tt2015381"
func GetMovieByImdbId(id string) (*MovieResult, error) {
	resp, err := omdbApiRequest("", id, "", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := new(MovieResult)
	err = json.NewDecoder(resp.Body).Decode(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if r.Response == "False" {
		return r, errors.New(r.Error)
	}
	return r, nil
}

func omdbApiRequest(s string, i string, t string, y string) (resp *http.Response, err error) {
	//s = Search Parameter, if this is != nil then its a searchMovies
	//i = Id Parameter, if this is != nil then its a getMovieByImdbId
	//t = Title Parameter, if this is != nil then its a getMovieByTitle
	//y = Year Parameter, Optional data for s and t search
	//var res http.Response
	var url string
	if s != "" {
		s = strings.Replace(s, " ", "%20", -1)
		url = fmt.Sprintf(baseUrl+"s=%s&y=%s", s, y)
	} else if i != "" {
		url = fmt.Sprintf(baseUrl+"i=%s&plot=%s&tomatoes=%s", i, plot, tomatoes)
	} else if t != "" {
		t = strings.Replace(t, " ", "%20", -1)
		url = fmt.Sprintf(baseUrl+"t=%s&plot=%s&tomatoes=%s&y=%s", t, plot, tomatoes, y)
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
		return errors.New(fmt.Sprintf("Status Code %d received from IMDB", status))
	} else {
		return nil
	}
}

//Stringer Interface for MovieResult
func (mr MovieResult) String() string {
	return fmt.Sprintf("#%s: %s (%s)", mr.ImdbID, mr.Title, mr.Year)
}

//Stringer Interface for SearchResult
func (sr SearchResult) String() string {
	return fmt.Sprintf("#%s: %s (%s) Type: %s", sr.ImdbId, sr.Title, sr.Year, sr.Type)
}
