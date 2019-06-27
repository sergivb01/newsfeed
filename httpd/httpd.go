package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/httpd/routes"
	"github.com/sergivb01/newsfeed/httpd/routes/api"
	"github.com/sergivb01/newsfeed/httpd/utils"
	"github.com/sergivb01/newsfeed/news"
)

func router() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/api/news", api.HandleNews)
	r.HandleFunc("/api/sources", api.HandleSources)
	r.HandleFunc("/api/client", api.HandleClient)
	r.HandleFunc("/", routes.HandleIndex)

	fs := http.FileServer(http.Dir("www/assets"))
	r.Handle("/assets/", http.StripPrefix("/assets/", fs))

	return r
}

func newHTTPServer(done chan bool, logger *log.Logger) {
	srv := &http.Server{
		Addr:         ":80",
		Handler:      utils.Tracing()(utils.Logging(logger)(router())),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go gracefulShutdown(done, quit, srv, logger)

	funcMap := template.FuncMap{
		"checkValid": func(x int) bool {
			return x%3 == 0 || x == 0
		},
		"checkValid2": func(x int) bool {
			return (x+1)%3 == 0
		},
		"renderCard": func(item news.Item) map[string]interface{} {
			src := client.CLI.GetSourceByName(item.Source)
			return map[string]interface{}{
				"Item":   item,
				"Source": src,
			}
		},
	}

	routes.Templates = template.Must(template.New("T").Funcs(funcMap).ParseGlob("www/templates/*"))
	// routes.Templates = template.Must(template.ParseGlob("www/templates/*")).Funcs(funcMap)

	logger.Println("Server is ready to handle requests")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen: %v\n", err)
	}
}

func gracefulShutdown(done chan bool, quit <-chan os.Signal, srv *http.Server, logger *log.Logger) {
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
