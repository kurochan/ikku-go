package ikku

import (
	"github.com/kurochan/ikku-go/internal"
	"github.com/samber/lo"
)

type Song struct {
	Phrases [][]Node
	Exactly bool
	Rule    []int
}

func internalSongToSong(is *internal.Song) *Song {
	if is == nil {
		return nil
	}
	return &Song{
		Phrases: lo.Map(is.Phrases, func(ph []internal.Node, _ int) []Node {
			return lo.Map(ph, func(n internal.Node, _ int) Node { return *internalNodeToNode(&n) })
		}),
		Exactly: is.Exactly,
		Rule:    is.Rule,
	}
}
