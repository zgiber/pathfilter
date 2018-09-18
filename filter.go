package pathfilter

import "math"

// FilterByAvgSpeed takes a path and a speed value (meter/sec)
// and removes the points from the route which would require
// an average speed to reach beyond the specified value
func FilterByAvgSpeed(p *Path, maxSpeedMps float64) {
	filteredPoints := []*point{}
	p.l.RLock()
	for i := 0; i < len(p.points)-1; i++ {
		if avgSpeedMps(p.points[i], p.points[i+1]) <= maxSpeedMps {
			filteredPoints = append(filteredPoints, p.points[i], p.points[i+1])
			continue
		}

		for nextValid := i + 1; nextValid < len(p.points); nextValid++ {
			if avgSpeedMps(p.points[i], p.points[nextValid]) <= maxSpeedMps {
				filteredPoints = append(filteredPoints, p.points[i], p.points[nextValid])
				i = nextValid
			}
		}
	}
	p.l.RUnlock()

	p.l.Lock()
	p.points = filteredPoints
	p.l.Unlock()
}

// returns the average speed between two points in meters / s
func avgSpeedMps(p1, p2 *point) float64 {
	return distance(p1, p2) / float64(p2.Timestamp-p1.Timestamp)
}

// returns the distance between two points in meters
func distance(p1, p2 *point) float64 {
	dLat := toRad(p2.Lat - p1.Lat)
	dLon := toRad(p2.Lon - p1.Lon)
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(toRad(p1.Lat))*math.Cos(toRad(p2.Lat))*math.Pow(math.Sin(dLon/2), 2)
	return 2 * earthRadius * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func toRad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}
