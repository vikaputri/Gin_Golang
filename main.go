package main

import "github.com/vikaputri/Gin_Golang/routers"

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
