# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

M-Engine is a lightweight 2D game engine written in Go, built on top of [raylib-go](https://github.com/gen2brain/raylib-go). It provides an entity-component architecture with scene management, physics, collision detection, and save systems.

## Common Commands

### Running Examples
```bash
# Run flappy bird example
make flappy
# or
go run examples/flappy/main.go

# Run nakama integration example
make nakama

# Dev mode with hot reload (requires air)
make dev-flappy
```

### Building
```bash
# Cross-platform build for flappy example
make build-flappy
```

Builds output to `builds/<example>/`.

## Architecture

### Core Game Loop

The engine follows a central game loop managed by the `Game` struct ([pkg/engine/game.go](pkg/engine/game.go)):

1. **SceneManager.Start()** - Initializes the current scene and its entities
2. **Per-frame loop:**
   - `SceneManager.Update()` - Updates all entities (BG → Main → UI order)
   - `PhysicsSystem.Update()` - Applies gravity and velocity
   - `CollisionSystem.Update()` - Checks collisions and fires callbacks

### Entity System

All game objects implement the `Entity` interface ([pkg/engine/entity.go](pkg/engine/entity.go)):
```go
type Entity interface {
    Start(g *Game)      // Called once when scene loads
    Update(g *Game)     // Called every frame
    Destroy(g *Game)    // Called on scene unload
}
```

Scenes organize entities into three layers:
- **Background entities** - Rendered first (behind everything)
- **Main entities** - The core game objects
- **UI entities** - Rendered last (on top)

### Scene Management

Scenes implement `IScene` which extends `Entity` with entity container methods ([pkg/engine/scene.go](pkg/engine/scene.go)). Use `SceneManager.LoadScene()` to transition between scenes—this properly clears entities and resets physics/collision systems.

Example scene pattern:
```go
type MyScene struct {
    engine.Scene  // Embed for entity collection methods
}

func (s *MyScene) Start(g *engine.Game) {
    // Create and register entities
    s.AddEntity(myEntity)
    s.AddUIEntity(myUI)
    s.AddBGEntity(background)
}

func (s *MyScene) Update(g *engine.Game) {
    rl.ClearBackground(rl.RayWhite)
    // Entities update automatically via SceneManager
}
```

### Physics System

The `PhysicsSystem` ([pkg/engine/physics_system.go](pkg/engine/physics_system.go)) handles:
- Gravity acceleration (configurable direction and multiplier)
- Velocity-based position updates

Entities must implement `PhysicsObject` to participate:
```go
type PhysicsObject interface {
    GetPosition() rl.Vector2
    SetPosition(rl.Vector2)
    GetVelocity() rl.Vector2
    SetVelocity(rl.Vector2)
    IsGravityEnabled() bool
}
```

### Collision System

The `CollisionSystem` ([pkg/engine/collision_system.go](pkg/engine/collision_system.go)) provides AABB collision detection with three callback phases:

1. **OnCollisionEnter** - First frame of collision (implements `CollisionEnterHandler`)
2. **OnCollision** - Every frame of collision (required via `RectCollidable`)
3. **OnCollisionExit** - Frame after collision ends (implements `CollisionExitHandler`)

Colliders implement `RectCollidable`:
```go
type RectCollidable interface {
    OnCollision(g *Game, other RectCollidable)
    GetColliderRect() rl.Rectangle
    SetColliderRect(rect rl.Rectangle)
}
```

Set `IsDebug: true` on the collision system to visualize collider bounds.

### Other Systems

- **GameSave** ([pkg/engine/game_save.go](pkg/engine/game_save.go)) - JSON-based save/load interface
- **Spritesheet** ([pkg/engine/spritesheet.go](pkg/engine/spritesheet.go)) - Tile-based sprite rendering helper
- **Nakama** ([pkg/nakama/nakama.go](pkg/nakama/nakama.go)) - HTTP client for Nakama server authentication

## Example Structure

See [examples/flappy/](examples/flappy/) for a complete game example:
- `main.go` - Game initialization and scene registration
- `internal/scenes/` - Scene implementations
- `internal/entities/` - Custom entities (bird, obstacles, UI)
- `assets/` - Sprites and textures

## Development Notes

- The engine uses raylib-go for rendering, input, and window management
- Scene transitions automatically clean up physics/collision systems
- Always call `scene.ClearEntities()` before destroying a scene manually
- Collision pairs track state to fire enter/exit callbacks correctly
- Spritesheets assume uniform tile size across the texture
