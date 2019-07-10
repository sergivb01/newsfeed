<!--
@host = http://localhost
-->
# 📰 Newsfeed
Newsfeed is a Go program that automatically reads RSS feed from a bunch of online news websites
and organizes the feed into one. It stores the feed into an storage service (currently only memory).

# 🎯 TODO:
- [X] Add [Microsiervos](https://www.microsiervos.com/)
- [X] Add a `yaml` configuration file
- [ ] Add support for flags
- [ ] Website and/or design
  - [ ] Add a navigation bar
  - [ ] Find a better way to organize and display the news
  - [ ] Dark mode?
- [ ] Implement redis storage

# 📖 Documentation
Each API route in the HTTP server will return a response in JSON. If an error occurs, it will be displayed as the following json format:

### 🚧 Error
- `???` depending on the error
```json
{
  "message": "json: unsupported type: map[interface {}]interface {}",
  "data": "🖖🏻"
}
```

### 📰 News
GET {{host}}/api/news

- `200 OK`
- `500 InternalServerError` on error
```json
[
  {
    "title": "The 'Super Mario Maker 2' Community Is a Haven of Player Creativity",
    "link": "https://www.wired.com/story/super-mario-maker-2-community",
    "published": "2019-07-10T13:00:00Z",
    "description": "The hallowed halls of Mario have become, in the hands of fans, shrines to the gods of difficulty.",
    "source": "Wired"
  }
]
```

### 💻 Client
GET {{host}}/api/client

- `200 OK` on success
- `500 InternalServerError` on error
```json
{
  "sources": [
    {
      "title": "Wired",
      "homepage": "https://wired.com",
      "rss": "https://www.wired.com/feed/rss",
      "withChannels": true
    }
  ],
  "fetchInterval": 10000000000,
  "lastFetched": "2019-07-10T18:01:43.6479822+02:00",
  "nextUpdate": "2019-07-10T18:01:53.4002519+02:00"
}
```

### 📚 Sources
GET {{host}}/api/sources

- `200 OK` on success
- `500 InternalServerError` on error
```json
[
  {
    "title": "Wired",
    "homepage": "https://wired.com",
    "rss": "https://www.wired.com/feed/rss",
    "withChannels": true
  }
]
```

# 📋 License
This project is licensed under [MIT License](./LICENSE "License document from the repository")
