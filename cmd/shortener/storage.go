package main

type Storage struct {
	Urls map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		Urls: make(map[string]string),
	}
}

func (s *Storage) Set(key, value string) {
	s.Urls[key] = value
}

func (s *Storage) Get(key string) (string, bool) {
	val, ok := s.Urls[key]
	return val, ok
}
