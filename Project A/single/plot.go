package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"fmt"
	"image/color"
)

// Plotinum compatible data list type
type _points struct {
	X, Y []float64
}

// Len makes _points satisfy plotter.XYer interface
// returns data length
func (p _points) Len() int {
	return len(p.Y)
}

// XY makes _points satisfy plotter.XYer interface
// returns x and y values depending on index
func (p _points) XY(i int) (x, y float64) {
	x = p.X[i]
	y = p.Y[i]
	return
}

// Preset colours for plotting multiple lines
var cols = []color.RGBA{
	color.RGBA{R: 255, A: 255},               // red
	color.RGBA{G: 255, A: 255},               // green
	color.RGBA{B: 255, A: 255},               // blue
	color.RGBA{R: 255, G: 255, A: 255},       // yellow
	color.RGBA{R: 128, G: 0, A: 128},         // purple
	color.RGBA{R: 139, G: 69, B: 19, A: 255}, // brown
}

// Function called by rest of code to plot graphs
func Plot(fname string, title string, labels []string, xdata []float64, ydata ...[]float64) {
	if len(fname) == 0 {
		fname = "output"
	}
	fname = fmt.Sprintf("out/%v h=%v m=%v l=%v γ=%v.png", fname, h, m, l, gamma)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	fmt.Print(">> Plotting \"" + title + "\" to \"" + fname + "\"")

	// Add grid
	p.Add(plotter.NewGrid())
	p.Title.Text = title
	p.Title.Padding = 10.0
	p.X.Label.Text = labels[0]
	p.Y.Label.Text = labels[1]
	// p.Y.Min = -60.0
	// p.Y.Max = 85.0
	p.Legend.Top = true
	p.Legend.Left = true
	p.Legend.YOffs = -5.0
	p.Legend.YOffs = -15.0
	p.Legend.Padding = 10.0
	p.Legend.ThumbnailWidth = 30.0

	var s *plotter.Scatter

	// This is to allow for plotting of vertical lines for stability graphs
	stabtest := false
	last := len(ydata)
	if (len(labels) + 4) == last {
		stabtest = true
		last -= 6
	}

	// Loop through data and
	for i := 0; i < last; i++ {
		s = plotter.NewScatter(&_points{xdata, ydata[i]})
		s.GlyphStyle.Color = cols[i]
		s.GlyphStyle.Radius = 1
		s.GlyphStyle.Shape = plot.CircleGlyph{}
		p.Add(s)
		if last > 1 {
			p.Legend.Add(labels[i+2], s)
		}
	}

	if stabtest {
		for i := 0; i < 3; i++ {
			s = plotter.NewScatter(&_points{ydata[2*i+3], ydata[2*i+4]})
			s.GlyphStyle.Color = cols[i]
			s.GlyphStyle.Radius = 1
			s.GlyphStyle.Shape = plot.CircleGlyph{}
			p.Add(s)
		}
	}

	// Save the plot to a PNG file.
	if err := p.Save(7, 7, fname); err != nil {
		panic(err)
	}
	fmt.Print(" ... Done\n")
}
