package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sergivb01/newsfeed/news/engines"
	"github.com/sergivb01/newsfeed/store"

	"github.com/sergivb01/newsfeed/news"
	"github.com/sergivb01/newsfeed/routes"
	"github.com/sergivb01/newsfeed/server"
	"github.com/sergivb01/newsfeed/utils"
)

var (
	logger *log.Logger
)

func main() {
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)

	server.Server = server.Srv{
		Something: "yes",
		Storage:   store.InMemory(),
		Sources: []news.Source{
			engines.HackerNewsSource,
			engines.WiredSource,
			engines.TheVergeSource,
			engines.NYTimesSource,
		},
	}

	srv := &http.Server{
		Addr:         ":80",
		Handler:      utils.Tracing()(utils.Logging(logger)(configRoutes())),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Listen for CTRL+C or kill and start shutting down the app without
	// disconnecting people by not taking any new requests. ("Graceful Shutdown")
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go gracefulShutdown(done, quit, srv)

	go automaticFetch()

	logger.Println("Server is ready to handle requests")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen: %v\n", err)
	}

	<-done
	logger.Println("Server stopped")
}

func configRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/api/sources", routes.HandleSources)
	r.HandleFunc("/api/news", routes.HandleNews)
	return r
}

func automaticFetch() {
	ticker := time.NewTicker(time.Second * 15)
	for t := range ticker.C {
		log.Printf("started to fetch resources at %s", t)

		var items []news.Item
		for id, source := range server.Server.Sources {
			srcItems, err := source.GetItems()
			if err != nil {
				log.Printf("error getting items for %s: %v", source.Title, err)
			}
			log.Printf("Got items %d for %s", len(srcItems), source.Title)
			items = append(items, srcItems...)

			server.Server.Sources[id].LastFetched = time.Now()
		}

		server.Server.Lock()
		server.Server.Set(items)
		server.Server.Unlock()
	}
}

func gracefulShutdown(done chan bool, quit <-chan os.Signal, srv *http.Server) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
