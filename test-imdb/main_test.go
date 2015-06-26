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

package main

import (
	imdb "github.com/eefret/go-imdb"
	"testing"
)

func TestMovieSearch(t *testing.T) {
	//Testing SearchMovies
	res, err := imdb.SearchMovies("The fifth element", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.Search[0].Title)
}

func TestGetMovieByTitle(t *testing.T) {
	//Testing GetMovieByTitle
	res2, err := imdb.GetMovieByTitle("True Grit", "1969")
	if err != nil {
		t.Error(err)
	}
	t.Log(res2.Title)
}

func TestGetMovieByImdbId(t *testing.T) {
	//Testing GetMovieByImdbId
	res3, err := imdb.GetMovieByImdbID("tt2015381")
	if err != nil {
		t.Error(err)
	}
	t.Log(res3.Title)
}
