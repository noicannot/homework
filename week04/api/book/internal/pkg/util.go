package pkg

import "strconv"

func StringToInt(s string) int {
	if x, err := strconv.Atoi(s); err != nil {
		return -1
	} else {
		return x
	}
}
