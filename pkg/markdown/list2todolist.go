package markdown

import (
	"strings"

	"github.com/russross/blackfriday"
)

type StackNode struct {
	NodeType blackfriday.NodeType
	Text     string
}

func List2TodoList(md string) string {
	markdown := blackfriday.New()
	rootNode := markdown.Parse([]byte(stripLF(md)))

	stackNodes := make([]*StackNode, 0)
	res := ""
	rootNode.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if entering {
			stackNodes = append(stackNodes, &StackNode{
				NodeType: node.Type,
				Text:     string(node.Literal),
			})
			return blackfriday.GoToNext
		}

		if node.Type == blackfriday.Document {
			for _, n := range stackNodes {
				res = res + n.Text
			}
			return blackfriday.SkipChildren
		}

		popDone := true
		currentType := node.Type
		text := ""
		for popDone {
			popNode, tmp := stackNodes[len(stackNodes)-1], stackNodes[:len(stackNodes)-1]
			stackNodes = tmp
			if popNode.NodeType == currentType {
				stackNodes = append(stackNodes, NewStackNode(currentType, text))
				popDone = false
			}

			text = popNode.Text + text
		}

		return blackfriday.GoToNext
	})

	return res
}

func NewStackNode(nodeType blackfriday.NodeType, text string) *StackNode {
	var res *StackNode
	switch nodeType {
	case blackfriday.Heading:
		res = &StackNode{
			NodeType: blackfriday.Text,
			Text:     "### " + text + "\n",
		}
	case blackfriday.Strong:
		res = &StackNode{
			NodeType: blackfriday.Text,
			Text:     "**" + text + "**",
		}
	case blackfriday.Item:
		res = &StackNode{
			NodeType: blackfriday.Text,
			Text:     "- [ ] " + text + "\n",
		}
	case blackfriday.Paragraph:
		res = &StackNode{
			NodeType: blackfriday.Text,
			Text:     text,
		}
	default:
		res = &StackNode{
			NodeType: blackfriday.Text,
			Text:     text,
		}
	}

	return res
}

func stripLF(md string) string {
	res := strings.ReplaceAll(md, "\\r", "\n")
	res = strings.ReplaceAll(res, "\r", "\n")
	res = strings.ReplaceAll(res, "\\n", "\n\n")

	return res
}
