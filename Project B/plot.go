package main

import (
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"fmt"
	"image/color"
)

// Plotinum compatible data list type
type valuer struct {
	data []float64
}

// Len makes valuer satisfy plotter.Valuer interface
// returns data length
func (v valuer) Len() int {
	return len(v.data)
}

// Value makes valuer satisfy plotter.Valuer interface
// returns value depending on index
func (v valuer) Value(i int) float64 {
	return v.data[i]
}

// createHistogram creates a histogram from given data and bin numbers
func createHistogram(data []float64, n int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	histdata := valuer{data}
	p.Add(plotter.NewHist(histdata, n))
	p.X.Label.Text = "time / ps"
	p.Y.Label.Text = "frequency"
	p.Title.Text = fmt.Sprintf("Frequency of lifetime data from lifetime.txt. %v bins.", n)

	if err := p.Save(5, 5, fmt.Sprintf("out/Histogram with %v bins.png", n)); err != nil {
		panic(err)
	}
}

// Plotinum compatible coordinates
type xyer struct {
	X, Y []float64
}

// Len makes xyer satisfy plotter.XYer interface
// returns length of data set
func (xy xyer) Len() int {
	return len(xy.X)
}

// XY makes xyer satisfy plotter.XYer interface
// returns x, y coordinates
func (xy xyer) XY(i int) (float64, float64) {
	return xy.X[i], xy.Y[i]
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

// createLine creates a line graph from provided x,y data and title
func createLine(xdat, ydat [][]float64, ylab []string, title string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	p.Title.Text = title
	p.Legend.Top = true
	p.Legend.XOffs = -10.0
	p.Legend.YOffs = -10.0

	var scatdata xyer
	var s *plotter.Line
	for i, _ := range ydat {
		scatdata = xyer{xdat[i], ydat[i]}
		s = plotter.NewLine(scatdata)
		s.LineStyle.Width = 2
		s.LineStyle.Color = cols[i]
		p.Add(s)
		p.Legend.Add(ylab[i], s)
	}
	p.X.Max = 2.5
	p.Y.Max = 3.5
	p.X.Label.Text = "Time / ps"
	p.Y.Label.Text = "Probability density"

	if err := p.Save(5, 5, "out/"+title+".png"); err != nil {
		panic(err)
	}
}

// createScatter creates a scatter graph from provided x,y data and title
func createScatter(xdat, ydat []float64, title string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	p.Title.Text = title

	scatdata := xyer{xdat, ydat}
	s := plotter.NewScatter(scatdata)
	s.GlyphStyle.Radius = 2
	s.GlyphStyle.Shape = &plot.BoxGlyph{}
	s.GlyphStyle.Color = color.RGBA{G: 100, A: 255}
	p.Add(s)

	if err := p.Save(5, 5, "out/"+title+".png"); err != nil {
		panic(err)
	}
}
