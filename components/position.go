package components

import "github.com/iopred/ecs"

const PositionIdentifier ecs.Identifier = "Position"

type Position struct {
	X float64
	Y float64
}

func (p *Position) ID() ecs.Identifier {
	return PositionIdentifier
}
