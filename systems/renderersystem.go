package systems

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/iopred/ecs"
	"github.com/iopred/ecs/components"
)

const RendererSystemIdentifier ecs.Identifier = "RendererSystem"

type RendererSystem struct {
	*ecs.BasicSystem
	Op *ebiten.DrawImageOptions
}

func (rs *RendererSystem) add(entity ecs.Entity) bool {
	return entity.HasAllComponents(components.PositionIdentifier, components.RenderableIdentifier)
}

func (rs *RendererSystem) update(screen *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}

	entities := rs.Entities()
	for _, e := range entities {
		position := e.GetComponent(components.PositionIdentifier).(*components.Position)
		renderable, ok := e.GetComponent(components.RenderableIdentifier).(*components.Renderable)
		if ok && renderable.Texture != nil {
			rs.Op.GeoM.Reset()
			rs.Op.GeoM.Translate(position.X, position.Y)
			screen.DrawImage(renderable.Texture, rs.Op)
		}
	}
}

func NewRendererSystem() *RendererSystem {
	rs := &RendererSystem{
		Op: &ebiten.DrawImageOptions{},
	}

	rs.BasicSystem = ecs.NewBasicSystem(RendererSystemIdentifier, 0, rs.add, rs.update)
	return rs
}
