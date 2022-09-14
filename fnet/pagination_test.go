package fnet_test

import (
	"testing"

	"github.com/lenon/gofii/fnet"
	"github.com/stretchr/testify/assert"
)

func TestPageOffset(t *testing.T) {
	assert.Equal(t, 0, fnet.PageOffset(1, 100))
	assert.Equal(t, 100, fnet.PageOffset(2, 100))
	assert.Equal(t, 200, fnet.PageOffset(3, 100))
	assert.Equal(t, 300, fnet.PageOffset(4, 100))
}

func TestNumberOfPages(t *testing.T) {
	assert.Equal(t, 0, fnet.NumberOfPages(0, 100))
	assert.Equal(t, 1, fnet.NumberOfPages(100, 100))
	assert.Equal(t, 6, fnet.NumberOfPages(531, 100))
	assert.Equal(t, 5, fnet.NumberOfPages(500, 100))
}
