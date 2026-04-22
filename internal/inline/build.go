package inline

import (
	"github.com/spcameron/scribe/internal/ast"
	"github.com/spcameron/scribe/internal/ir"
	"github.com/spcameron/scribe/internal/source"
)

// Build resolves a tokenized inline span into AST inline nodes.
func Build(src *source.Source, defs map[string]ir.ReferenceDefinition, span source.ByteSpan, tokens []Token) ([]ast.Inline, error) {
	c := NewCursor(src, defs, span, tokens)

	inlines, err := c.Build()
	if err != nil {
		return []ast.Inline{}, err
	}

	return inlines, nil
}
