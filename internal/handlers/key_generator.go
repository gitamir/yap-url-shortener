package handlers

import (
	"math/rand"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length = 8
)

type Repository interface {
	Set(string, string)
	Get(string) (string, bool)
}

type Generator struct {
	storage Repository
}

func NewGenerator(storage Repository) *Generator {
	return &Generator{
		storage: storage,
	}
}

func (g *Generator) Generate() string {
	return randString(length)
}

func randStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randString(length int) string {
	return randStringWithCharset(length, charset)
}
