package main

import (
	"os"
	"os/exec"
)

func splitLines(s string) []string {
	var out []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		out = append(out, s[start:])
	}
	return out
}

func compareLines(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equal(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	strNum := ""
	for n > 0 {
		digit := n % 10
		s = string('0'+digit) + s // prepend the digit
		n /= 10
	}

	return sign + strNum
}

func main() {
	//chk output of given
	input, _ := os.ReadFile(os.Stdin.Name())
	if len(input) == 0 {
		os.Stdout.WriteString("Not a quad function")
		return
	}
	lines := splitLines(string(input))
	h := len(lines)
	if h == 0 {
		os.Stdout.WriteString("Not a quad function")
		return
	}
	w := len(lines[0])
	for _, l := range lines {
		if len(l) != w {
			os.Stdout.WriteString("Not a quad function")
			return
		}
	}
//slices to for testing matches....names for output matches...matches slice for storing matches
	quads := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	names := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	matches := []string{}

	for i := 0; i < len(quads); i++ {
		//exec.Command takes quad name and x,y as parameters stored in cmd
		cmd := exec.Command(quads[i], itoa(w), itoa(h))
		//comd.Output stores output in out
		out, err := cmd.Output()
		if err != nil {
			continue
		}
		//split lines of quad and store for comparison
		genLines := splitLines(string(out))
		if compareLines(lines, genLines) {//check which quad matches the given quad at stdin and store in matches slice
			matches = append(matches, "["+names[i]+"] ["+itoa(w)+"] ["+itoa(h)+"]")
		}
	}

	if len(matches) == 0 {//no matches, not a quad function
		os.Stdout.WriteString("Not a quad function")
		return
	}

	for i := 0; i < len(matches); i++ {//add || to out put
		if i > 0 {
			os.Stdout.WriteString(" || ")
		}
		os.Stdout.WriteString(matches[i])
	}
}
