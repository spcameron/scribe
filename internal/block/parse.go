package block

import (
	"github.com/spcameron/scribe/internal/ir"
	"github.com/spcameron/scribe/internal/source"
)

// Parse scans source lines and builds the block-level IR document.
func Parse(src *source.Source) (ir.Document, error) {
	lines, err := Scan(src)
	if err != nil {
		return ir.Document{}, err
	}

	out, err := Build(src, lines)
	if err != nil {
		return ir.Document{}, err
	}

	return out, nil
}
