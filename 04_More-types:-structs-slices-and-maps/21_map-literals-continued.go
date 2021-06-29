package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs": {1.0, 2.0},
	"Google":    {3.0, 4.0},
}

func main() {
	fmt.Println(m)
}
