package components

import (
	"github.com/iopred/ecs"
)

const TimeScaleIdentifier ecs.Identifier = "TimeScale"

type TimeScale struct {
	Scalar float64
}

func (p *TimeScale) ID() ecs.Identifier {
	return TimeScaleIdentifier
}
