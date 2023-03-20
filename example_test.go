// nolint
package ikku_test

import (
	"fmt"

	"github.com/ikawaha/kagome-dict/ipa"
	ikku "github.com/kurochan/ikku-go"
)

func Example() {
	r, err := ikku.NewReviewer(ipa.Dict())
	if err != nil {
		panic(err)
	}
	// This is Haiku.
	fmt.Println(r.Find("古池や蛙飛び込む水の音"))
	// This contains Haiku.
	fmt.Println(r.Find("まさに古池や蛙飛び込む水の音ですね。"))
	// This is NOT Haiku.
	fmt.Println(r.Find("今日もいい天気だ。"))
}

func Example_print() {
	r, err := ikku.NewReviewer(ipa.Dict())
	if err != nil {
		panic(err)
	}
	song := r.Find("まさに古池や蛙飛び込む水の音ですね。")
	fmt.Println(song.String())
	// Output: 古池や蛙飛び込む水の音
}

func ExampleReviewerOptionExactly() {
	ikku.NewReviewer(
		ipa.Dict(),
		ikku.ReviewerOptionExactly(true),
	)
}

func ExampleReviewerOptionRule() {
	//nolint
	ikku.NewReviewer(
		ipa.Dict(),
		ikku.ReviewerOptionRule([]int{5, 7, 5, 7, 7}),
	)
}
