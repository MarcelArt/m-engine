package engine

type IScene interface {
	Entity
	AddEntity(e Entity)
	GetEntities() []Entity
	ClearEntities()
}

type Scene struct {
	entities []Entity
}

func (s *Scene) AddEntity(e Entity) {
	if e == nil {
		s.entities = make([]Entity, 0)
	}

	s.entities = append(s.entities, e)
}

func (s *Scene) GetEntities() []Entity {
	return s.entities
}

func (s *Scene) ClearEntities() {
	s.entities = nil
}
