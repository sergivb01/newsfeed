package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sergivb01/newsfeed/news"
	"gopkg.in/yaml.v2"
)

type config struct {
	FetchInterval string        `yaml:"fetchInterval"`
	Sources       []news.Source `yaml:"sources"`
}

func loadConfig(filePath string) (config, error) {
	var c config

	if err := checkIfExists(filePath); err != nil {
		return c, fmt.Errorf("error creating default config: %v", err)
	}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return c, err
	}
	if err := yaml.Unmarshal(b, &c); err != nil {
		return c, err
	}

	return c, nil
}

func checkIfExists(filePath string) error {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return nil
	}
	log.Printf("config file does not exists, creating a \"%s\" configuration for you...", filePath)

	return ioutil.WriteFile(filePath, []byte(defaultConfig), 0644)
}

const defaultConfig = `fetchInterval: "10s"
sources:
  - title: Microsiervos
    shortName: microsiervos
    homepage: https://www.microsiervos.com
    rss: https://www.microsiervos.com/index.xml
    withChannels: true

  - title: Wired
    shortName: wired
    homepage: https://wired.com
    rss: https://www.wired.com/feed/rss
    withChannels: true

  - title: The Verge
    shortName: verge
    homepage: https://www.theverge.com/
    rss: https://www.theverge.com/rss/index.xml
    withChannels: false

  - title: Gizmodo
    shortName: gizmodo
    homepage: https://es.gizmodo.com
    rss: https://es.gizmodo.com/rss
    withChannels: true

  - title: Hacker News
    shortName: hackernews
    homepage: https://news.ycombinator.com
    rss: https://hnrss.org/frontpage
    withChannels: true

  - title: The New York Times
    shortName: nytimes
    homepage: https://nytimes.com
    rss: https://www.nytimes.com/svc/collections/v1/publish/https://www.nytimes.com/section/technology/rss.xml
    withChannels: true
`
