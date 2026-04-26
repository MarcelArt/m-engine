---
name: create-entity
description: Scaffold entity files for M-Engine (Go game engine). Creates Go structs with Start/Update/Destroy lifecycle methods. Use this skill whenever the user asks to create, add, make, or scaffold an entity, game object, component, actor, sprite, character, enemy, player, NPC, or any similar game object terminology in the context of Go game development or the m-engine codebase. This includes phrases like "create a player entity", "add an enemy", "make a coin object", "scaffold a bird entity", etc.
---

# Create Entity

Scaffolds a new entity file for the M-Engine game engine.

## When to use

Use this whenever the user wants to create a new game object/entity in their M-Engine project. The user might say:
- "create a Player entity"
- "add an Enemy with physics"
- "make a Coin entity that collects on collision"
- "scaffold a Bird entity for the flappy example"

## How it works

1. Parse the entity name from user input (convert to PascalCase)
2. Determine the target example (flappy, nakama, or create new)
3. Create the entity file with lifecycle methods
4. Add optional features based on user request

## File location

Entities go in: `examples/<example>/internal/entities/<name>.go`

## Basic entity template

```go
package entities

import (
    "github.com/MarcelArt/m-engine/pkg/engine"
    rl "github.com/gen2brain/raylib-go/raylib"
)

type EntityName struct {
    // Fields go here
}

func (e *EntityName) Start(g *engine.Game) {
    // Initialization
}

func (e *EntityName) Update(g *engine.Game) {
    // Per-frame logic
}

func (e *EntityName) Destroy(g *engine.Game) {
    // Cleanup
}

var _ engine.Entity = &EntityName{}
```

## Optional features

Add these based on user request:

### Physics
Add these methods and fields:
```go
type EntityName struct {
    Position rl.Vector2
    Velocity rl.Vector2
}

func (e *EntityName) GetPosition() rl.Vector2 { return e.Position }
func (e *EntityName) SetPosition(pos rl.Vector2) { e.Position = pos }
func (e *EntityName) GetVelocity() rl.Vector2 { return e.Velocity }
func (e *EntityName) SetVelocity(vel rl.Vector2) { e.Velocity = vel }
func (e *EntityName) IsGravityEnabled() bool { return true }

var _ engine.PhysicsObject = &EntityName{}
```

### Collision
Add these methods:
```go
func (e *EntityName) GetColliderRect() rl.Rectangle { /* return bounds */ }
func (e *EntityName) SetColliderRect(rect rl.Rectangle) { /* update bounds */ }
func (e *EntityName) OnCollision(g *engine.Game, other engine.RectCollidable) {}

var _ engine.RectCollidable = &EntityName{}
```

For collision enter/exit callbacks:
```go
func (e *EntityName) OnCollisionEnter(g *engine.Game, other engine.RectCollidable) {}
func (e *EntityName) OnCollisionExit(g *engine.Game, other engine.RectCollidable) {}

var _ engine.CollisionEnterHandler = &EntityName{}
var _ engine.CollisionExitHandler = &EntityName{}
```

### Sprite
Add field and example usage:
```go
type EntityName struct {
    Sprite *engine.Spritesheet
}

// In Update():
// e.Sprite.DrawTile(tileIndex, position, rl.White)
```

## Instructions

1. Extract entity name - convert to PascalCase (e.g., "player" → "Player", "flappy_bird" → "FlappyBird")
2. Detect target example from context or current working directory
3. Ask clarifying questions if the example is ambiguous
4. Determine which features to add (physics, collision, sprite, animation)
5. Create the file with proper Go formatting
6. Verify compilation with `go build`
7. Report the created file path

## Examples

| User input | Entity name | Features |
|------------|-------------|----------|
| "create a Player entity" | Player | basic |
| "add an Enemy with physics" | Enemy | physics |
| "make a Coin that collects on collision" | Coin | collision |
| "create a Bird entity with sprite animation" | Bird | sprite, animation |
