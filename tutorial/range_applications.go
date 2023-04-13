package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Location struct {
	Lat, Long float64
}

var city map[string]Location

func main() {
	// struct array - range
	var vertex_arr [4]Vertex

	vertex_arr[0] = Vertex{0, 0}
	vertex_arr[1] = Vertex{1, 0}
	vertex_arr[2] = Vertex{0, 1}
	vertex_arr[3] = Vertex{1, 1}

	for i, v := range vertex_arr {
		// fmt.Printf("Vertex %d : %v\n", i, v)
		fmt.Printf("Vertex %d : ", i)
		fmt.Println(v)
	}

	// map - range
	city = make(map[string]Location)
	city["Seoul"] = Location{37.532600, 127.024612}
	city["Daejeon"] = Location{36.351002, 127.385002}
	city["Gwangju"] = Location{35.166668, 126.916664}
	city["Busan"] = Location{35.166668, 129.066666}

	for i, v := range city {
		fmt.Printf("key: %s, ", i)
		fmt.Printf("value: %v \n", v)
	}
}
