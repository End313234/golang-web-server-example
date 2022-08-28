package utils

import (
	"errors"
)

type Movie struct {
	Name        string `json:"name"`
	ReleaseYear int    `json:"releaseYear"`
}

type Movies []Movie

func (movies *Movies) Add(movie Movie) *Movies {
	*movies = append(*movies, movie)
	return movies
}

func (movies *Movies) Update(movie Movie, newMovie Movie) (*Movies, error) {
	targetIndex := -1

	for index, item := range *movies {
		if item == movie {
			targetIndex = index
		}
	}

	if targetIndex == -1 {
		return &Movies{}, errors.New("movie not found")
	}

	(*movies)[targetIndex] = newMovie

	return movies, nil
}

func (movies *Movies) Get(movie Movie) (Movie, error) {
	for _, item := range *movies {
		if item == movie {
			return movie, nil
		}
	}

	return Movie{}, errors.New("movie not found")
}

func (movies *Movies) GetAll() Movies {
	return *movies
}

func (movies *Movies) Delete(movie Movie) (Movie, error) {
	var deletedMovie Movie

	for index, item := range *movies {
		if item == movie {
			deletedMovie = movie
			*movies = append((*movies)[:index], (*movies)[index+1:]...)
		}
	}

	if (deletedMovie == Movie{}) {
		return deletedMovie, errors.New("movie not found")
	}

	return deletedMovie, nil
}

var MoviesDatabase Movies
