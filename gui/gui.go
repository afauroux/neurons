//Package gui deal with presenting graphical user interface of neural networks
package gui

import (
	"image/color"
	"math"

	"github.com/faiface/pixel/pixelgl"
	"github.com/h8gi/canvas"
)

// RADIUS of the neurons in pixels
var RADIUS = 20.0

// DIST between neurons in pixels
var DIST = 40.0

// THICKNESS of neural links in pixels
var THICKNESS = 2.0

// WIDTH of the window
const WIDTH = 800

// HEIGHT of the window
const HEIGHT = 600

// scale allow the scalling of the neural net representation
func scale(s float64) {
	RADIUS *= s
	DIST *= s
	THICKNESS *= s
}

// CreateCanvas tests canvas
func CreateCanvas(nn *Network) {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     WIDTH,
		Height:    HEIGHT,
		FrameRate: 60,
		Title:     "Hebbian Discrete LTP Neural network",
	})
	// generating coordinates of neurons from their position in the 2d array and the var "DIST"
	coords := makeCoords(nn)

	c.Draw(func(ctx *canvas.Context) {
		ctx.DrawRectangle(0, 0, float64(WIDTH), float64(HEIGHT))
		ctx.SetColor(color.White)
		ctx.Fill()

		drawNeuralNetwork(ctx, nn, coords)
		if ctx.IsMouseDragged {
			for _, layer := range nn.NN {
				for _, n := range layer {
					dx := math.Abs(coords[n.ID][0]-ctx.Mouse.X) / RADIUS
					dy := math.Abs(coords[n.ID][1]-ctx.Mouse.Y) / RADIUS
					// dx and dy are in [0,1] if the user clicked inside this neuron's circle
					if dx < 1 && dy < 1 {
						if n.Potential >= 0 {
							n.Fire()
						}
					}
				}
			}

		}
		if ctx.IsKeyPressed(pixelgl.KeyPageDown) {
			scale(0.8)
			coords = makeCoords(nn)
		}

		if ctx.IsKeyPressed(pixelgl.KeyPageUp) {
			scale(1.2)
			coords = makeCoords(nn)
		}

	})
}

// Click is called whenever a neuron is clicked
func Click(n Neuron) {
	n.Fire()
}

// drawLink draw a line to represent axones or dendrites
// which color represent the weight from green (MAXSIG)
// to red (-MAXSIG)
func drawLink(ctx *canvas.Context, X0, Y0, X1, Y1, weight float64) {
	ctx.DrawLine(X0, Y0, X1, Y1)
	w := weight / MAXSIG
	if w >= 0 {
		ctx.SetRGB(0, w, 0)
	} else {
		ctx.SetRGB(math.Abs(w), 0, 0)
	}

	ctx.SetLineWidth(THICKNESS * 2)
	ctx.Stroke()
}

// drawNeuron as a circle with a variable inner brightness
// representing its activation potential. Also the stroke
// is red whenever the potential is negative
func drawNeuron(ctx *canvas.Context, X, Y, potential float64) {
	//inner neuron represent potential
	ctx.DrawCircle(X, Y, RADIUS)

	// p from 0 to 110 neutral 10 -> 0 and 1
	p := (float64(potential) + math.Abs(LOWEND)) / (float64(TRESH) + math.Abs(LOWEND))

	ctx.SetRGB(p, p, p)
	ctx.Fill()

	//contour
	ctx.DrawCircle(X, Y, RADIUS)
	var r float64
	if potential < 0 {
		r = 1
	}

	ctx.SetRGB(r, 0, 0)
	ctx.SetLineWidth(THICKNESS)
	ctx.Stroke()
}

// drawNeuralNetwork as a bunch of circle and lines with explicit dynamic colors
// see drawLink and drawNeuron for more explainations
func drawNeuralNetwork(ctx *canvas.Context, nn *Network, coord map[int][2]float64) {
	for _, layer := range nn.NN {
		for _, n := range layer {
			for _, d := range n.Dendrites {
				drawLink(ctx,
					coord[d.Pre.ID][0],
					coord[d.Pre.ID][1],
					coord[n.ID][0],
					coord[n.ID][1],
					d.SignedWeight(),
				)
			}
		}
	}
	for _, layer := range nn.NN {
		for _, n := range layer {
			drawNeuron(ctx,
				coord[n.ID][0],
				coord[n.ID][1],
				n.Potential,
			)
		}
	}
}

func layer2coord(net Network, i, j int) (x, y float64) {
	x = float64(2*int(DIST)*i) + WIDTH/2
	y = float64(2*int(DIST)*(j-len(net.NN[i])/2)) + HEIGHT/2
	return x, y
}

func coord2layer(net Network, x, y float64) (i, j int) {
	i = int((x - WIDTH/2) / 2 * DIST)
	j = int((y-HEIGHT/2)/2*DIST) + len(net.NN[i])/2
	return i, j
}

// makeCoords generate a etping between neurons IDs and their XY coordinates
func makeCoords(nn *Network) (coords map[int][2]float64) {
	coords = make(map[int][2]float64)
	for i, layer := range nn.NN {
		for j, n := range layer {
			coords[n.ID] = [2]float64{
				float64(2*int(DIST)*i) + WIDTH/2,
				float64(2*int(DIST)*(j-len(layer)/2)) + HEIGHT/2,
			}
		}
	}
	return coords
}
