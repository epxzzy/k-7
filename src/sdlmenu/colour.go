package main

import (
	"fmt"
	"math"
)

func rainbowColor(timeStep int) [3]int {
	//this is def not taken from epxzzy sabers on curseforge and modrinth
	localTimeStep := math.Mod(math.Abs(float64(timeStep)), 1536)
	timeStepInPhase := math.Mod(localTimeStep, 256);
	phaseBlue := localTimeStep / 256;
	red := colorInPhase(phaseBlue + 4, timeStepInPhase);
	green := colorInPhase(phaseBlue + 2, timeStepInPhase);
	blue := colorInPhase(phaseBlue, timeStepInPhase);
	return [3]int{red, green, blue};
}

func colorInPhase(phase, progress float64) int {
	phase = math.Mod(phase, 6);
	if (phase <= 1){
		return 0
	}
	if (phase == 2){
		return int(progress);
	} 
	if (phase <= 4){
		return 255;
	} else {
		return 255 - int(progress);
	}
}

func keygencolour(timeStep int, s, l float64) string {
	modlel := timeStep % 1536

	angle := float64(modlel) / 1536.0 * 360.0 

	fmt.Printf("angle: %f\n", angle)

	regbees := HSLtoRGB(angle, s, l)

	return RGBtoHEX(regbees[0], regbees[1], regbees[2])
}

func HSLtoRGB(h, s, l float64) [3]int {
	h = h / 360.0 // Normalize to 0-1
	if s == 0 {
		rgb := [3]float64{l, l, l}
		return [3]int{
			int(math.Round(rgb[0] * 255)),
			int(math.Round(rgb[1] * 255)),
			int(math.Round(rgb[2] * 255)),
		}
	}

	q := l
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - (l * s)
	}
	p := 2*l - q
	hue2rgb := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		if t < 1.0/6 {
			return p + (q-p)*6*t
		}
		if t < 0.5 {
			return q
		}
		if t < 2.0/3 {
			return p + (q-p)*(2.0/3-t)*6
		}
		return p
	}

	r := hue2rgb(p, q, h+1.0/3)
	g := hue2rgb(p, q, h)
	b := hue2rgb(p, q, h-1.0/3)

	return [3]int{
		int(math.Round(r * 255)),
		int(math.Round(g * 255)),
		int(math.Round(b * 255)),
	}
}

func RGBtoHEX(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
