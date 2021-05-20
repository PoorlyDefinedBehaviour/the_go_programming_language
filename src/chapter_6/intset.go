package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) has(x int) bool {
	word := x / 64
	bit := uint(x % 64)

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) add(x int) {
	word := x / 64
	bit := uint(x % 64)

	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
}

func (s *IntSet) union(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buffer bytes.Buffer

	buffer.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buffer.Len() > len("{") {
					buffer.WriteByte(' ')
				}

				fmt.Fprintf(&buffer, "%d", 64*i+j)
			}
		}

	}

	buffer.WriteByte('}')

	return buffer.String()
}

func main() {
	set := IntSet{}

	set.add(1)
	set.add(2)
	set.add(2)
	set.add(3)

	fmt.Println(set.String())
}
