package day8

import (
	mapset "github.com/deckarep/golang-set"
	"math"
	"strings"
)

// SegmentPattern is a set containing all of the runes that make up the
// signals which constitute a digit on a seven segment display
type SegmentPattern mapset.Set

func parseSegmentPattern(s string) SegmentPattern {
	r := mapset.NewSet()
	for _, runeValue := range s {
		r.Add(runeValue)
	}
	return r
}

// SegmentKey contains all of the possible segment patterns 0-9. They are stored
// in a map, where the key is the pattern length.
type SegmentKey map[int][]SegmentPattern

func parseSegmentKey(pattern string) SegmentKey {
	p := make(SegmentKey)
	for _, s := range strings.SplitN(pattern, " ", 10) {
		segmentPattern := parseSegmentPattern(s)
		p[segmentPattern.Cardinality()] = append(p[segmentPattern.Cardinality()], segmentPattern)
	}
	return p
}

func (p *SegmentKey) GetCipher() SegmentCipher {
	c := make(SegmentCipher)
	c[(*p)[2][0]] = 1
	c[(*p)[3][0]] = 7
	c[(*p)[4][0]] = 4
	c[(*p)[7][0]] = 8
	sixes := (*p)[6]
	for i, pattern := range sixes {
		four := c.GetPattern(4)
		if four.IsSubset(pattern) {
			c[pattern] = 9
			sixes = append(sixes[:i], sixes[i+1:]...)
		}
	}
	for i, pattern := range sixes {
		if c.GetPattern(1).IsSubset(pattern) {
			c[pattern] = 0
			sixes = append(sixes[:i], sixes[i+1:]...)
		}
	}
	c[sixes[0]] = 6
	fives := (*p)[5]
	for i, pattern := range fives {
		if c.GetPattern(1).IsSubset(pattern) {
			c[pattern] = 3
			fives = append(fives[:i], fives[i+1:]...)
		}
	}
	for i, pattern := range fives {
		if pattern.IsSubset(c.GetPattern(6)) {
			c[pattern] = 5
			fives = append(fives[:i], fives[i+1:]...)
		}
	}
	c[fives[0]] = 2
	return c
}

// Message contains the encrypted four segment patterns in the message.
type Message [4]SegmentPattern

func parseMessage(msg string) Message {
	var m Message
	for i, s := range strings.SplitN(msg, " ", 4) {
		m[i] = parseSegmentPattern(s)
	}
	return m
}

type EncryptedMessage struct {
	Key     SegmentKey
	Message Message
}

func (s *EncryptedMessage) Decrypt() DecryptedMessage {
	c := s.Key.GetCipher()
	var message DecryptedMessage
	for i, pattern := range s.Message {
		message[i] = c.Decode(pattern)
	}
	return message
}

func parseEncryptedMessage(input string) EncryptedMessage {
	splits := strings.SplitN(input, " | ", 2)
	return EncryptedMessage{
		Key:     parseSegmentKey(splits[0]),
		Message: parseMessage(splits[1]),
	}
}

type DecryptedMessage [4]int

func (m *DecryptedMessage) ToInt() int {
	s := 0
	for i := 0; i < 4; i++ {
		s += (*m)[i] * int(math.Pow(10, float64(3-i)))
	}
	return s
}

// SegmentCipher is a map from encrypted segment patterns to their decrypted
// integer values.
type SegmentCipher map[SegmentPattern]int

func (c *SegmentCipher) Decode(p SegmentPattern) int {
	for k, v := range *c {
		if k.Equal(p) {
			return v
		}
	}
	return 0
}

func (c *SegmentCipher) GetPattern(n int) SegmentPattern {
	for k, v := range *c {
		if v == n {
			return k
		}
	}
	return nil
}
