package components

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/iopred/ecs"
)

const RenderableIdentifier ecs.Identifier = "Renderable"

type Renderable struct {
	Texture *ebiten.Image
}

func (p *Renderable) ID() ecs.Identifier {
	return RenderableIdentifier
}
