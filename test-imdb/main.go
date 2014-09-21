package main

import (
	"fmt"
	imdb "github.com/eefret/go-imdb"
	"log"
)

func main() {

	//Testing SearchMovies
	res, err := imdb.SearchMovies("The fifth element", "")
	if err != nil {
		log.Fatal(err)
	}

	//Testing GetMovieByTitle
	res2, err := imdb.GetMovieByTitle("True Grit", "1969")
	if err != nil {
		log.Fatal(err)
	}

	//Testing GetMovieByImdbId
	res3, err := imdb.GetMovieByImdbId("tt2015381")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Search[0].Title)
	fmt.Println(res2.Title)
	fmt.Println(res3.Title)
}
