package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type collisionPair struct {
	entityA RectCollidable
	entityB RectCollidable
}

type CollisionSystem struct {
	RectColliders     []RectCollidable
	activeCollisions  map[collisionPair]struct{}
	IsDebug           bool
}

func NewCollisionSystem(isDebug bool) *CollisionSystem {
	return &CollisionSystem{
		IsDebug:       isDebug,
		RectColliders: make([]RectCollidable, 0),
		activeCollisions: make(map[collisionPair]struct{}),
	}
}

func (c *CollisionSystem) Update() {
	c.checkCollisionLoop()
	c.drawDebug()
}

func (c *CollisionSystem) AddRectCollidable(collider RectCollidable) {
	c.RectColliders = append(c.RectColliders, collider)
}

func (c *CollisionSystem) RemoveRectCollidable(entity RectCollidable) {
	// Fire exit callbacks for all active collisions involving this entity
	for pair := range c.activeCollisions {
		if pair.entityA == entity || pair.entityB == entity {
			if handler, ok := pair.entityA.(CollisionExitHandler); ok && pair.entityB == entity {
				handler.OnCollisionExit(pair.entityB)
			}
			if handler, ok := pair.entityB.(CollisionExitHandler); ok && pair.entityA == entity {
				handler.OnCollisionExit(pair.entityA)
			}
		}
	}

	// Remove from the slice
	for i, e := range c.RectColliders {
		if e == entity {
			c.RectColliders = append(c.RectColliders[:i], c.RectColliders[i+1:]...)
			break
		}
	}
}

func (c *CollisionSystem) checkCollisionLoop() {
	currentCollisions := make(map[collisionPair]struct{})

	// Build current frame collisions and fire enter callbacks
	for i := 0; i < len(c.RectColliders); i++ {
		for j := i + 1; j < len(c.RectColliders); j++ {
			collider := c.RectColliders[i]
			other := c.RectColliders[j]

			if rl.CheckCollisionRecs(collider.GetColliderRect(), other.GetColliderRect()) {
				pair := collisionPair{entityA: collider, entityB: other}
				currentCollisions[pair] = struct{}{}

				// OnEnter: new collision
				if _, wasActive := c.activeCollisions[pair]; !wasActive {
					if handler, ok := collider.(CollisionEnterHandler); ok {
						handler.OnCollisionEnter(other)
					}
					if handler, ok := other.(CollisionEnterHandler); ok {
						handler.OnCollisionEnter(collider)
					}
				}

				// OnCollision: continuous callback
				collider.OnCollision(other)
				other.OnCollision(collider)
			}
		}
	}

	// OnExit: ended collisions
	for pair := range c.activeCollisions {
		if _, stillActive := currentCollisions[pair]; !stillActive {
			if handler, ok := pair.entityA.(CollisionExitHandler); ok {
				handler.OnCollisionExit(pair.entityB)
			}
			if handler, ok := pair.entityB.(CollisionExitHandler); ok {
				handler.OnCollisionExit(pair.entityA)
			}
		}
	}

	c.activeCollisions = currentCollisions
}

func (c *CollisionSystem) drawDebug() {
	if c.IsDebug {
		for _, collider := range c.RectColliders {
			rect := collider.GetColliderRect()
			rl.DrawRectangleLines(int32(rect.X), int32(rect.Y), int32(rect.Width), int32(rect.Height), rl.Red)
		}
	}
}
