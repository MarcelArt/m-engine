package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type CollisionSystem struct {
	RectColliders []RectCollidable
	IsDebug       bool
}

func NewCollisionSystem(isDebug bool) *CollisionSystem {
	return &CollisionSystem{
		IsDebug:       isDebug,
		RectColliders: make([]RectCollidable, 0),
	}
}

func (c *CollisionSystem) Update() {
	c.checkCollisionLoop()
	c.drawDebug()
}

func (c *CollisionSystem) AddRectCollidable(collider RectCollidable) {
	c.RectColliders = append(c.RectColliders, collider)
}

func (c *CollisionSystem) checkCollisionLoop() {
	for _, collider := range c.RectColliders {
		for _, other := range c.RectColliders {
			if collider != other {
				if rl.CheckCollisionRecs(collider.GetColliderRect(), other.GetColliderRect()) {
					collider.OnCollision(other)
					other.OnCollision(collider)
				}
			}
		}
	}
}

func (c *CollisionSystem) drawDebug() {
	if c.IsDebug {
		for _, collider := range c.RectColliders {
			rl.DrawRectangleLines(int32(collider.GetColliderRect().X), int32(collider.GetColliderRect().Y), int32(collider.GetColliderRect().Width), int32(collider.GetColliderRect().Height), rl.Red)
		}
	}
}
