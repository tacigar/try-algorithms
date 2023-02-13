package main

import (
	"fmt"
	"math"
	"strconv"
)

const M = 100000

type HashTableElem struct {
	key   string
	value string
}

type HashTable [M]*HashTableElem

func Hash(s string) int {
	res := 1
	for _, c := range s {
		n := int(byte(c))
		res *= n
	}
	return int(math.Mod(float64(res), M))
}

func (h *HashTable) Insert(key, value string) {
	hash := Hash(key)
	for i := 0; ; i++ {
		if h[hash] == nil {
			h[hash] = &HashTableElem{key: key, value: value}
			return
		} else {
			hash = Hash(key + strconv.Itoa(i))
		}
	}
}

func (h *HashTable) At(key string) string {
	hash := Hash(key)
	for i := 0; ; i++ {
		if h[hash] == nil {
			return ""
		}
		if h[hash].key == key {
			return h[hash].value
		}
		hash = Hash(key + strconv.Itoa(i))
	}
}

func main() {
	var h HashTable
	h.Insert("key1", "value1")
	h.Insert("key2", "value2")
	h.Insert("key3", "value3")
	fmt.Println(h.At("key1"))
	fmt.Println(h.At("key2"))
	fmt.Println(h.At("key3"))
	fmt.Println(h.At("key4"))
}
