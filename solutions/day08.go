package solutions

import "fmt"

var width08 = 25
var height08 = 6

// Day08 ...
func Day08(input string) {
	layerSize := width08 * height08
	numLayers := len(input) / layerSize

	layers := make([]string, numLayers)
	for i := 0; i < numLayers; i++ {
		layers[i] = input[i*layerSize : (i+1)*layerSize]
	}

	for y := 0; y < height08; y++ {
		for x := 0; x < width08; x++ {
			idx := y*width08 + x
			for _, layer := range layers {
				pixel := layer[idx]
				if pixel == '1' {
					fmt.Print("X")
					break
				} else if pixel == '0' {
					fmt.Print(" ")
					break
				}
			}
		}
		fmt.Println("")
	}
}
