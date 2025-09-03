package generator_test

import (
	"testing"
	"url-shortener/internal/generator"

	"github.com/stretchr/testify/assert"
)

func TestLengthId(t *testing.T) {

	gen := generator.NanoidGenerator{}

	result := gen.Generate()

	assert.Equal(t, len(result), 10)

}
