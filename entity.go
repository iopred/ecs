package ecs

import (
	"strconv"
	"sync/atomic"
)

type Entity interface {
	Identifiable
	Prioritizable
	AddComponent(Component) bool
	RemoveComponent(Identifier) bool
	GetComponent(Identifier) Component
	ReplaceComponent(Identifier, Component) bool
	Components() map[Identifier]Component
	HasComponent(Identifier) bool
	HasAllComponents(...Identifier) bool
}

type Entities []Entity

func (e Entities) Len() int           { return len(e) }
func (e Entities) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Entities) Less(i, j int) bool { return e[i].Priority() > e[j].Priority() }

type BasicEntity struct {
	id         Identifier
	priority   int
	components map[Identifier]Component
}

func (e *BasicEntity) ID() Identifier {
	return e.id
}

func (e *BasicEntity) Priority() int {
	return e.priority
}

func (e *BasicEntity) AddComponent(component Component) bool {
	if component == nil {
		return false
	}

	id := component.ID()
	if _, ok := e.components[id]; ok {
		return false
	}

	e.components[id] = component
	return true
}

func (e *BasicEntity) RemoveComponent(id Identifier) bool {
	if _, ok := e.components[id]; !ok {
		return false
	}

	delete(e.components, id)
	return true
}

func (e *BasicEntity) GetComponent(id Identifier) Component {
	return e.components[id]
}

func (e *BasicEntity) ReplaceComponent(id Identifier, component Component) bool {
	if _, ok := e.components[id]; !ok {
		return false
	}

	e.components[id] = component
	return true
}

func (e *BasicEntity) Components() map[Identifier]Component {
	return e.components
}

func (e *BasicEntity) HasComponent(id Identifier) bool {
	_, ok := e.components[id]
	return ok
}

func (e *BasicEntity) HasAllComponents(identifiers ...Identifier) bool {
	for _, identifier := range identifiers {
		if !e.HasComponent(identifier) {
			return false
		}
	}
	return true
}

var id int64

func NewBasicEntity() *BasicEntity {
	atomic.AddInt64(&id, 1)
	return &BasicEntity{
		id:         Identifier(strconv.Itoa(int(id))),
		components: map[Identifier]Component{},
	}
}
