package main

import (
    "net/http"
	"github.com/muskiteer/GoCP/server/routes"
)

func main() {

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)
}