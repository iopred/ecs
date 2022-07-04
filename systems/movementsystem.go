package systems

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/iopred/ecs"
	"github.com/iopred/ecs/components"
)

const MovementSystemIdentifier ecs.Identifier = "MovementSystem"

type MovementSystem struct {
	*ecs.BasicSystem
	width  float64
	height float64
}

func (ms *MovementSystem) add(entity ecs.Entity) bool {
	return entity.HasAllComponents(components.PositionIdentifier, components.VelocityIdentifier)
}

func (ms *MovementSystem) update(screen *ebiten.Image) {
	entities := ms.Entities()
	for _, e := range entities {
		position := e.GetComponent(components.PositionIdentifier).(*components.Position)
		velocity := e.GetComponent(components.VelocityIdentifier).(*components.Velocity)

		w := 0.0
		h := 0.0
		size, ok := e.GetComponent(components.SizeIdentifier).(*components.Size)
		if ok {
			w = size.Width
			h = size.Height
		}

		timeScale, ok := e.GetComponent(components.TimeScaleIdentifier).(*components.TimeScale)
		scalar := 1.0
		if ok {
			scalar = timeScale.Scalar
		}

		position.X += velocity.X * scalar
		position.Y += velocity.Y * scalar

		if position.X+w > ms.width && velocity.X > 0 {
			velocity.X *= -0.9
			position.X -= (position.X + w) - ms.width
		}
		if position.X < 0 && velocity.X < 0 {
			velocity.X *= -0.9
			position.X -= position.X
		}
		if position.Y+h > ms.height && velocity.Y > 0 {
			velocity.X *= 0.9
			velocity.Y *= -0.9
			position.Y -= (position.Y + h) - ms.height
		}
		if position.Y < 0 && velocity.Y < 0 {
			velocity.X *= 0.9
			velocity.Y *= -0.9
			position.Y -= position.Y
		}

		if math.Abs(velocity.X) < 0.001 && math.Abs(velocity.Y) < 0.001 {
			velocity.X = 0
			velocity.Y = 0
			position.Y = ms.height - h
		}
	}
}

func NewMovementSystem(width, height float64) *MovementSystem {
	ms := &MovementSystem{
		width:  width,
		height: height,
	}
	ms.BasicSystem = ecs.NewBasicSystem(MovementSystemIdentifier, 0, ms.add, ms.update)
	return ms
}
