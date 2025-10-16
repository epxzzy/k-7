package main

import (
  "math"
  "math/rand"
  "github.com/tfriedel6/canvas"
)

func updatePatternPosition(pattern *PatternConfig) {
  radians := degreesToRadians(pattern.Angle)
  offsetX := math.Cos(radians) * pattern.Speed
  offsetY := math.Sin(radians) * pattern.Speed
  /*
pattern.X += offsetX
pattern.Y += offsetY

if pattern.X >= pattern.TileSize || pattern.X <= -pattern.TileSize ||
pattern.Y >= pattern.TileSize || pattern.Y <= -pattern.TileSize {
pattern.X = offsetX
pattern.Y = offsetY
}
*/

  // Update position with continuous wrapping
  pattern.X = math.Mod(pattern.X+offsetX, pattern.TileSize)
  pattern.Y = math.Mod(pattern.Y+offsetY, pattern.TileSize)

  // Handle negative values for proper wrapping in all directions
  if pattern.X < 0 {
    pattern.X += pattern.TileSize
  }
  if pattern.Y < 0 {
    pattern.Y += pattern.TileSize
  }
}

func drawPattern(cv *canvas.Canvas, p *PatternConfig, width, height int) {
  cv.SetStrokeStyle(p.Color)
  cv.SetLineWidth(p.Width)

  ts := p.TileSize
  w := float64(width)
  h := float64(height)

  for x := -ts; x < w; x += ts {
    for y := -ts; y < h; y += ts {
      cv.BeginPath()
      cv.Rect(float64(x)+p.X, float64(y)+p.Y, ts, ts)
      cv.Stroke()
    }
  }
}

func updateNormalMode(cv *canvas.Canvas, state *AnimationState, width, height int) {
  go updatePatternPosition(&state.Normal.Back)
  drawPattern(cv, &state.Normal.Back, width, height)
  go updatePatternPosition(&state.Normal.Front)
  drawPattern(cv, &state.Normal.Front, width, height)
}

func updateHyperdriveMode(cv *canvas.Canvas, state *AnimationState, width, height int) {
  state.Hyperdrive.Front.Speed = 10
  state.Hyperdrive.Back.Speed = 10

  lightBlue := []string{"#42C2FF", "#85F4FF", "#B8FFF9"}
  darkBlue := []string{"#0B2447", "#19376D", "#576CBC"}

  currWidth := rand.Float64()*(17-0.5) + 0.5
  colour := rand.Intn(3)

  state.Hyperdrive.Back.Color = darkBlue[colour]
  state.Hyperdrive.Back.Width = currWidth
  state.Hyperdrive.Front.Color = lightBlue[colour]
  state.Hyperdrive.Front.Width = currWidth

  go updatePatternPosition(&state.Hyperdrive.Back)
  drawPattern(cv, &state.Hyperdrive.Back, width, height)
  go updatePatternPosition(&state.Hyperdrive.Front)
  drawPattern(cv, &state.Hyperdrive.Front, width, height)
}

func updateKeygenMode(cv *canvas.Canvas, state *AnimationState, width, height int) {
  state.Keygen.Front.Speed = 1
  state.Keygen.Back.Speed = 1

  rgb_colour := []string{"#42C2FF", "#85F4FF", "#B8FFF9"}
  inverted_rgb_colour := []string{"#0B2447", "#19376D", "#576CBC"}

  colour := rand.Intn(3)

  state.Keygen.Back.Color = inverted_rgb_colour[colour]
  state.Keygen.Front.Color = rgb_colour[colour]

  go updatePatternPosition(&state.Keygen.Back)
  drawPattern(cv, &state.Keygen.Back, width, height)
  go updatePatternPosition(&state.Keygen.Front)
  drawPattern(cv, &state.Keygen.Front, width, height)
}


func updateRedShakeMode(cv *canvas.Canvas, state *AnimationState, width, height int) {
  if state.ShakeCounter%120 == 0 {
    state.RedShake.Front.Angle, state.RedShake.Back.Angle = 
    state.RedShake.Back.Angle, state.RedShake.Front.Angle
    state.RedShake.Front.Color = "#DC0000"
    state.RedShake.Back.Color = "#850000"
  }

  state.RedShake.Front.X += rand.Float64()*10 - 5
  state.RedShake.Front.Y += rand.Float64()*10 - 5
  state.RedShake.Back.X += rand.Float64()*15 - 5
  state.RedShake.Back.Y += rand.Float64()*15 - 5

  go updatePatternPosition(&state.RedShake.Front)
  drawPattern(cv, &state.RedShake.Front, width, height)
  go updatePatternPosition(&state.RedShake.Back)
  drawPattern(cv, &state.RedShake.Back, width, height)

  state.ShakeCounter++
}

func updateGlitchMatrixMode(cv *canvas.Canvas, state *AnimationState, width, height int) {
  if state.GlitchCounter >= state.GlitchDur {
    state.GlitchState = !state.GlitchState
    state.GlitchCounter = 0

    lightGreen := []string{"#38b000", "#70e000", "#9ef01a"}
    darkGreen := []string{"#538d22", "#538d22", "#aad576"}
    colour := rand.Intn(3)

    if state.GlitchState {
      state.GlitchDur = rand.Intn(20)
      state.GlitchMatrix.Front = state.Normal.Back
      state.GlitchMatrix.Back = state.Normal.Front
      state.GlitchMatrix.Front.Color = lightGreen[colour]
      state.GlitchMatrix.Back.Color = darkGreen[colour]
      state.GlitchMatrix.Front.Speed = 3
      state.GlitchMatrix.Back.Speed = 3
    } else {
      state.GlitchDur = rand.Intn(10)
      state.GlitchMatrix.Front = state.Normal.Front
      state.GlitchMatrix.Back = state.Normal.Back
      state.GlitchMatrix.Front.Speed = 0.3
      state.GlitchMatrix.Back.Speed = 0.3
    }
  }

  if state.GlitchState && rand.Float64() < 0.47 {
    drawPattern(cv, &state.GlitchMatrix.Front, width, height)
    drawPattern(cv, &state.GlitchMatrix.Back, width, height)
  } else {
    go updatePatternPosition(&state.GlitchMatrix.Front)
    drawPattern(cv, &state.GlitchMatrix.Front, width, height)
    go updatePatternPosition(&state.GlitchMatrix.Back)
    drawPattern(cv, &state.GlitchMatrix.Back, width, height)
  }

  state.GlitchCounter++
}
