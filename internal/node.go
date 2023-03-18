package internal

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/samber/lo"
)

type Node struct {
	Token    tokenizer.Token
	features []string
}

func NewNode(token tokenizer.Token) *Node {
	features := token.Features()
	return &Node{
		Token:    token,
		features: features,
	}
}

func (n *Node) FirstOfIkku() bool {
	switch {
	case !n.FirstOfPhrase():
		return false
	case n._type() == "記号" && !lo.Contains([]string{"括弧開", "括弧閉"}, n.subtype1()):
		return false
	default:
		return true
	}
}

func (n *Node) PronunciationLength() int {
	if n.pronunciation() == "" {
		return 0
	}
	mora := n.pronunciation()
	for i := 0; i <= 'う'-'ぁ'; i++ {
		old := string(rune('ぁ' + i))
		new := string(rune('ァ' + i))
		mora = strings.ReplaceAll(mora, old, new)
	}
	re, _ := regexp.Compile(`[^アイウエオカ-モヤユヨラ-ロワヲンヴー]`)
	mora = re.ReplaceAllString(mora, "")
	return utf8.RuneCountInString(mora)
}

func (n *Node) ElementOfIkku() bool {
	return lo.Contains([]tokenizer.TokenClass{tokenizer.KNOWN, tokenizer.USER}, n.Token.Class)
}

func (n *Node) FirstOfPhrase() bool {
	switch {
	case lo.Contains([]string{"助詞", "助動詞"}, n._type()):
		return false
	case lo.Contains([]string{"非自立", "接尾"}, n.subtype1()):
		return false
	case n.subtype1() == "自立" && lo.Contains([]string{"する", "できる"}, n.Token.Surface):
		return false
	default:
		return true
	}
}

func (n *Node) LastOfPhrase() bool {
	return n._type() != "接頭詞"
}

func (n *Node) LastOfIkku() bool {
	switch {
	case lo.Contains([]string{"名詞接続", "格助詞", "係助詞", "連体化", "接続助詞", "並立助詞", "副詞化", "数接続", "連体詞"}, n._type()):
		return false
	case n.confugation2() == "連用タ接続":
		return false
	case n.confugation1() == "サ変・スル" && n.confugation2() == "連用形":
		return false
	case n._type() == "動詞" && lo.Contains([]string{"仮定形", "未然形"}, n.confugation2()):
		return false
	case n._type() == "名詞" && n.subtype1() == "非自立" && n.pronunciation() == "ン":
		return false
	default:
		return true
	}
}

func (n *Node) _type() string {
	return n.features[0]
}

func (n *Node) subtype1() string {
	return n.features[1]
}

func (n *Node) confugation1() string {
	i, _ := n.Token.InflectionalForm()
	return i
}

func (n *Node) confugation2() string {
	i, _ := n.Token.InflectionalType()
	return i
}

func (n *Node) pronunciation() string {
	p, _ := n.Token.Pronunciation()
	return p
}
