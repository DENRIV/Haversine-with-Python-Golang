// haversine
// go get "github.com/umahmood/haversine"
// go run main.go

package main

import (
    "fmt"

    "github.com/umahmood/haversine"
)

func main() {
    oxford := haversine.Coord{Lat: 51.45, Lon: 1.15}  // Oxford, UK
    turin  := haversine.Coord{Lat: 45.04, Lon: 7.42}  // Turin, Italy
    mi, km := haversine.Distance(oxford, turin)
	fmt.Println("From Oxford, UK")
	fmt.Println("To Turin, Italy")
    fmt.Println("Miles:", mi, "Kilometers:", km)
}