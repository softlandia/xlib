package xlib

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tSortBytes = struct {
	i []byte
	o []byte
}

var dSortBytes = []tSortBytes{
	tSortBytes{
		[]byte{0, 'u', 'c', 'k'},
		[]byte{0, 'c', 'k', 'u'}},
	tSortBytes{
		[]byte{0, 'c', 'k', 0},
		[]byte{0, 0, 'c', 'k'}},
	tSortBytes{
		[]byte{'f', 'u', 'c', 'k', 'e', 'm', 'a', 'l', 'l'},
		[]byte{'a', 'c', 'e', 'f', 'k', 'l', 'l', 'm', 'u'}},
}

func TestSortBytes(t *testing.T) {
	for _, d := range dSortBytes {
		c := SortBytes(d.i)
		assert.Equal(t, true, bytes.Equal(c, d.o), fmt.Sprintf("%v\n", c))
	}
}
