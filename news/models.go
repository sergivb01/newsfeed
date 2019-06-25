package news

type withChannels struct {
	Channel struct {
		Items []struct {
			Title       string `xml:"title"`
			PublishDate string `xml:"pubDate"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

type withoutChannels struct {
	Items []struct {
		Title       string `xml:"title"`
		PublishDate string `xml:"published"`
		Content     struct {
			Text string `xml:",chardata"`
		} `xml:"content"`
		Link struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
	} `xml:"entry"`
}
