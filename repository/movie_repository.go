package repository

import "irisdemo20191120/model"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := model.Movie{Name: "《这是一个挑食的可爱胖子》"}
	return movie.Name
}
