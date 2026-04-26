package entities

import (
	"github.com/MarcelArt/m-engine/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextInput struct {
	Value           string
	Placeholder     string
	IsDisabled      bool
	IsSecret        bool
	OnValueChange   func(string)
	BackgroundColor rl.Color
	TextColor       rl.Color
	FontSize        int32
	Rect            rl.Rectangle

	isFocused      bool
	frames         int32
	backspaceTimer float32
}

func (t *TextInput) Start(g *engine.Game) {
}

func (t *TextInput) Update(g *engine.Game) {
	mouse := rl.GetMousePosition()
	isHovered := rl.CheckCollisionPointRec(mouse, t.Rect)

	clicked := isHovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)
	if clicked && !t.IsDisabled {
		t.isFocused = true
	} else if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !isHovered {
		t.isFocused = false
	}

	t.handleKeyboard()

	bgColor := t.BackgroundColor
	if t.isFocused {
		borderColor := rl.NewColor(100, 150, 255, 255)
		rl.DrawRectangleLinesEx(t.Rect, 2, borderColor)
	} else {
		borderColor := rl.NewColor(150, 150, 150, 255)
		rl.DrawRectangleLinesEx(t.Rect, 1, borderColor)
	}
	rl.DrawRectangleRec(t.Rect, bgColor)

	padding := float32(10)
	displayText := t.Value
	if displayText == "" {
		displayText = t.Placeholder
	}
	if t.IsSecret && len(t.Value) > 0 {
		displayText = ""
		for range t.Value {
			displayText += "*"
		}
	}

	textX := int32(t.Rect.X + padding)
	textY := int32(t.Rect.Y + (t.Rect.Height / 2) - float32(t.FontSize)/2)

	textColor := t.TextColor
	if t.Value == "" {
		textColor = rl.NewColor(150, 150, 150, 255)
	}
	rl.DrawText(displayText, textX, textY, t.FontSize, textColor)

	if t.isFocused && t.frames%60 < 30 {
		textWidth := rl.MeasureText(displayText, t.FontSize)
		cursorX := textX + textWidth
		cursorHeight := t.FontSize
		rl.DrawRectangle(cursorX, textY, 2, cursorHeight, t.TextColor)
	}
	t.frames++
}

func (t *TextInput) Destroy(g *engine.Game) {
	t.isFocused = false
}

func (t *TextInput) handleKeyboard() {
	t.backspaceTimer += rl.GetFrameTime()

	if !t.isFocused || t.IsDisabled {
		return
	}

	if rl.IsKeyDown(rl.KeyEnter) || rl.IsKeyDown(rl.KeyKpEnter) {
		t.isFocused = false
		return
	}

	if rl.IsKeyDown(rl.KeyBackspace) && t.backspaceTimer > 0.08 {
		if len(t.Value) > 0 {
			t.Value = t.Value[:len(t.Value)-1]
			if t.OnValueChange != nil {
				t.OnValueChange(t.Value)
			}
			t.backspaceTimer = 0
		}
	}

	char := rl.GetCharPressed()
	for char != 0 {
		t.Value += string(rune(char))
		if t.OnValueChange != nil {
			t.OnValueChange(t.Value)
		}
		char = rl.GetCharPressed()
	}
}

var _ engine.Entity = &TextInput{}
