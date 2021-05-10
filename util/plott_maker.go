package util

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func MakerPlot(mainFile string) {

	file, err := os.Open(mainFile)
	if err == io.EOF {
		log.Fatal(err)
	}
	defer file.Close()

	irisDF := dataframe.ReadCSV(file)
	yVals := irisDF.Col("Reviews").Float()
	//fmt.Println(yVals)
	fmt.Println(irisDF.Names())

	for _, colName := range irisDF.Names() {
		switch colName {
		case "Rating", "Installs":
			plots := make(plotter.XYs, irisDF.Nrow())

			for i, floatVal := range irisDF.Col(colName).Float() {

				if !math.IsNaN(yVals[i]) && !math.IsNaN(floatVal) {
					plots[i].X = floatVal
					plots[i].Y = yVals[i]
				}

			}

			// Create a new plot
			p := plot.New()

			p.X.Label.Text = colName
			p.Y.Label.Text = "Reviews"

			p.Add(plotter.NewGrid())

			s, err := plotter.NewScatter(plots)
			if err != nil {
				log.Fatal("error-> ", err)
			}
			s.GlyphStyle.Radius = vg.Points(3)
			p.Add(s)

			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
				log.Fatal("Error creating: ", err)
			}

		default:
			continue
		}
	}

}
