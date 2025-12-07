package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read piped data from std input
	scanner := bufio.NewScanner(os.Stdin)
	out := ""
	x := 0
	y := 0

	for scanner.Scan() { // Read line by line
		rowCols := scanner.Text()
		out += rowCols + "\n"
		x = len(rowCols)
		y++ // Each line represents a row
	}
	
	if scanner.Err() != nil {
		return
	}

	res := ""
	if QuadA(x, y) == out {
		res += "[quadA] [" + strconv.Itoa(x) + "] [" + strconv.Itoa(y) + "]"
	}

	if QuadB(x, y) == out {
		if len(res) > 0 {
			res += " || "
		}
		res += "[quadB] [" + strconv.Itoa(x) + "] [" + strconv.Itoa(y) + "]"
	}

	if QuadC(x, y) == out {
		if len(res) > 0 {
			res += " || "
		}
		res += "[quadC] [" + strconv.Itoa(x) + "] [" + strconv.Itoa(y) + "]"
	}

	if QuadD(x, y) == out {
		if len(res) > 0 {
			res += " || "
		}
		res += "[quadD] [" + strconv.Itoa(x) + "] [" + strconv.Itoa(y) + "]"
	}

	if QuadE(x, y) == out {
		if len(res) > 0 {
			res += " || "
		}
		res += "[quadE] [" + strconv.Itoa(x) + "] [" + strconv.Itoa(y) + "]"
	}

	if res == "" {
		fmt.Printf("Not a quad function\n")
		return
	}

	res += "\n"
	fmt.Printf("%v", res)
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
