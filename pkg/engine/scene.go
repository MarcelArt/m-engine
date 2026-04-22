package engine

type IScene interface {
	Entity
	AddEntity(e Entity)
	GetEntities() []Entity
	AddUIEntity(e Entity)
	GetUIEntities() []Entity
	AddBGEntity(e Entity)
	GetBGEntities() []Entity
	ClearEntities()
}

type Scene struct {
	entities   []Entity
	uiEntities []Entity
	bgEntities []Entity
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

func (s *Scene) AddUIEntity(e Entity) {
	s.uiEntities = append(s.uiEntities, e)
}

func (s *Scene) GetUIEntities() []Entity {
	return s.uiEntities
}

func (s *Scene) AddBGEntity(e Entity) {
	s.bgEntities = append(s.bgEntities, e)
}

func (s *Scene) GetBGEntities() []Entity {
	return s.bgEntities
}

func (s *Scene) ClearEntities() {
	s.entities = nil
	s.uiEntities = nil
}
