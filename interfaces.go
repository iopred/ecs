package ecs

type Identifier string

type Identifiable interface {
	ID() Identifier
}

type Prioritizable interface {
	Priority() int
}

type Component interface {
	Identifiable
}
