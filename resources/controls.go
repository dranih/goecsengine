package resources

import (
	"fmt"

	"github.com/x-hgg-x/goecsengine/utils"

	"github.com/hajimehoshi/ebiten"
)

// Key is a US keyboard key
type Key struct {
	Key ebiten.Key
}

// UnmarshalText fills structure fields from text data
func (k *Key) UnmarshalText(text []byte) error {
	if key, ok := utils.KeyMap[string(text)]; ok {
		k.Key = key
		return nil
	}
	return fmt.Errorf("unknown key: '%s'", string(text))
}

// MouseButton is a mouse button
type MouseButton struct {
	MouseButton ebiten.MouseButton
}

// UnmarshalText fills structure fields from text data
func (b *MouseButton) UnmarshalText(text []byte) error {
	if mouseButton, ok := utils.MouseButtonMap[string(text)]; ok {
		b.MouseButton = mouseButton
		return nil
	}
	return fmt.Errorf("unknown mouse button: '%s'", string(text))
}

// ControllerButton is a gamepad button
type ControllerButton struct {
	ID            int
	GamepadButton ebiten.GamepadButton
}

// UnmarshalTOML fills structure fields from TOML data
func (b *ControllerButton) UnmarshalTOML(i interface{}) error {
	data := i.(map[string]interface{})
	if gamepadButton, ok := utils.GamepadButtonMap[data["button"].(string)]; ok {
		b.ID = int(data["id"].(int64))
		b.GamepadButton = gamepadButton
		return nil
	}
	return fmt.Errorf("unknown gamepad button: '%s'", data["button"].(string))
}

// Button can be a US keyboard key, a mouse button or a gamepad button
type Button struct {
	Type             string
	Key              *Key
	MouseButton      *MouseButton      `toml:"mouse_button"`
	ControllerButton *ControllerButton `toml:"controller"`
}

// Emulated is an emulated axis
type Emulated struct {
	Pos Button
	Neg Button
}

// ControllerAxis is a gamepad axis
type ControllerAxis struct {
	ID       int
	Axis     int
	Invert   bool
	DeadZone float64 `toml:"dead_zone"`
}

// MouseAxis is a mouse axis
type MouseAxis struct {
	Axis int
}

// Axis can be an emulated axis, a gamepad axis or a mouse axis
type Axis struct {
	Type           string
	Emulated       *Emulated
	ControllerAxis *ControllerAxis `toml:"controller_axis"`
	MouseAxis      *MouseAxis      `toml:"mouse_axis"`
}

// Action contains buttons combinations with settings
type Action struct {
	// Combinations contains buttons combinations
	Combinations [][]Button
	// Once determines if the action should be triggered every frame when the button is pressed (default) or only once
	Once bool
}

// Controls contains input controls
type Controls struct {
	// Axes contains axis controls, used for inputs represented by a float value from -1 to 1
	Axes map[string]Axis
	// Actions contains buttons combinations, used for general inputs
	Actions map[string]Action
}

// InputHandler contains input axis values and actions corresponding to specified controls
type InputHandler struct {
	// Axes contains input axis values
	Axes map[string]float64
	// Actions contains input actions
	Actions map[string]bool
}
