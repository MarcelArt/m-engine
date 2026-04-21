package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type RectCollidable interface {
	OnCollision(g *Game, other RectCollidable)
	GetColliderRect() rl.Rectangle
	SetColliderRect(rect rl.Rectangle)
}

type CollisionEnterHandler interface {
	OnCollisionEnter(g *Game, other RectCollidable)
}

type CollisionExitHandler interface {
	OnCollisionExit(g *Game, other RectCollidable)
}
