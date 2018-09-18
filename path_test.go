package pathfilter

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadValidPath(t *testing.T) {
	// quick check if parsing the input is OK
	p, err := NewPathFromCSV("testpoints.csv")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 227, len(p.points))
	assert.Equal(t, 51.49871493, p.points[0].Lat)
	assert.Equal(t, -0.160117799, p.points[0].Lon)
	assert.Equal(t, int64(1326378718), p.points[0].Timestamp)
}

func TestPathDistances(t *testing.T) {
	p, err := NewPathFromCSV("testpoints.csv")
	if err != nil {
		t.Fatal(err)
	}

	// eyeball if the values look correct (best effort without test result set)
	for i := 0; i < len(p.points)-1; i++ {
		d := distance(p.points[i], p.points[i+1])
		td := float64(p.points[i+1].Timestamp - p.points[i].Timestamp)
		fmt.Printf("distance: %dm\nspeed: %dkm/h\n-----\n", int(d), int(3.6*d/td))
	}
}

func TestFiltered(t *testing.T) {
	p, err := NewPathFromCSV("testpoints.csv")
	if err != nil {
		t.Fatal(err)
	}

	// eyeball if the numbers look correct (best effort without test result set)
	// note, speed here is in m/s where 20 is ~72km/h which is a very generous
	// limit inside London
	FilterByAvgSpeed(p, 20)

	for _, point := range p.points {
		// printing a list of points to be plotted
		// using: http://www.hamstermap.com/quickmap.php
		fmt.Printf("%f,%f\n", point.Lat, point.Lon)
	}
}

func TestExportCSV(t *testing.T) {
	expected, err := NewPathFromCSV("testpoints.csv")
	if err != nil {
		t.Fatal(err)
	}

	err = expected.ExportCSV("testpoints_backup.csv")
	if err != nil {
		t.Fatal(err)
	}

	got, err := NewPathFromCSV("testpoints_backup.csv")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(expected.points), len(got.points))
	for i := 0; i < len(expected.points); i++ {
		assert.Equal(t, expected.points[i].Lat, got.points[i].Lat, fmt.Sprintf("[%d] Latitude does not match", i))
		assert.Equal(t, expected.points[i].Lon, got.points[i].Lon, fmt.Sprintf("[%d] Longitude does not match", i))
		assert.Equal(t, expected.points[i].Timestamp, got.points[i].Timestamp, fmt.Sprintf("[%d] Timestamp does not match", i))
	}

	// try to clean up
	os.Remove("testpoints_backup.csv")
}
