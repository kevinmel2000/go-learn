package main

func Auth(res http.ResponseWriter, req *http.Request){
  if req.Header.Get("API-KEY") != "secret123" {
    http.Error(res, "Fuck", 401)
  }
}
