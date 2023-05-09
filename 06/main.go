package main

import (
	"fmt"
	"strings"
)

func main() {

}

func convert(s string, numRows int) string {
	lines := make([]string, numRows)
	size := len(s)
	pos := 0
	for pos < size {
		for i := 0; i < numRows; i++ {
			lines[i] = fmt.Sprintf("%s%s", lines[i], string(s[pos]))
			pos++
			if pos >= size {
				break
			}
		}
		if pos >= size {
			break
		}
		for j := numRows - 2; j > 0; j-- {
			lines[j] = fmt.Sprintf("%s%s", lines[j], string(s[pos]))
			pos++
			if pos >= size {
				break
			}
		}
	}
	builder := strings.Builder{}
	for i := 0; i < numRows; i++ {
		builder.WriteString(lines[i])
	}
	return builder.String()
}
