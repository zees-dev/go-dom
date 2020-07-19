package models

import "testing"

func TestNode(t *testing.T) {

	t.Run("deep equality success", func(t *testing.T) {
		image := Node{
			tag: "img",
			src: "http://example.com/logo.svg",
			alt: "Example's Logo",
		}
		p := Node{
			tag:      "p",
			text:     "And this is some text in a paragraph. And next to it there's an image.",
			children: []*Node{&image},
		}

		span := Node{
			tag:  "span",
			id:   "copyright",
			text: "2019 &copy; Ilija Eftimov",
		}
		div := Node{
			tag:      "div",
			class:    "footer",
			text:     "This is the footer of the page.",
			children: []*Node{&span},
		}

		h1 := Node{
			tag:  "h1",
			text: "This is a H1",
		}

		body := Node{
			tag:      "body",
			children: []*Node{&h1, &p, &div},
		}

		got := Node{
			tag:      "html",
			children: []*Node{&body},
		}

		want := Node{
			tag: "html",
			children: []*Node{
				&Node{
					tag: "body",
					children: []*Node{
						&h1,
						&Node{
							tag:  "p",
							text: "And this is some text in a paragraph. And next to it there's an image.",
							children: []*Node{
								&Node{
									tag: "img",
									src: "http://example.com/logo.svg",
									alt: "Example's Logo",
								},
							},
						},
						&div,
					},
				},
			},
		}

		if !got.Equals(&want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("deep equality failure", func(t *testing.T) {
		image := Node{
			tag: "img",
			src: "http://example.com/logo.svg",
			alt: "Example's Logo",
		}

		p := Node{
			tag:      "p",
			text:     "And this is some text in a paragraph. And next to it there's an image.",
			children: []*Node{&image},
		}

		span := Node{
			tag:  "span",
			id:   "copyright",
			text: "2019 &copy; Ilija Eftimov",
		}

		div := Node{
			tag:      "div",
			class:    "footer",
			text:     "This is the footer of the page.",
			children: []*Node{&span},
		}

		h1 := Node{
			tag:  "h1",
			text: "This is a H1",
		}

		body := Node{
			tag:      "body",
			children: []*Node{&h1, &p, &div},
		}

		got := Node{
			tag:      "html",
			children: []*Node{&body},
		}

		want := Node{
			tag: "html",
			children: []*Node{
				&Node{
					tag: "body",
					children: []*Node{
						&h1,
						&Node{
							tag:  "p",
							text: "And this is some text in a paragraph. And next to it there's an image.",
							children: []*Node{
								&Node{
									tag: "img",
									src: "http://example.com/logo.svg",
									alt: "MISMATCH HERE",
								},
							},
						},
						&div,
					},
				},
			},
		}

		if got.Equals(&want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
