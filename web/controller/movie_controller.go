package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"irisdemo20191120/repository"
	"irisdemo20191120/service"
)

type MovieController struct {
	ctx iris.Context
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repository.NewMovieManager()
	movieService := service.MovieService(movieRepository)
	movieResult := movieService.GetMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}
