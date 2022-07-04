package systems

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/iopred/ecs"
	"github.com/iopred/ecs/components"
)

const GravitySystemIdentifier ecs.Identifier = "GravitySystem"

type GravitySystem struct {
	*ecs.BasicSystem
	width  float64
	height float64
}

func (gs *GravitySystem) add(entity ecs.Entity) bool {
	return entity.HasAllComponents(components.VelocityIdentifier)
}

func (gs *GravitySystem) update(screen *ebiten.Image) {
	entities := gs.Entities()
	for _, e := range entities {
		velocity := e.GetComponent(components.VelocityIdentifier).(*components.Velocity)
		timeScale, ok := e.GetComponent(components.TimeScaleIdentifier).(*components.TimeScale)
		scalar := 1.0
		if ok {
			scalar = timeScale.Scalar
		}

		velocity.Y += 1 * scalar
	}
}

func NewGravitySystem() *GravitySystem {
	gs := &GravitySystem{}
	gs.BasicSystem = ecs.NewBasicSystem(GravitySystemIdentifier, 0, gs.add, gs.update)
	return gs
}
