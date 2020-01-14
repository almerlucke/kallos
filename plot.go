package kallos

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Plottable interface for Kallos
type Plottable interface {
	ToXYs() plotter.XYs
}

// ToXYs returns a plotter XY's slice
func (vls Values) ToXYs() plotter.XYs {
	pts := make(plotter.XYs, len(vls))

	for i, v := range vls {
		pts[i] = plotter.XY{
			X: float64(i + 1),
			Y: v[0],
		}
	}

	return pts
}

// Plot values
func Plot(p Plottable, resultPath string, width int, height int) error {
	data := p.ToXYs()

	pl, err := plot.New()
	if err != nil {
		return err
	}

	pl.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(data)
	if err != nil {
		return err
	}

	s.GlyphStyle.Radius = vg.Points(1)

	pl.Add(s)

	err = pl.Save(vg.Length(width), vg.Length(height), resultPath)
	return err
}
