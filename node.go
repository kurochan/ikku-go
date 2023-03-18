package ikku

import (
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/kurochan/ikku-go/internal"
)

type Node struct {
	Token tokenizer.Token
}

func (n *Node) String() string {
	return n.Token.Surface
}

func internalNodeToNode(in *internal.Node) *Node {
	if in == nil {
		return nil
	}
	return &Node{
		Token: in.Token,
	}
}
