---
name: create-scene
description: Scaffold scene files for M-Engine (Go game engine). Creates Go structs with Start/Update/Destroy lifecycle methods that embed engine.Scene. Use this skill whenever the user asks to create, add, make, or scaffold a scene, level, game state, menu, screen, game mode, or any similar scene terminology in the context of Go game development or the m-engine codebase. This includes phrases like "create a menu scene", "add a game level", "make a settings screen", "scaffold a lobby scene", etc.
---

# Create Scene

Scaffolds a new scene file for the M-Engine game engine.

## When to use

Use this whenever the user wants to create a new scene/level/screen in their M-Engine project. The user might say:
- "create a Menu scene"
- "add a GameLevel scene with physics"
- "make a Settings screen"
- "scaffold a Lobby scene for the nakama example"

## How it works

1. Parse the scene name from user input (convert to PascalCase, add "Scene" suffix)
2. Determine the target example (flappy, nakama, or create new)
3. Create the scene file with lifecycle methods
4. Add optional system setup based on user request

## File location

Scenes go in: `examples/<example>/internal/scenes/<name>.scene.go` or `<name>_scene.go`

## Basic scene template

```go
package scenes

import (
    "github.com/MarcelArt/m-engine/pkg/engine"
    rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneNameScene struct {
    engine.Scene
}

func (s *SceneNameScene) Start(g *engine.Game) {
    // Create and register entities
    // s.AddEntity(entity)
    // s.AddUIEntity(uiEntity)
    // s.AddBGEntity(background)
}

func (s *SceneNameScene) Update(g *engine.Game) {
    rl.ClearBackground(rl.RayWhite)
    // Entities update automatically via SceneManager
}

func (s *SceneNameScene) Destroy(g *engine.Game) {
    // Cleanup before scene unload
}

var _ engine.IScene = &SceneNameScene{}
```

## Optional system setup

Add these in Start() based on user request:

### Physics system
```go
physics := engine.NewPhysicsSystem(rl.NewVector2(0, 1), 800)
// physics.AddEntity(someEntity)
g.SetPhysicsSystem(physics)
```

### Collision system
```go
collision := engine.NewCollisionSystem(false) // true for debug mode
// collision.AddRectCollidable(someEntity)
g.SetCollisionSystem(collision)
```

### Background entity
```go
background := &entities.Background{
    Texture: rl.LoadTexture("examples/<example>/assets/background.png"),
}
s.AddBGEntity(background)
```

### Game state entity
```go
state := &entities.GameState{
    Score: 0,
    // ... other fields
}
s.AddEntity(state)
```

## Instructions

1. Extract scene name - convert to PascalCase, add "Scene" suffix if not present (e.g., "menu" → "MenuScene", "game_level" → "GameLevelScene")
2. Detect target example from context or current working directory
3. Ask clarifying questions if the example is ambiguous
4. Determine which systems/features to add (physics, collision, background)
5. Create the file with proper Go formatting
6. Verify compilation with `go build`
7. Report the created file path

## Examples

| User input | Scene name | Features |
|------------|------------|----------|
| "create a Menu scene" | MenuScene | basic |
| "add a GameLevel scene with physics" | GameLevelScene | physics system |
| "make a Settings screen" | SettingsScene | basic |
| "create a Lobby scene for nakama" | LobbyScene | basic, nakama context |
