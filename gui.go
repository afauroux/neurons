package neurons

import (
	"image/color"
	"math"

	"github.com/h8gi/canvas"
)

const radius = 20
const dist = 40
const width = 800
const height = 600

// CreateCanvas tests canvas
func CreateCanvas(nmap [][]*Neuron) {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     width,
		Height:    height,
		FrameRate: 20,
		Title:     "Hebbian Discrete LTP Neural network",
	})
	// generating list of Coord and Links for faster ploting
	var coords []*Coord
	var links []*Link
	for _, layer := range nmap {
		for _, n := range layer {
			coords = append(coords, getCoord(n))
			links = append(links, getLinks(n)...)
		}
	}

	c.Draw(func(ctx *canvas.Context) {
		ctx.DrawRectangle(0, 0, float64(width), float64(height))
		ctx.SetColor(color.White)
		ctx.Fill()

		DrawNeuralNetwork(ctx, coords, links)
		if ctx.IsMouseDragged {
			for _, c := range coords {
				if (math.Abs(c.X-ctx.Mouse.X) <= radius) && (math.Abs(c.Y-ctx.Mouse.Y) <= radius) {
					c.N.Fire()
				}
			}
		}

	})
}

// DrawNeuralNetwork ibid
func DrawNeuralNetwork(ctx *canvas.Context, coord []*Coord, links []*Link) {
	// Drawing connections
	for _, l := range links {
		ctx.DrawLine(l.X0, l.Y0, l.X1, l.Y1)
		w := 1 - float64(l.N.Weights[l.ID])/float64(maxSig)
		if w >= 0 {
			ctx.SetRGB(0, w, 0)
		} else {
			w = -w
			ctx.SetRGB(w, 0, 0)
		}

		ctx.SetLineWidth(2)
		ctx.Stroke()

	}
	// Drawing neurons
	for _, c := range coord {
		//inner neuron represent potential
		ctx.DrawCircle(c.X, c.Y, radius)
		p := float64(c.N.Potential) / float64(tresh)
		ctx.SetRGB(p, p, p)
		ctx.Fill()

		//contour
		ctx.DrawCircle(c.X, c.Y, radius)
		ctx.SetRGB(0, 0, 0)
		ctx.SetLineWidth(2)
		ctx.Stroke()
	}
}

// Coord maps neurons to coordinates
type Coord struct {
	X, Y float64
	N    *Neuron
}

// Link is a mapping of synapses to
// pre and post-synaptic neuron coordinates
type Link struct {
	X0, Y0, X1, Y1 float64
	ID             int
	N              *Neuron
}

func getCoord(n *Neuron) *Coord {
	return &Coord{
		X: float64(2*dist*n.X + width/2),
		Y: float64(2*dist*n.Y + height/2),
		N: n,
	}
}

func getLinks(n *Neuron) []*Link {
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
