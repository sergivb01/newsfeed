package engines

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sergivb01/newsfeed/news"
)

var HackerNewsSource = news.Source{
	Title:    "Hacker News",
	Homepage: "https://news.ycombinator.com",
	RSS:      "https://hacker-news.firebaseio.com/v0/topstories.json",
	Parser: func(b []byte) ([]news.Item, error) {
		var res [500]int

		if err := json.Unmarshal(b, &res); err != nil {
			return nil, fmt.Errorf("error while marshaling bytes to array of int for HN: %v", err)
		}

		var items []news.Item

		for _, id := range res[:40] {
			if len(items) == 30 {
				break
			}

			url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
			resp, err := http.Get(url)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			var jsonRes struct {
				By    string `json:"by"`
				ID    int64  `json:"id"`
				Text  string `json:"text"`
				Time  int64  `json:"time"`
				Title string `json:"title"`
				Type  string `json:"type"`
				URL   string `json:"url"`
			}

			if err := json.Unmarshal(b, &jsonRes); err != nil {
				return nil, fmt.Errorf("error while marshaling post from HN to struct: %v", err)
			}

			// I don't want job offers in my feed
			if jsonRes.Type != "story" || jsonRes.URL == "" {
				continue
			}

			cut := len(jsonRes.Text)
			if cut > 175 {
				cut = 175
			}

			link := jsonRes.URL
			if link == "" {
				link = fmt.Sprintf("https://news.ycombinator.com/item?id=%d", id)
			}

			items = append(items, news.Item{
				Title:       jsonRes.Title,
				Link:        link,
				PublishDate: time.Unix(jsonRes.Time, 0),
				Description: jsonRes.Text[:cut],
			})
		}

		return items, nil
	},
}
