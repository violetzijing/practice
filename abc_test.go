package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbc(t *testing.T) {
	val := abc()
	assert.NotNil(t, val)
}
