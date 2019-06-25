package store

import "github.com/sergivb01/newsfeed/news"

type RedisStore struct {
}

func (s *RedisStore) Get() []news.Item {
	return nil
}

func (s *RedisStore) Set(items []news.Item) {

}

func inRedis() *RedisStore {
	return &RedisStore{}
}
