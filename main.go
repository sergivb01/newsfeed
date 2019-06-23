package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/sergivb01/newsfeed/news"
	"github.com/sergivb01/newsfeed/news/engines"
)

func main() {
	os.Mkdir("output", 0544)

	sources := map[string]news.Source{
		"hn":       engines.HackerNewsSource,
		"wired":    engines.WiredSource,
		"theverge": engines.TheVergeSource,
		"nytimes":  engines.NYTimesSource,
	}

	var wg sync.WaitGroup
	wg.Add(len(sources))

	for name, source := range sources {
		go doThings(name, source, &wg)
	}

	wg.Wait()
}

func doThings(name string, source news.Source, wg *sync.WaitGroup) {
	items, err := source.GetItems()
	if err != nil {
		fmt.Printf("error while getting items for %s: %v\n", name, err)
	} else {
		fmt.Printf("Found %d news in %s!\n", len(items), name)
	}

	b, _ := json.Marshal(items)
	ioutil.WriteFile(fmt.Sprintf("output/%s.json", name), b, 0644)

	wg.Done()
}
