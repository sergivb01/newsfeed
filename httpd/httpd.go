package main

import (
	"context"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sergivb01/newsfeed/httpd/routes"
	"github.com/sergivb01/newsfeed/httpd/routes/api"
	"github.com/sergivb01/newsfeed/httpd/utils"
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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

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
		"checkIfValid": func(x int) bool {
			return x%4 == 0 || x == 0
		},
		"checkIfValid2": func(x int) bool {
			x++
			return x%4 == 0 && x != 0
		},
		"rndString": func() string {
			b := make([]byte, 8)
			for i := range b {
				b[i] = charset[seededRand.Intn(len(charset))]
			}
			return string(b)
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
