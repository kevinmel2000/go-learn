package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	//goauth2 "github.com/golang/oauth2"
	//"github.com/martini-contrib/oauth2"
	//"github.com/martini-contrib/sessions"
	)

func init(){
	//martini.Env = martini.Prod
	martini.Env = martini.Dev
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
	  Layout: "home",
	}))

  m.Get("/", func(r render.Render){
    r.HTML(200, "halo", "ibnu")
  })

  m.Get("/api", Auth, func(r render.Render) {
    r.JSON(200, map[string]interface{}{"hello": "world","id":1})
  })

  m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request){
	if req.Header.Get("API-KEY") != "secret123" {
		http.Error(res, "Fuck", 401)
	}
}
