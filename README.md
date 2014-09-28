The Golang Omdb API
=======

Author: Christopher T. Herrera (eefretsoul AT gmail DOT com)

<iframe src="http://githubbadge.appspot.com/eefret" style="border: 0;height: 142px;width: 200px;overflow: hidden;" frameBorder="0"></iframe>

This API uses the [omdbapi.com](http://omdbapi.com/) API by Brian Fritz

***
### OMDBAPI.com
This is an excellent open database for movie and film content.

I *strongly* encourage you to check it out and contribute to keep it growing.

### http://www.omdbapi.com
***
Project Usage
-------------
The API usage is very simple. Just import the go-imdb package

	import (
		imdb "github.com/eefret/go-imdb"
	)

And use any of the methods 

	res, err := imdb.SearchMovies("The fifth element", "")
	res2, err := imdb.GetMovieByTitle("True Grit", "1969")
	res3, err := imdb.GetMovieByImdbId("tt2015381")

See the project documentation to see the Response Objects and stuff

Project Documentation
---------------------
The automatically generated documentation can be found in godocs.
[![GoDoc](https://godoc.org/github.com/eefret/go-imdb?status.svg)](https://godoc.org/github.com/eefret/go-imdb)
