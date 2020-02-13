package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/russross/blackfriday"
)

type StackNode struct {
	NodeType blackfriday.NodeType
	Text     string
}

func main() {
	flag.Parse()
	args := flag.Args()

	md := strings.ReplaceAll(args[0], "\\r", "\n")
	md = strings.ReplaceAll(md, "\\n", "\n\n")
	markdown := blackfriday.New()
	rootNode := markdown.Parse([]byte(md))

	stackNodes := make([]*StackNode, 0)
	res := ""
	rootNode.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if entering {
			stackNodes = append(stackNodes, &StackNode{
				NodeType: node.Type,
				Text:     string(node.Literal),
			})
		} else {
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
					switch currentType {
					case blackfriday.Heading:
						stackNodes = append(stackNodes, &StackNode{
							NodeType: blackfriday.Text,
							Text:     "### " + text + "\n",
						})
					case blackfriday.Strong:
						stackNodes = append(stackNodes, &StackNode{
							NodeType: blackfriday.Text,
							Text:     "**" + text + "**",
						})
					case blackfriday.Item:
						stackNodes = append(stackNodes, &StackNode{
							NodeType: blackfriday.Text,
							Text:     "- [ ] " + text + "\n",
						})
					default:
						stackNodes = append(stackNodes, &StackNode{
							NodeType: blackfriday.Text,
							Text:     text,
						})
					}
					popDone = false
				}

				text = popNode.Text + text
			}
		}

		return blackfriday.GoToNext
	})

	fmt.Println(res)
}
