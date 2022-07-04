package ecs

import (
	"sort"

	"github.com/hajimehoshi/ebiten"
)

type ECS struct {
	systems Systems
}

func (e *ECS) AddSystem(system System) {
	if system == nil {
		return
	}

	e.systems = append(e.systems, system)
	sort.Sort(e.systems)
}

func (e *ECS) GetSystem(id Identifier) System {
	for _, s := range e.systems {
		if s.ID() == id {
			return s
		}
	}
	return nil
}

func (e *ECS) AddEntity(entity Entity) {
	if entity == nil {
		return
	}

	for _, s := range e.systems {
		s.AddEntity(entity)
	}
}

func (e *ECS) RemoveEntity(entity Entity) {
	for _, s := range e.systems {
		s.RemoveEntity(entity)
	}
}

func (e *ECS) Update(screen *ebiten.Image) {
	for _, s := range e.systems {
		s.Update(screen)
	}
}
