package main

type PatternConfig struct {
	Angle    float64
	Speed    float64
	TileSize float64
	Width    float64
	Color    string
	X        float64
	Y        float64
}

type AnimationState struct {
	CurrentMode   string
	Normal        ModeConfig
	Hyperdrive    ModeConfig
	Keygen        ModeConfig
	RedShake      ModeConfig
	GlitchMatrix  ModeConfig
	ShakeCounter  int
	GlitchCounter int
	GlitchState   bool
	GlitchDur     int
}

type ModeConfig struct {
	Front PatternConfig
	Back  PatternConfig
}

func NewAnimationState() *AnimationState {
	baseFront := PatternConfig{
		Angle:    225, Speed: 0.3, TileSize: 60,
		Width: 5, Color: "#490049", X: 60, Y: 60,
	}

	baseBack := PatternConfig{
		Angle:    45, Speed: 0.3, TileSize: 60,
		Width: 5, Color: "#240024", X: 0, Y: 0,
	}

	return &AnimationState{
		CurrentMode: "NORMAL",
		Normal:      ModeConfig{Front: baseFront, Back: baseBack},
		Hyperdrive:  ModeConfig{Front: baseFront, Back: baseBack},
		Keygen:  		ModeConfig{Front: baseFront, Back: baseBack},
		RedShake:    ModeConfig{Front: baseFront, Back: baseBack},
		GlitchMatrix: ModeConfig{
			Front: baseFront,
			Back:  baseBack,
		},
	}
}
