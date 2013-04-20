package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func readFile(fname string) (data [][]float64) {
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		println("Cannot read " + fname)
		return
	}

	data = make([][]float64, 0, 6) // Make data table with 6 columns space
	var ci int                     // Column index in data
	num := make([]byte, 0, 32)     // Number buffer, length 32
	var numf float64               // num converted to float64

	// add adds the current buffer to our data list
	// After adding the found float, the column index is increased
	add := func() {
		// Abort if nothing in buffer
		if len(num) == 0 {
			return
		}

		numf, _ = strconv.ParseFloat(string(num), 64) // Convert buffer to float64

		// If column does not exist
		if ci > len(data)-1 {
			data = append(data, make([]float64, 0, 512)) // Create column with starting size of 512
		}

		data[ci] = append(data[ci], numf) // Add float to column
		num = num[:0]                     // Reset buffer
		ci++                              // Increase column number
	}

	// Over all bytes in file
	for _, c := range f {
		switch c {
		// If encountering whitespace, attempt add
		case ' ', ',', '\t':
			add()
		// If line changed, attempt to add and reset column index
		case '\r', '\n':
			add()
			ci = 0
		// If anything else, add character to buffer
		default:
			num = append(num, c)
		}
	}

	return
}

// saveFile saves data as a basic csv
func saveFile(fname string, data [][]float64) {
	lc, lr := len(data), len(data[0])
	dat := make([]byte, 0, lc*lr*5)
	for r := 0; r < lr; r++ {
		for c := 0; c < lc; c++ {
			if c > 0 {
				dat = append(dat, '\t')
			}
			dat = append(dat, []byte(fmt.Sprintf("%f", data[c][r]))...)
		}
		dat = append(dat, '\n')
	}
	ioutil.WriteFile(fname, dat, os.FileMode(os.O_CREATE|os.O_RDWR))
}
