package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create window and canvas
	wnd, cv, err := sdlcanvas.CreateWindow(1024, 768, "survey_program")
	if err != nil {
		panic(err);
	}
	defer wnd.Destroy()	


	go initWsClient();

	state := NewAnimationState()
	mode := "NORMAL"

  wnd.KeyDown = func(scancode int, rn rune, name string) {
    switch scancode {
    case sdl.SCANCODE_N:
      mode ="NORMAL"
      fmt.Println("current mode is normal")

    case sdl.SCANCODE_H:
      mode = "HYPERDRIVE"
      fmt.Println("current mode is hyperdrive")

    case sdl.SCANCODE_L:
      mode = "KEYGEN"
      fmt.Println("current mode is keygen")

    case sdl.SCANCODE_R:
      mode = "REDSHAKE"
      fmt.Println("current mode is redshake")

    case sdl.SCANCODE_G:
      mode = "GLITCHMATRIX"
      fmt.Println("current mode is glitchmatrix")

    default:
			SendToServer(name)
      fmt.Println("sending over " + name)

    }

	if RecieveQueueC.IsEmpty() == false {
		for {
			if RecieveQueueC.GetLength() == 0 {break}

				println("Recieved this from the server: "+RecieveQueueC.Dequeue())
			}
		}
	}


	wnd.MainLoop(func() {
		w, h := float64(cv.Width()), float64(cv.Height())
		cv.SetFillStyle("#000")
		cv.FillRect(0, 0, w, h)

		UpdateAnimation(cv, state, mode )
	})
}




func UpdateAnimation(cv *canvas.Canvas, state *AnimationState, mode string) {
	width, height := cv.Size()

	// Clear canvas
	cv.SetFillStyle(0, 0, 0, 1)
	cv.FillRect(0, 0, float64(width), float64(height))

	// Update and draw current mode
	switch mode{
	case "NORMAL":
		updateNormalMode(cv, state, width, height)
	case "HYPERDRIVE":
		updateHyperdriveMode(cv, state, width, height)
	case "KEYGEN":
		updateKeygenMode(cv, state, width, height)
	case "REDSHAKE":
		updateRedShakeMode(cv, state, width, height)
	case "GLITCHMATRIX":
		updateGlitchMatrixMode(cv, state, width, height)
	}

	if RecieveQueueC.IsEmpty() == false {
		for {
			if RecieveQueueC.GetLength() == 0 {break}

				println("Recieved this from the server: "+RecieveQueueC.Dequeue())
			}
		}

}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
