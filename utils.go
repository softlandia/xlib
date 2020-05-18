package xlib

import (
	"hash/fnv"
	"sort"
)

//Epsilon - precission
const Epsilon float64 = 0.01

// Max - return max of two int
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// SortBytes - return sorted slice of bytes
func SortBytes(b []byte) []byte {
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return b
}

// StrHash - make hash from string
func StrHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
