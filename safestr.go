package safestr

import "fmt"

func Parse(value string, replacement string, visible int) (*string, error) {
	if visible <= 0 {
		visible = 4
	}

	vLength := len(value)

	if vLength < visible {
		return nil, fmt.Errorf("the length of the value must be greater than the visible part")
	}

	chars := []rune(value)

	out := ""
	for i := 0; i < vLength; i++ {
		if i >= (vLength - visible) {
			out += string(chars[i])
			continue
		}
		out += replacement
	}

	return &out, nil
}