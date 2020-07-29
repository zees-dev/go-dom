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

func BenchmarkGetElementByIDViaWaitGroup(b *testing.B) {
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := domNormal.GetElementByIDViaWaitGroup("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domBFSOptimized.GetElementByIDViaWaitGroup("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domDFSOptimized.GetElementByIDViaWaitGroup("id-to-find")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetElementByIDViaGoRoutines(b *testing.B) {
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := domNormal.GetElementByIDViaGoRoutines("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domBFSOptimized.GetElementByIDViaGoRoutines("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domDFSOptimized.GetElementByIDViaGoRoutines("id-to-find")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkBreadthFirstSearch(b *testing.B) {
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := domNormal.GetElementByIDBFS("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domBFSOptimized.GetElementByIDBFS("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domDFSOptimized.GetElementByIDBFS("id-to-find")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDepthFirstSearch(b *testing.B) {
	domNormal := generateExampleDom()
	domBFSOptimized := generateExampleDomBFSOptimized()
	domDFSOptimized := generateExampleDomDFSOptimized()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := domNormal.GetElementByIDDFS("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domBFSOptimized.GetElementByIDDFS("id-to-find")
		if err != nil {
			b.Error(err)
		}

		_, err = domDFSOptimized.GetElementByIDDFS("id-to-find")
		if err != nil {
			b.Error(err)
		}
	}
}
