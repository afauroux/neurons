package neurons

import (
	"image/color"
	"math"

	"github.com/h8gi/canvas"
)

const radius = 20
const dist = 40
const width = 800
const height = 500

// DrawNeuralNetwork ibid
func DrawNeuralNetwork(ctx *canvas.Context, coord map[*Neuron][2]float64, links map[*Neuron]map[int][4]float64) {
	// Drawing connections
	for n, m := range links {
		for id, l := range m {
			ctx.DrawLine(l[0], l[1], l[2], l[3])
			w := float64(n.Weights[id]) / float64(maxSig)
			ctx.SetRGB(w, w, w)
			ctx.SetLineWidth(2)
			ctx.Stroke()
		}
	}
	// Drawing neurons
	for n, c := range coord {
		//inner neuron represent potential
		ctx.DrawCircle(c[0], c[1], radius)
		p := float64(n.Potential) / float64(tresh)
		ctx.SetRGB(p, p, p)
		ctx.Fill()

		//contour
		ctx.DrawCircle(c[0], c[1], radius)
		ctx.SetRGB(0, 0, 0)
		ctx.SetLineWidth(2)
		ctx.Stroke()
	}
}

// MakeNeuralNetwork create the coordinates of neurons and of the links connecting them
func MakeNeuralNetwork(n [][]*Neuron) (coord map[*Neuron][2]float64, links map[*Neuron]map[int][4]float64) {
	coord = make(map[*Neuron][2]float64)
	for i, layer := range n {
		for j, neuron := range layer {
			coord[neuron] = [2]float64{float64(2*dist*(len(n)/2-i) + width/2), float64(2*dist*(j-len(layer)/2) + height/2)}
		}
	}
	links = make(map[*Neuron]map[int][4]float64) // map[Neuron][Parent.ID]{x0,y0,x1,y1}
	for _, layer := range n {
		for _, neuron := range layer {
			links[neuron] = make(map[int][4]float64, len(neuron.Parents))
			for _, pre := range neuron.Parents {
				links[neuron][pre.ID] = [4]float64{coord[pre][0], coord[pre][1], coord[neuron][0], coord[neuron][1]}
			}
		}
	}
	return coord, links
}

// TestCanvas tests canvas
func TestCanvas(n [][]*Neuron) {
	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     width,
		Height:    height,
		FrameRate: 60,
	})

	coords, links := MakeNeuralNetwork(n)
	c.Draw(func(ctx *canvas.Context) {
		ctx.DrawRectangle(0, 0, float64(width), float64(height))
		ctx.SetColor(color.White)
		ctx.Fill()

		DrawNeuralNetwork(ctx, coords, links)
		if ctx.IsMouseDragged {
			for n, c := range coords {
				if (math.Abs(c[0]-ctx.Mouse.X) <= radius) && (math.Abs(c[1]-ctx.Mouse.Y) <= radius) {
					n.Fire()
				}
			}
		}

	})
}
