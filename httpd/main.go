package main

import (
	"log"
	"os"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/httpd/store"
)

func main() {
	if err := client.NewClient("config.yml"); err != nil {
		log.Printf("error while loading config: %v", err)
		return
	}
	client.CLI.UseStore(store.InMemory())

	go func() {
		for {
			if err := client.CLI.FetchSources(); err != nil {
				log.Printf("error while fetching sources: %v", err)
			}
		}
	}()

	done := make(chan bool)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	newHTTPServer(done, logger)

	<-done
	logger.Println("Server stopped")
}
