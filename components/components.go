package components

import (
	"io"
	"sort"
	"strings"

	g "github.com/sempaxjeen/gohtml"
	h "github.com/sempaxjeen/gohtml/html"
)

type HTML5Props struct {
	Title       string
	Description string
	Language    string
	Head        []g.Node
	Body        []g.Node
	HTMLAttrs   []g.Node
}

func HTML5(p HTML5Props) g.Node {
	return h.Doctype(
		h.HTML(g.If(p.Language != "", h.Lang(p.Language)), g.Group(p.HTMLAttrs),
			h.Head(
				h.Meta(h.Charset("utf-8")),
				h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
				h.TitleEl(g.Text(p.Title)),
				g.If(p.Description != "", h.Meta(h.Name("description"), h.Content(p.Description))),
				g.Group(p.Head),
			),
			h.Body(g.Group(p.Body)),
		),
	)
}

type Classes map[string]bool

func (c Classes) Render(w io.Writer) error {
	included := make([]string, 0, len(c))
	for c, include := range c {
		if include {
			included = append(included, c)
		}
	}
	sort.Strings(included)
	return h.Class(strings.Join(included, " ")).Render(w)
}

func (c Classes) Type() g.NodeType {
	return g.AttributeType
}

func (c Classes) String() string {
	var b strings.Builder
	_ = c.Render(&b)
	return b.String()
}
