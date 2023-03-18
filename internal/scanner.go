package internal

import (
	"github.com/samber/lo"
)

type Scanner struct {
	nodes   []Node
	exactly bool
	rule    []int
	phrases [][]Node
	count   int
}

func NewScanner(nodes []Node, exactly bool, rule []int) *Scanner {
	return &Scanner{
		nodes:   nodes,
		exactly: exactly,
		rule:    rule,
	}
}

func (s *Scanner) Scan() *Song {
	s.phrases = make([][]Node, len(s.rule))
	s.count = 0

	if !s.hasValidFirstNode() {
		return nil
	}
	for _, node := range s.nodes {
		if s.consume(node) {
			if s.satisfied() && !s.exactly {
				return NewSong(s.phrases, s.exactly, s.rule)
			}
		} else {
			return nil
		}
	}
	if s.satisfied() {
		return NewSong(s.phrases, s.exactly, s.rule)
	} else {
		return nil
	}
}

func (s *Scanner) consume(node Node) bool {
	switch {
	case node.PronunciationLength() > s.maxConsumableLength():
		return false
	case !node.ElementOfIkku():
		return false
	case s.firstOfPhrase() && !node.FirstOfPhrase():
		return false
	case node.PronunciationLength() == s.maxConsumableLength() && !node.LastOfPhrase():
		return false
	default:
		s.phrases[s.phraseIndex()] = append(s.phrases[s.phraseIndex()], node)
		s.count += node.PronunciationLength()
		return true
	}
}

func (s *Scanner) maxConsumableLength() int {
	return lo.Sum(s.rule[0:s.phraseIndex()+1]) - s.count
}

func (s *Scanner) hasValidFirstNode() bool {
	if len(s.nodes) == 0 {
		return false
	}
	return s.nodes[0].FirstOfIkku()
}

func (s *Scanner) hasFullCount() bool {
	return s.count == lo.Sum(s.rule)
}

func (s *Scanner) hasValidLastNode() bool {
	if len(s.phrases) == 0 {
		return false
	}
	lastPh := s.phrases[len(s.phrases)-1]
	if len(lastPh) == 0 {
		return false
	}
	return lastPh[len(lastPh)-1].LastOfIkku()
}

func (s *Scanner) firstOfPhrase() bool {
	c := 0
	for _, i := range s.rule {
		c += i
		if c == s.count {
			return true
		}
	}
	return false
}

func (s *Scanner) phraseIndex() int {
	index := len(s.rule) - 1
	for i := range s.rule {
		if s.count < lo.Sum(s.rule[0:i+1]) {
			index = i
			break
		}
	}
	return index
}

func (s *Scanner) satisfied() bool {
	return s.hasFullCount() && s.hasValidLastNode()
}
