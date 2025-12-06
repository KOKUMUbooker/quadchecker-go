package main

func main() {
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
