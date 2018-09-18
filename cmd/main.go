package main

import (
	"flag"
	"log"

	"github.com/zgiber/pathfilter"
)

var (
	input    = flag.String("in", "testpoints.csv", "path for the input .csv file")
	output   = flag.String("out", "filtered.csv", "path for the filtered output .csv file")
	maxSpeed = flag.Float64("vmax", 20, "maximum average speed in m/s (points which require speeds above this will be removed)")
)

func main() {
	flag.Parse()
	path, err := pathfilter.NewPathFromCSV(*input)
	if err != nil {
		log.Fatal(err)
	}

	pathfilter.FilterByAvgSpeed(path, *maxSpeed)
	err = path.ExportCSV(*output)
	if err != nil {
		log.Fatal(err)
	}
}
