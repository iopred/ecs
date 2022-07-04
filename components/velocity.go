package components

import "github.com/iopred/ecs"

const VelocityIdentifier ecs.Identifier = "Velocity"

type Velocity struct {
	X float64
	Y float64
}

func (p *Velocity) ID() ecs.Identifier {
	return VelocityIdentifier
}
