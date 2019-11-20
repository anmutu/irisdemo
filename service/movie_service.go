package service

import "irisdemo20191120/repository"

type MovieService interface {
	GetMovieName() string
}

type MovieServiceManager struct {
	repo repository.MovieRepository
}

func NewMovieServiceManager(repo repository.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) GetMovieName() string {
	return m.repo.GetMovieName()
}
