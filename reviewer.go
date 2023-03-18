/*
ikku is Ikku("一句") detector, Ikku is something like Japanese Haiku("俳句").

Inspired by [r7kamura/ikku].

[r7kamura/ikku]: https://github.com/r7kamura/ikku
*/
package ikku

import (
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/kurochan/ikku-go/internal"
	"github.com/samber/lo"
)

type Reviewer struct {
	tokenizer *tokenizer.Tokenizer
	exactly   bool
	rule      []int
}

func NewReviewer(dict *dict.Dict, opts ...ReviewerOption) (*Reviewer, error) {
	option := defaultReviewerConf
	for _, o := range opts {
		o.Apply(&option)
	}

	t, err := tokenizer.New(dict, tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	return &Reviewer{
		tokenizer: t,
		exactly:   option.exactly,
		rule:      option.rule,
	}, nil
}

func (r *Reviewer) Find(text string) *Song {
	nodes := lo.Map(r.tokenizer.Tokenize(text),
		func(token tokenizer.Token, _ int) internal.Node { return *internal.NewNode(token) },
	)
	for i := range nodes {
		scanner := internal.NewScanner(nodes[i:], r.exactly, r.rule)
		if song := scanner.Scan(); song != nil && song.Valid() {
			return internalSongToSong(song)
		}
	}
	return nil
}

func (r *Reviewer) Judge(text string) bool {
	nodes := lo.Map(r.tokenizer.Tokenize(text),
		func(token tokenizer.Token, _ int) internal.Node { return *internal.NewNode(token) },
	)
	song := internal.NewScanner(nodes, r.exactly, r.rule).Scan()
	return song != nil
}

func (r *Reviewer) Search(text string) []Song {
	nodes := lo.Map(r.tokenizer.Tokenize(text),
		func(token tokenizer.Token, _ int) internal.Node { return *internal.NewNode(token) },
	)
	songs := make([]Song, 0)
	for i := range nodes {
		scanner := internal.NewScanner(nodes[i:], r.exactly, r.rule)
		if song := scanner.Scan(); song != nil && song.Valid() {
			songs = append(songs, *internalSongToSong(song))
		}
	}
	return songs
}

type ReviewerOption interface {
	Apply(*reviewerOption)
}

type reviewerOption struct {
	rule    []int
	exactly bool
}

var defaultReviewerConf reviewerOption = reviewerOption{
	rule:    []int{5, 7, 5},
	exactly: false,
}

type reviewerOptionExactly bool

// Requires exact match.
func ReviewerOptionExactly(exactly bool) reviewerOptionExactly {
	return reviewerOptionExactly(exactly)
}
func (o reviewerOptionExactly) Apply(c *reviewerOption) {
	c.exactly = bool(o)
}

type reviewerOptionRule []int

// Customize count rule.
func ReviewerOptionRule(rule []int) reviewerOptionRule {
	return reviewerOptionRule(rule)
}
func (o reviewerOptionRule) Apply(c *reviewerOption) {
	c.rule = []int(o)
}
