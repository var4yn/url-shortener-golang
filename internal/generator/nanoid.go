package generator

import gonanoid "github.com/matoous/go-nanoid/v2"

type NanoidGenerator struct{}

func (ng NanoidGenerator) Generate() string {
	id, _ := gonanoid.New(10)
	return id
}
