package main

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length = 8
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type KeyGenerator interface {
	Generate() string
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
	newStr := randString(length)
	str, _ := g.storage.Get(newStr)
	if str != "" {
		return g.Generate()
	}
	return newStr
}

func randStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randString(length int) string {
	return randStringWithCharset(length, charset)
}
