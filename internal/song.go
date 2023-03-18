package internal

type Song struct {
	Phrases [][]Node
	Exactly bool
	Rule    []int
}

func NewSong(phrases [][]Node, exactly bool, rule []int) *Song {
	return &Song{
		Phrases: phrases,
		Exactly: exactly,
		Rule:    rule,
	}
}

func (s *Song) Valid() bool {
	switch {
	case len(s.Phrases) == 0:
		return false
	default:
		return true
	}
}
