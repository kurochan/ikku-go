package ikku_test

import (
	"testing"

	"github.com/ikawaha/kagome-dict/ipa"
	ikku "github.com/kurochan/ikku-go"
)

func TestSongString(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "ikku/古池や蛙飛び込む水の音",
			text: "古池や蛙飛び込む水の音",
			want: "古池や蛙飛び込む水の音",
		},
		{
			name: "ikku/ああ古池や蛙飛び込む水の音がします",
			text: "ああ古池や蛙飛び込む水の音がします",
			want: "古池や蛙飛び込む水の音",
		},
	}
	for _, tt := range tests {
		r, err := ikku.NewReviewer(ipa.Dict())
		if err != nil {
			t.Error(err)
		}
		t.Run(tt.name, func(t *testing.T) {
			got := r.Find(tt.text).String()
			if tt.want != got {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
