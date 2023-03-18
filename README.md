# ikku-go (一句)
`ikku-go` is Ikku("一句") detector, Ikku is something like Japanese Haiku("俳句").

Inspired by [r7kamura/ikku](https://github.com/r7kamura/ikku).

## Example
```go
import (
	"fmt"
	"github.com/ikawaha/kagome-dict/ipa"
	ikku "github.com/kurochan/ikku-go"
)

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

song := r.Find("まさに古池や蛙飛び込む水の音ですね。")
fmt.Println(song.String())
// Output: 古池や蛙飛び込む水の音
```
