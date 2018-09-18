package pathfilter

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

const (
	earthRadius = 6371.0 * 1000 // in meters
)

var (
	errInvalidRecord = errors.New("invalid record")
)

type point struct {
	Lat       float64
	Lon       float64
	Timestamp int64
}

// Path is a collection of points
type Path struct {
	l      sync.RWMutex
	points []*point
}

// NewPathFromCSV returns a path from a CSV file
// the returned *Path maintains the original order
// from the input file.
func NewPathFromCSV(filePath string) (*Path, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var p Path
	p.points = make([]*point, len(records))

	for i, rec := range records {
		if len(rec) != 3 {
			return nil, errors.Wrap(errInvalidRecord, "length must be exactly 3")
		}

		var (
			lat, lon float64
			t        int64
		)

		if lat, err = strconv.ParseFloat(rec[0], 64); err != nil {
			return nil, errors.Wrap(errInvalidRecord, err.Error())
		}

		if lon, err = strconv.ParseFloat(rec[1], 64); err != nil {
			return nil, errors.Wrap(errInvalidRecord, err.Error())
		}

		if t, err = strconv.ParseInt(rec[2], 10, 64); err != nil {
			return nil, errors.Wrap(errInvalidRecord, err.Error())
		}

		p.points[i] = &point{
			Lat:       lat,
			Lon:       lon,
			Timestamp: t,
		}
	}
	return &p, nil
}

// ExportCSV saves a CSV version of the path
func (p *Path) ExportCSV(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	for _, point := range p.points {
		csvWriter.Write(
			[]string{
				strconv.FormatFloat(point.Lat, 'f', 9, 64),
				strconv.FormatFloat(point.Lon, 'f', 9, 64),
				strconv.FormatInt(point.Timestamp, 10),
			})
	}
	csvWriter.Flush()
	return csvWriter.Error()
}
