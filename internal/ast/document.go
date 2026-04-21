package ast

import "github.com/spcameron/scribe/internal/source"

type Document struct {
	Source *source.Source
	Blocks []Block
}
