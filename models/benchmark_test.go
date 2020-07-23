package models

import (
	"testing"
)

func generateExampleDomBFSOptimized() *Node {
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
					{
						tag:      "h1",
						text:     "This is a H1",
						children: manyNodes,
					},
					{
						tag:      "div",
						text:     "And this is some text in a paragraph. And next to it there's an image.",
						children: manyNodes,
					},
				},
			},
		},
	}

	return &html
}

func generateExampleDomDFSOptimized() *Node {
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
						tag:      "h1",
						text:     "This is a H1",
						children: manyNodes,
					},
					{
						tag:      "div",
						text:     "And this is some text in a paragraph. And next to it there's an image.",
						children: manyNodes,
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
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		domNormal.GetElementByID("id-to-find")
		domBFSOptimized.GetElementByID("id-to-find")
		domDFSOptimized.GetElementByID("id-to-find")
	}
}

func BenchmarkBreadthFirstSearch(b *testing.B) {
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		domNormal.GetElementByIDBFS("id-to-find")
		domBFSOptimized.GetElementByIDBFS("id-to-find")
		domDFSOptimized.GetElementByIDBFS("id-to-find")
	}
}

func BenchmarkDepthFirstSearch(b *testing.B) {
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		domNormal.GetElementByIDDFS("id-to-find")
		domBFSOptimized.GetElementByIDDFS("id-to-find")
		domDFSOptimized.GetElementByIDDFS("id-to-find")
	}
}
