package main

import (
        "context"
        "log"

        "cricket/api"
        "cricket/config"
)

func main() {
        ctx := context.Background()
        cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(cfg.HTTPServer)
	server.Start(ctx)
}
