package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create a slice of slices, specifically, a slice of []uint8 with dy length
	// This means we have a 2D array of dy rows, we have the Y axis of the image
	// Each element of this outer slice is initially nil, meaning we don't have the X axis of the image yet
	img := make([][]uint8, dy)
	
	// Loop over each row, top to bottom since we are going over dy
	// The outer loops controls which row we are filling
	for y := 0; y < dy; y++ {
		// Now we create the X axis for each row
		row := make([]uint8, dx) // Create a new row of pixels
		// The inner loop controls which column we are filling
		// So we go 
		// [0][0], [0][1], [0][2], until [0][dx-1], 
		// [1][0], [1][1], [1][2], until [0][dx-1], 
		// [2][0], [2][1], [2][2], until [0][dx-1], 
		// until we reach [dy-1][dx-1]
		for x := 0; x < dx; x++ { // For each column, calculate a number based on x ^ y
			row[x] = uint8((x ^y)) // Run a function on each element
		} 
		// Finally, we say that the []uint8 that should be in this row (Y axis), is this set of columns (X axis)
		img[y] = row // Add the row to the slice
	}
	
	return img // Return the slice
}

func picSlices() {
	// pic.Show(Pic) will intrepet the data as an image and display it
	pic.Show(Pic) // Call the function
}
