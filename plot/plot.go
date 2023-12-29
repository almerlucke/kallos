package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Plottable interface for Kallos
type Plottable interface {
	ToXYs() plotter.XYs
}

// Plot values
func Plot(p Plottable, resultPath string, width int, height int) error {
	data := p.ToXYs()

	pl := plot.New()

	pl.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(data)
	if err != nil {
		return err
	}

	s.GlyphStyle.Radius = vg.Points(1)

	pl.Add(s)

	return pl.Save(vg.Length(width), vg.Length(height), resultPath)
}
