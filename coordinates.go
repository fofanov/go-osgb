package transformation

import (
	"fmt"
	"math"
)

const (
	degreeInRadians = 180 / math.Pi
	radianInDegrees = 1 / degreeInRadians
)

type direction string

const (
	north direction = "N"
	south direction = "S"
	east  direction = "E"
	west  direction = "W"
)

type ETRS89Coordinate struct {
	Lat, Lon, Height float64
}

type OSGB36Coordinate struct {
	Easting, Northing, Height float64
}

func NewETRS89Radians(lat, lon, height float64) *ETRS89Coordinate {
	return &ETRS89Coordinate{
		Lat:    lat,
		Lon:    lon,
		Height: height,
	}
}

func NewOSGB36(easting, northing, height float64) *OSGB36Coordinate {
	return &OSGB36Coordinate{
		Easting:  easting,
		Northing: northing,
		Height:   height,
	}
}

type geographicCoord struct {
	lat, lon, height float64
}

type cartesianCoord struct {
	x, y, z float64
}

type planeCoord struct {
	easting, northing float64
}

func dmsToDecimal(degrees, minutes, seconds float64, direction direction) (float64, error) {
	if direction == "N" || direction == "S" {
		if degrees < 0 || degrees > 90 {
			return 0, fmt.Errorf("invalid latitude degrees %f", degrees)
		}
	} else if direction == "E" || direction == "W" {
		if degrees < 0 || degrees > 180 {
			return 0, fmt.Errorf("invalid longitude degrees %f", degrees)
		}
	} else {
		return 0, fmt.Errorf("invalid direction %s", direction)
	}

	if minutes < 0 || minutes > 60 {
		return 0, fmt.Errorf("invalid minutes %f", minutes)
	}
	if seconds < 0 || seconds > 60 {
		return 0, fmt.Errorf("invalid secondss %f", seconds)
	}

	rad := (degrees + minutes/60 + seconds/3600)
	if direction == "N" || direction == "E" {
		return rad, nil
	}
	return rad * -1, nil
}

func radiansToDegrees(rad float64) float64 {
	return rad * degreeInRadians
}

func degreesToRadians(degrees float64) float64 {
	return degrees * radianInDegrees
}
