package scribe

import (
	"io"

	"github.com/spcameron/scribe/internal/codegen"
	"github.com/spcameron/scribe/internal/html"
	"github.com/spcameron/scribe/internal/lower"
	"github.com/spcameron/scribe/internal/source"
)

type Document interface {
	Write(io.Writer) error
}

// Compile parses Markdown and returns a renderable document.
func Compile(md string) (Document, error) {
	src := source.NewSource(md)

	irDoc, err := block.Parse(src)
	if err != nil {
		return nil, err
	}

	astDoc, err := lower.Document(irDoc)
	if err != nil {
		return nil, err
	}

	tree, err := codegen.HTML(astDoc)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

// HTML parses Markdown and renders the result as an HTML string.
func HTML(md string) (string, error) {
	tree, err := Compile(md)
	if err != nil {
		return "", err
	}

	htmlStr, err := html.Render(tree)
	if err != nil {
		return "", err
	}

	return htmlStr, nil
}
