package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/rkoesters/xkcd"
)

func main() {

	m := martini.Classic()
	//Includes static files(css,img,js etc.)
	m.Use(martini.Static("static"))
	//Renderes the templates
	m.Use(render.Renderer())
	m.Get("/", func(r render.Render) {
		//Get the comic data
		comic, err := xkcd.GetCurrent()
		perror(err)
		r.HTML(200, "index", comic)
	})
	m.Run()
}

//Error handling
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
