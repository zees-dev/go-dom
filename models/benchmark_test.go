package models

import (
	"testing"
)

func generateExampleDom() *Node {
	exampleNodeToReplicate := &Node{
		tag:  "p",
		text: "And this is some text in a paragraph. And next to it there's an image.",
		children: []*Node{
			{
				tag: "img",
				src: "http://example.com/logo.svg",
				alt: "Example's Logo",
			},
		},
	}

	manyNodes := make([]*Node, 1000)
	for i := range manyNodes {
		manyNodes[i] = exampleNodeToReplicate
	}

	html := Node{
		tag: "html",
		children: []*Node{
			{
				tag: "body",
				children: []*Node{
					{
						tag:      "div",
						text:     "And this is some text in a paragraph. And next to it there's an image.",
						children: manyNodes,
					},
					{
						tag:  "h1",
						text: "This is a H1",
					},
					{
						tag:   "div",
						class: "footer",
						text:  "This is the footer of the page.",
						children: []*Node{
							{
								tag:  "span",
								text: "2019 &copy; Ilija Eftimov",
							},
							{
								tag:  "p",
								text: "And this is some text in a paragraph. And next to it there's an image.",
								children: []*Node{
									{
										tag: "img",
										src: "http://example.com/logo.svg",
										alt: "Example's Logo",
									},
								},
							},
							{
								tag:  "h2",
								text: "This is a H2",
								id:   "id-to-find",
							},
						},
					},
				},
			},
		},
	}

	return &html
}

func BenchmarkGetElementByID(b *testing.B) {
	dom := generateExampleDom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dom.GetElementByID("id-to-find")
	}
}

func BenchmarkDepthFirstSearch(b *testing.B) {
	dom := generateExampleDom()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dom.GetElementByIDBFS("id-to-find")
	}
}
