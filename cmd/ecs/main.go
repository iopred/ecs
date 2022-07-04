package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images/platformer"
	"github.com/iopred/ecs"
	"github.com/iopred/ecs/components"
	"github.com/iopred/ecs/systems"
)

const (
	numSprites = 100
)

var (
	game           = &ecs.ECS{}
	rendererSystem = systems.NewRendererSystem()
)

func main() {
	screenWidth := 1920.0 / 2.0
	screenHeight := 1080.0 / 2.0

	game.AddSystem(systems.NewGravitySystem())
	game.AddSystem(systems.NewMovementSystem(screenWidth, screenHeight))
	rendererSystem.Op.ColorM.Scale(1, 1, 1, 0.25)
	game.AddSystem(rendererSystem)

	rand.Seed(time.Now().UnixNano())

	img, _, err := image.Decode(bytes.NewReader(platformer.MainChar_png))
	if err != nil {
		log.Fatal(err)
	}

	ebitenImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := ebitenImage.Size()
	size := &components.Size{float64(w), float64(h)}
	renderable := &components.Renderable{ebitenImage}
	timeScale := &components.TimeScale{0.05}
	for i := 0; i < numSprites; i++ {
		entity := ecs.NewBasicEntity()
		entity.AddComponent(&components.Position{rand.Float64() * (screenWidth - size.Width), rand.Float64() * (screenHeight - size.Height)})
		entity.AddComponent(size)
		entity.AddComponent(&components.Velocity{-10.0 + rand.Float64()*10.0, -10.0 + rand.Float64()*10.0})
		entity.AddComponent(renderable)
		entity.AddComponent(timeScale)
		game.AddEntity(entity)
	}

	ebiten.SetMaxTPS(120)
	if err := ebiten.Run(update, int(screenWidth), int(screenHeight), 1, "ECS"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	game.Update(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d sprites, %0.2f FPS", numSprites, ebiten.CurrentTPS()))
	return nil
}
