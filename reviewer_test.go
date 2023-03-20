package ikku_test

import (
	"reflect"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/ikawaha/kagome-dict/ipa"
	ikku "github.com/kurochan/ikku-go"
	"github.com/samber/lo"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name string
		text string
		opts []ikku.ReviewerOption
		want [][]string
	}{
		{
			name: "ikku/古池や蛙飛び込む水の音",
			text: "古池や蛙飛び込む水の音",
			want: [][]string{
				{"古池", "や"},
				{"蛙", "飛び込む"},
				{"水", "の", "音"},
			},
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "ikku/exactly/古池や蛙飛び込む水の音",
			text: "古池や蛙飛び込む水の音",
			want: [][]string{
				{"古池", "や"},
				{"蛙", "飛び込む"},
				{"水", "の", "音"},
			},
			opts: []ikku.ReviewerOption{ikku.ReviewerOptionExactly(true)},
		},
		{
			name: "ikku/ああ古池や蛙飛び込む水の音がします",
			text: "ああ古池や蛙飛び込む水の音がします",
			want: [][]string{
				{"古池", "や"},
				{"蛙", "飛び込む"},
				{"水", "の", "音"},
			},
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "no ikku/exactly/古池や蛙飛び込む水の音がします",
			text: "古池や蛙飛び込む水の音がします",
			want: nil,
			opts: []ikku.ReviewerOption{ikku.ReviewerOptionExactly(true)},
		},
		{
			name: "no ikku/今日はいい天気ですね",
			text: "今日はいい天気ですね",
			want: nil,
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "ikku/久かたの光のどけき春の日にしづ心なく花の散るらん",
			text: "久かたの光のどけき春の日にしづ心なく花の散るらん",
			want: [][]string{
				{"久", "かた", "の"},
				{"光", "のどけき"},
				{"春", "の", "日", "に"},
				{"しづ", "心", "なく"},
				{"花", "の", "散る", "らん"},
			},
			opts: []ikku.ReviewerOption{ikku.ReviewerOptionRule([]int{5, 7, 5, 7, 7})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ikku.NewReviewer(ipa.Dict(), tt.opts...)
			if err != nil {
				t.Error(err)
			}
			got := songToTexts(r.Find(tt.text))
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestJudge(t *testing.T) {
	tests := []struct {
		name string
		text string
		opts []ikku.ReviewerOption
		want bool
	}{
		{
			name: "ikku/古池や蛙飛び込む水の音",
			text: "古池や蛙飛び込む水の音",
			want: true,
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "ikku/exactly/古池や蛙飛び込む水の音",
			text: "古池や蛙飛び込む水の音",
			want: true,
			opts: []ikku.ReviewerOption{ikku.ReviewerOptionExactly(true)},
		},
		{
			name: "ikku/古池や蛙飛び込む水の音がします",
			text: "古池や蛙飛び込む水の音がします",
			want: true,
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "no ikku/ああ古池や蛙飛び込む水の音がします",
			text: "ああ古池や蛙飛び込む水の音がします",
			want: false,
			opts: []ikku.ReviewerOption{ikku.ReviewerOptionExactly(true)},
		},
		{
			name: "no ikku/exactly/古池や蛙飛び込む水の音がします",
			text: "古池や蛙飛び込む水の音がします",
			want: false,
			opts: []ikku.ReviewerOption{ikku.ReviewerOptionExactly(true)},
		},
		{
			name: "no ikku/今日はいい天気ですね",
			text: "今日はいい天気ですね",
			want: false,
			opts: []ikku.ReviewerOption{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ikku.NewReviewer(ipa.Dict(), tt.opts...)
			if err != nil {
				t.Error(err)
			}
			got := r.Judge(tt.text)
			if tt.want != got {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name string
		text string
		opts []ikku.ReviewerOption
		want [][][]string
	}{
		{
			name: "ikku/古池や蛙飛び込む水の音",
			text: "古池や蛙飛び込む水の音",
			want: [][][]string{
				{
					{"古池", "や"},
					{"蛙", "飛び込む"},
					{"水", "の", "音"},
				},
			},
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "ikku/今日はいい天気ですね。古池や蛙飛び込む水の音。",
			text: "今日はいい天気ですね。古池や蛙飛び込む水の音。",
			want: [][][]string{
				{
					{"古池", "や"},
					{"蛙", "飛び込む"},
					{"水", "の", "音"},
				},
			},
			opts: []ikku.ReviewerOption{},
		},
		{
			name: "ikku/ああ古池や蛙飛び込む水の音。ああ五月雨を集めてはやし最上川。",
			text: "ああ古池や蛙飛び込む水の音。ああ五月雨を集めてはやし最上川。",
			want: [][][]string{
				{
					{"古池", "や"},
					{"蛙", "飛び込む"},
					{"水", "の", "音"},
				},
				{
					{"五月雨", "を"},
					{"集め", "て", "はやし"},
					{"最上川"},
				},
			},
			opts: []ikku.ReviewerOption{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ikku.NewReviewer(ipa.Dict(), tt.opts...)
			if err != nil {
				t.Error(err)
			}
			got := lo.Map(r.Search(tt.text),
				func(s ikku.Song, _ int) [][]string { return songToTexts(&s) },
			)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func FuzzTestJudgeAllRunes(f *testing.F) {
	f.Add("古池や蛙飛び込む水の音", true)
	f.Add("古池や蛙飛び込む水の音", false)
	f.Add("五月雨を集めてはやし最上川", true)
	f.Add("五月雨を集めてはやし最上川", false)

	f.Fuzz(func(t *testing.T, text string, exactly bool) {
		if utf8.RuneCountInString(text) > 24 {
			t.SkipNow()
		}
		r, err := ikku.NewReviewer(
			ipa.Dict(),
			ikku.ReviewerOptionExactly(exactly),
			ikku.ReviewerOptionRule([]int{2, 3, 2}),
		)
		if err != nil {
			t.Error(err)
		}
		r.Judge(text)
	})
}

func FuzzTestJudgeHiraganaKatakana(f *testing.F) {
	f.Add("古池や蛙飛び込む水の音", true)
	f.Add("古池や蛙飛び込む水の音", false)
	f.Add("五月雨を集めてはやし最上川", true)
	f.Add("五月雨を集めてはやし最上川", false)

	f.Fuzz(func(t *testing.T, text string, exactly bool) {
		if utf8.RuneCountInString(text) > 24 || !isTargetText(text) {
			t.SkipNow()
		}
		r, err := ikku.NewReviewer(
			ipa.Dict(),
			ikku.ReviewerOptionExactly(exactly),
			ikku.ReviewerOptionRule([]int{2, 3, 2}),
		)
		if err != nil {
			t.Error(err)
		}
		r.Judge(text)
	})
}

func songToTexts(s *ikku.Song) [][]string {
	if s == nil {
		return nil
	}
	return lo.Map(s.Phrases, func(ph []ikku.Node, _ int) []string {
		return lo.Map(ph, func(n ikku.Node, _ int) string { return n.String() })
	})
}

func isTargetText(text string) bool {
	for _, r := range text {
		if !(unicode.In(r, unicode.Hiragana) || unicode.In(r, unicode.Han)) {
			return false
		}
	}
	return true
}
