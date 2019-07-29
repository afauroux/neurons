package gui

import (
	"image/color"
	"math"

	"github.com/faiface/pixel/pixelgl"

	nn "github.com/afauroux/neurons/nnetwork"
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

func getCoordsandlinks(nmap [][]*nn.Neuron) (coords []*Coord, links []*Link) {
	for _, layer := range nmap {
		for _, n := range layer {
			coords = append(coords, getCoord(n))
			links = append(links, getLinks(n)...)
		}
	}
	return coords, links
}

// CreateCanvas tests canvas
func CreateCanvas(nmap [][]*nn.Neuron) {
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
					c.N.Input <- -1
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
		w := float64(l.N.Weights[l.ID]) / float64(nn.MAXSIG)
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
		p := (float64(c.N.Potential) + math.Abs(nn.LOWEND)) / (float64(nn.TRESH) + math.Abs(nn.LOWEND))

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
	N    *nn.Neuron
}

// Link is a mapping of synapses to
// pre and post-synaptic neuron coordinates
type Link struct {
	X0, Y0, X1, Y1 float64
	ID             int
	N              *nn.Neuron
}

func getCoord(n *nn.Neuron) *Coord {
	return &Coord{
		X: float64(2*int(dist)*n.X + width/2),
		Y: float64(2*int(dist)*n.Y + height/2),
		N: n,
	}
}

func getLinks(n *nn.Neuron) []*Link {
	var links []*Link
	for id, p := range n.Parents {
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
