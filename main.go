package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	width := -1
	height := 0

	// Read input and validate row widths
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" && len(lines) == 0 {
			continue // skip empty line at beginning
		}
		if line == "" && scanner.Err() == nil && !scanner.Scan() {
			break // ignore the empty line at the end of the file
		}
		if width == -1 {
			width = len(line)
		} else if len(line) != width { // Ensure all widths are same
			fmt.Println("Not a quad function")
			return
		}

		lines = append(lines, line)
		height++
	}

	if scanner.Err() != nil {
		fmt.Println("Not a quad function")
		return
	}

	// Reject if no rows or no columns
	if width <= 0 || height <= 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Join lines with newline to match quad output format
	input := strings.Join(lines, "\n")

	matches := []string{}

	if compareQuad(input, QuadA(width, height)) {
		matches = append(matches, "[quadA] ["+strconv.Itoa(width)+"] ["+strconv.Itoa(height)+"]")
	}
	if compareQuad(input, QuadB(width, height)) {
		matches = append(matches, "[quadB] ["+strconv.Itoa(width)+"] ["+strconv.Itoa(height)+"]")
	}
	if compareQuad(input, QuadC(width, height)) {
		matches = append(matches, "[quadC] ["+strconv.Itoa(width)+"] ["+strconv.Itoa(height)+"]")
	}
	if compareQuad(input, QuadD(width, height)) {
		matches = append(matches, "[quadD] ["+strconv.Itoa(width)+"] ["+strconv.Itoa(height)+"]")
	}
	if compareQuad(input, QuadE(width, height)) {
		matches = append(matches, "[quadE] ["+strconv.Itoa(width)+"] ["+strconv.Itoa(height)+"]")
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(matches, " || "))
}

func compareQuad(input, quad string) bool { // Ensure for all \n characters in last line get removed for both strings
	return strings.TrimRight(input, "\n") == strings.TrimRight(quad, "\n")
}

func QuadA(x, y int) string {
	res := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 { // Top left corner
				res += "o"
			} else if row == 1 && col == x { // Top right corner
				res += "o"
			} else if row == y && col == 1 { // Bottom left corner
				res += "o"
			} else if row == y && col == x { // Bottom right corner
				res += "o"
			} else if row == 1 || row == y { // bottom and top border
				res += "-"
			} else if col == 1 || col == x { // side borders
				res += "|"
			} else { // empty insides
				res += " "
			}
		}
		res += "\n" // new line after row
	}
	return res
}

func QuadB(x, y int) string {
	res := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 { // Top left corner
				res += "/"
			} else if row == 1 && col == x { // Top right corner
				res += "\\"
			} else if row == y && col == 1 { // Bottom left corner
				res += "\\"
			} else if row == y && col == x { // Bottom right corner
				res += "/"
			} else if row == 1 || row == y { // bottom and top border
				res += "*"
			} else if col == 1 || col == x { // side borders
				res += "*"
			} else { // empty insides
				res += " "
			}
		}
		res += "\n" // new line after row
	}
	return res
}

func QuadC(x, y int) string {
	res := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 { // Top left corner
				res += "A"
			} else if row == 1 && col == x { // Top right corner
				res += "A"
			} else if row == y && col == 1 { // Bottom left corner
				res += "C"
			} else if row == y && col == x { // Bottom right corner
				res += "C"
			} else if row == 1 || row == y { // bottom and top border
				res += "B"
			} else if col == 1 || col == x { // side borders
				res += "B"
			} else { // empty insides
				res += " "
			}
		}
		res += "\n" // new line after row
	}
	return res
}

func QuadD(x, y int) string {
	res := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 { // Top left corner
				res += "A"
			} else if row == 1 && col == x { // Top right corner
				res += "C"
			} else if row == y && col == 1 { // Bottom left corner
				res += "A"
			} else if row == y && col == x { // Bottom right corner
				res += "C"
			} else if row == 1 || row == y { // bottom and top border
				res += "B"
			} else if col == 1 || col == x { // side borders
				res += "B"
			} else { // empty insides
				res += " "
			}
		}
		res += "\n" // new line after row
	}
	return res
}

func QuadE(x, y int) string {
	res := ""
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 { // Top left corner
				res += "A"
			} else if row == 1 && col == x { // Top right corner
				res += "C"
			} else if row == y && col == 1 { // Bottom left corner
				res += "C"
			} else if row == y && col == x { // Bottom right corner
				res += "A"
			} else if row == 1 || row == y { // bottom and top border
				res += "B"
			} else if col == 1 || col == x { // side borders
				res += "B"
			} else { // empty insides
				res += " "
			}
		}
		res += "\n" // new line after row
	}
	return res
}
