package components

import "github.com/iopred/ecs"

const SizeIdentifier ecs.Identifier = "Size"

type Size struct {
	Width  float64
	Height float64
}

func (p *Size) ID() ecs.Identifier {
	return SizeIdentifier
}
