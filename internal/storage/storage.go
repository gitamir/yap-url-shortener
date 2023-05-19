package storage

import (
	"fmt"
	"sync"
)

type Storage struct {
	urls sync.Map
}

func NewStorage() *Storage {
	return &Storage{
		urls: sync.Map{},
	}
}

func (s *Storage) Set(key, value string) {
	s.urls.Store(key, value)
}

func (s *Storage) Get(key string) (string, bool) {
	val, ok := s.urls.Load(key)
	stringValue := fmt.Sprintf("%v", val)
	return stringValue, ok
}
