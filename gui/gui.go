package gui

import (
	"image/color"
	"math"

	"github.com/faiface/pixel/pixelgl"

	N "github.com/afauroux/neurons/neuron"
	"github.com/h8gi/canvas"
)

var radius = 20.0
var dist = 40.0
var thickness = 2.0

func scale(s float64) {
	radius *= s
	dist *= s
	thickness *= s
}

const clickStrength = 20
const width = 800
const height = 600

func getCoordsandlinks(nmap [][]N.Neuron) (coords []*Coord, links []*Link) {
	for _, layer := range nmap {
		for _, n := range layer {
			coords = append(coords, getCoord(n))
			links = append(links, getLinks(n)...)
		}
	}
	return coords, links
}

// CreateCanvas tests canvas
func CreateCanvas(nmap [][]N.Neuron) {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     width,
		Height:    height,
		FrameRate: 60,
		Title:     "Hebbian Discrete LTP Neural network",
	})
	// generating list of Coord and Links for faster ploting
	coords, links := getCoordsandlinks(nmap)

	c.Draw(func(ctx *canvas.Context) {
		ctx.DrawRectangle(0, 0, float64(width), float64(height))
		ctx.SetColor(color.White)
		ctx.Fill()

		DrawNeuralNetwork(ctx, coords, links)
		if ctx.IsMouseDragged {
			for _, c := range coords {
				if (math.Abs(c.X-ctx.Mouse.X) <= radius) && (math.Abs(c.Y-ctx.Mouse.Y) <= radius) {
					c.N.GetInput() <- -1
				}
			}
		}
		if ctx.IsKeyPressed(pixelgl.KeyPageDown) {
			scale(0.8)
			coords, links = getCoordsandlinks(nmap)
		}

		if ctx.IsKeyPressed(pixelgl.KeyPageUp) {
			scale(1.2)
			coords, links = getCoordsandlinks(nmap)
		}

	})
}

// DrawNeuralNetwork ibid
func DrawNeuralNetwork(ctx *canvas.Context, coord []*Coord, links []*Link) {
	// Drawing connections
	for _, l := range links {
		ctx.DrawLine(l.X0, l.Y0, l.X1, l.Y1)
		w := float64(l.N.Weights[l.ID]) / float64(N.MAXSIG)
		if w >= 0 {
			ctx.SetRGB(0, w, 0)
		} else {
			ctx.SetRGB(math.Abs(w), 0, 0)
		}

		ctx.SetLineWidth(thickness * 2)
		ctx.Stroke()

	}
	// Drawing neurons
	for _, c := range coord {
		//inner neuron represent potential
		ctx.DrawCircle(c.X, c.Y, radius)

		// p from 0 to 110 neutral 10 -> 0 and 1
		p := (float64(c.N.Potential) + math.Abs(N.LOWEND)) / (float64(N.TRESH) + math.Abs(N.LOWEND))

		ctx.SetRGB(p, p, p)
		ctx.Fill()

		//contour
		ctx.DrawCircle(c.X, c.Y, radius)
		var r float64
		if c.N.Potential < 0 {
			r = 1
		}

		ctx.SetRGB(r, 0, 0)
		ctx.SetLineWidth(thickness)
		ctx.Stroke()
	}
}

// Coord maps neurons to coordinates
type Coord struct {
	X, Y float64
	N    N.Neuron
}

// Link is a mapping of synapses to
// pre and post-synaptic neuron coordinates
type Link struct {
	X0, Y0, X1, Y1 float64
	ID             int
	N              N.Neuron
}

func getCoord(n N.Neuron) *Coord {
	return &Coord{
		X: float64(2*int(dist)*N.GetX() + width/2),
		Y: float64(2*int(dist)*N.GetY() + height/2),
		N: n,
	}
}

func getLinks(n N.Neuron) []*Link {
	var links []*Link
	for id, p := range N.GetParents() {
		c0 := getCoord(n)
		c1 := getCoord(p)
		links = append(links, &Link{
			X0: c0.X,
			Y0: c0.Y,
			X1: c1.X,
			Y1: c1.Y,
			ID: id,
			N:  n,
		})
	}
	return links
}
