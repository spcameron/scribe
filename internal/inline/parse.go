package inline

import (
	"github.com/spcameron/scribe/internal/ast"
	"github.com/spcameron/scribe/internal/ir"
	"github.com/spcameron/scribe/internal/source"
)

// Parse scans and builds inline content within span into AST inline nodes.
func Parse(src *source.Source, defs map[string]ir.ReferenceDefinition, span source.ByteSpan) ([]ast.Inline, error) {
	tokens, err := Scan(src, span)
	if err != nil {
		return nil, err
	}

	out, err := Build(src, defs, span, tokens)
	if err != nil {
		return nil, err
	}

	return out, nil
}
