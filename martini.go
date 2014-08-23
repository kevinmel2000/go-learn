package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
  Layout: "home",
}))

  m.Get("/", func(r render.Render){
    r.HTML(200, "halo", "jeremy")
  })

  m.Get("/api", func(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world"})
  })

  m.Run()
}
