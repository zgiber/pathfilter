# Path filter

[![GoDoc](https://godoc.org/github.com/zgiber/pathfilter?status.svg)](https://godoc.org/github.com/zgiber/pathfilter)

## Description

Given a series of points in a .csv file (latitude, longitude, timestamp) describing a journey from A-B, returns a filtered dataset from which the potentionally erroneous points are removed. Use it for cleaning up tracked paths from various applications (sports tracker, courier application, fleet tracker etc.).

It utilises a simple average speed filter, with a static speed limit which is defined in the SI unit meter/sec. Possible improvement would be to weight the path's points as 'potentionally erroneous' and remove a defined upper percentile. That would allow the program to be used with a range of transportation methods with fewer adjustmets and opens up possibilities to be fine-tuned to the path perhaps also considering speed limits on the paths.

The ultimate solution could be to also mark passable / impassable terrains and check poligon boundaries.

## Usage

Compile the binary within ./cmd `go build -o=filter main.go` or run `go install`.

Run the program with the following parameters:
```
  -in string
        path for the input .csv file (default "testpoints.csv")
  -out string
        path for the filtered output .csv file (default "filtered.csv")
  -vmax float
        maximum average speed (points which require speeds above this will be removed) (default 15)
```