package ikku_test

import (
	"testing"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	ikku "github.com/kurochan/ikku-go"
)

func TestNodeString(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "天気",
			text: "天気",
			want: "天気",
		},
		{
			name: "明日も晴れるかな",
			text: "明日も晴れるかな",
			want: "明日",
		},
	}
	for _, tt := range tests {
		tn, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
		if err != nil {
			t.Error(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			tokens := tn.Tokenize(tt.text)
			if len(tokens) == 0 {
				t.Error("failed to tokenize")
			}
			node := ikku.Node{Token: tokens[0]}
			got := node.String()
			if tt.want != got {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
