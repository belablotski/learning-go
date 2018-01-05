// From https://github.com/gonum/plot/blob/master/plotter/timeseries_test.go

// To re-define ticks UTC time to PST
// plot.UTCUnixTime = plot.UnixTimeIn(time.Local)		

package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	rnd := rand.New(rand.NewSource(1))
	_ = rnd

	// xticks defines how we convert and display time.Time values.
	xticks := plot.TimeTicks{Format: "2006-01-02"}
	yticks := plot.TimeTicks{Format: "15:04"}

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)
		for i := range pts {
			dt := time.Date(2018, 1, i+1, (i * 3) % 5, (i * 31) % 60, 0, 0, time.UTC)
            truncatedDate := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
            truncatedTime := time.Date(1, 1, 1, dt.Hour(), dt.Minute(), 0, 0, dt.Location())
			pts[i].X = float64(truncatedDate.Unix())
			pts[i].Y = float64(truncatedTime.Unix())
		}
		return pts
	}

	n := 10
	data := randomPoints(n)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}
	p.Title.Text = "Date/Time"
	p.X.Label.Text = "Date"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Time"
	p.Y.Tick.Marker = yticks
	p.Add(plotter.NewGrid())

	line, points, err := plotter.NewLinePoints(data)
	if err != nil {
		log.Panic(err)
	}
	line.Color = color.RGBA{G: 255, A: 255}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{R: 255, A: 255}

	p.Add(line, points)

	err = p.Save(25*vg.Centimeter, 15*vg.Centimeter, "06-time_series_xy.png")
	if err != nil {
		log.Panic(err)
	}

	err = p.Save(25*vg.Centimeter, 15*vg.Centimeter, "06-time_series_xy.svg")
	if err != nil {
		log.Panic(err)
	}
}