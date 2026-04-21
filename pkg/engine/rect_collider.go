package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type RectCollidable interface {
	OnCollision(other RectCollidable)
	GetColliderRect() rl.Rectangle
	SetColliderRect(rect rl.Rectangle)
}

type CollisionEnterHandler interface {
	OnCollisionEnter(other RectCollidable)
}

type CollisionExitHandler interface {
	OnCollisionExit(other RectCollidable)
}
