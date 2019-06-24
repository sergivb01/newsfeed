package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/httpd/store"
	"github.com/sergivb01/newsfeed/news/sources"

	"github.com/sergivb01/newsfeed/news"
)

var (
	logger        *log.Logger
	fetchInterval = time.Second * 30
)

func main() {
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)

	client.NewClient(time.Second*10, []news.Source{
		sources.MicrosiervosSource,
		sources.WiredSource,
		sources.TheVergeSource,
		sources.GizmodoSource,
		sources.HackerNewsSource,
		sources.NYTimesSource,
	})
	client.CLI.UseStore(store.InMemory())

	go automaticFetch()

	done := make(chan bool)
	newHTTPServer(done, logger)

	<-done
	logger.Println("Server stopped")
}

func automaticFetch() {
	for {
		t := time.Now()

		var tempStorage struct {
			items []news.Item
			sync.Mutex
		}
		var wg sync.WaitGroup
		wg.Add(len(client.CLI.Sources))

		for _, source := range client.CLI.Sources {
			go func(source news.Source) {
				srcItems, err := source.GetItems()
				if err != nil {
					log.Printf("error getting items for %s: %v", source.Title, err)
				}
				tempStorage.Lock()
				tempStorage.items = append(tempStorage.items, srcItems...)
				tempStorage.Unlock()

				wg.Done()
			}(source)
		}
		wg.Wait()

		client.CLI.Lock()
		client.CLI.Store.Set(tempStorage.items)
		client.CLI.Unlock()

		client.CLI.LastFetched = time.Now()
		client.CLI.NextUpdate = t.Add(fetchInterval)

		log.Printf("It took %s to fetch %d items from %d sources!", time.Since(t), len(tempStorage.items), len(client.CLI.Sources))
		log.Printf("Fetching again at %s (in %s)", time.Now().Add(fetchInterval), fetchInterval)

		time.Sleep(fetchInterval)
	}
}
