package ecs

import (
	"sort"

	"github.com/hajimehoshi/ebiten"
)

type System interface {
	Identifiable
	Prioritizable
	AddEntity(Entity) bool
	RemoveEntity(Identifiable) bool
	ReplaceEntity(Identifiable, Entity) bool
	Update(*ebiten.Image)
	Entities() Entities
}

type Systems []System

func (s Systems) Len() int           { return len(s) }
func (s Systems) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Systems) Less(i, j int) bool { return s[i].Priority() > s[j].Priority() }

type BasicSystem struct {
	id        Identifier
	priority  int
	add       func(entity Entity) bool
	update    func(*ebiten.Image)
	entities  Entities
	entityMap map[Identifier]Entity
	dirty     bool
}

func (s *BasicSystem) ID() Identifier {
	return s.id
}

func (s *BasicSystem) Priority() int {
	return s.priority
}

func (s *BasicSystem) Update(screen *ebiten.Image) {
	if s.dirty {
		sort.Sort(s.entities)
		s.dirty = false
	}
	s.update(screen)
}

func (s *BasicSystem) AddEntity(entity Entity) bool {
	if entity == nil {
		return false
	}

	if !s.add(entity) {
		return false
	}

	id := entity.ID()
	if _, ok := s.entityMap[id]; ok {
		return false
	}

	s.entityMap[id] = entity
	s.entities = append(s.entities, entity)
	s.dirty = true
	return true
}

func (s *BasicSystem) RemoveEntity(identifiable Identifiable) bool {
	id := identifiable.ID()
	if _, ok := s.entityMap[id]; !ok {
		return false
	}

	delete(s.entityMap, id)
	for i, e := range s.entities {
		if e.ID() == id {
			s.entities = append(s.entities[:i], s.entities[i+1:]...)
			break
		}
	}
	return true
}

func (s *BasicSystem) GetEntity(identifiable Identifiable) Entity {
	return s.entityMap[identifiable.ID()]
}

func (s *BasicSystem) ReplaceEntity(identifiable Identifiable, entity Entity) bool {
	id := identifiable.ID()
	if _, ok := s.entityMap[id]; !ok {
		return false
	}

	s.entityMap[id] = entity
	for i, e := range s.entities {
		if e.ID() == id {
			s.entities[i] = entity
			break
		}
	}
	return true
}

func (s *BasicSystem) Entities() Entities {
	return s.entities
}

func NewBasicSystem(id Identifier, priority int, add func(Entity) bool, update func(*ebiten.Image)) *BasicSystem {
	return &BasicSystem{
		id:        id,
		priority:  priority,
		add:       add,
		update:    update,
		entityMap: map[Identifier]Entity{},
		dirty:     false,
	}
}
