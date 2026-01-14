package main

import (
    // "net/http"
	// "github.com/muskiteer/GoCP/routes"
	"github.com/muskiteer/GoCP/registery"
	"log"
	"context"
)

func main() {
	ctx := context.Background()

// 	mux := http.NewServeMux()
// 	routes.SetupRoutes(mux)
	 manifest, err := registery.LoadToolManifest("/home/muskiteer/Desktop/GoCP/server/schema/tools.json")
    if err != nil {
        log.Fatal(err)
    }

    registry, err := registery.InitRegistry(manifest)
    if err != nil {
        log.Fatal(err)
    }
	result, err := registry.Execute(
		ctx,
		// context.Background(),
		"fetching_crypto",
		map[string]any{
			"coin":     "bitcoin",
			"currency": "usd",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result);
}