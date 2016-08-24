package main

import (
	"bytes"
	"fmt"
)

type InSet struct {
	word []int64
}

func main() {
	var x, y InSet
	x.Add(1)
	x.Add(13)
	x.Add(5)
	fmt.Println(x.String())

	y.Add(34)
	y.Add(2)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
}

func (s *InSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return len(s.word) > int(word) && s.word[x]&(1<<bit) != 0
}

func (s *InSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for len(s.word) <= word {
		s.word = append(s.word, 0)
	}
	s.word[word] |= 1 << bit
}

func (s *InSet) UnionWith(p *InSet) {
	for k, v := range p.word {
		if len(s.word) <= k {
			s.word = append(s.word, v)
		} else {
			s.word[k] |= v
		}
	}
}

func (s *InSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range s.word {
		if v == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if v&(1<<uint(i)) != 0 {
				//if buf.Len() > len("{") {
				//	buf.WriteByte('}')
				//}
				fmt.Fprintf(&buf, " %d ", 64*k+i)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
