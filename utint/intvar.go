package utint

import "strconv"

func ConvertFromString(str string, def int) int {
	newInt, err := strconv.Atoi(str)
	if err != nil {
		return def
	}

	return newInt
}

func Convert64FromString(str string, def int64) int64 {
	newInt, err := strconv.Atoi(str)
	if err != nil {
		return def
	}

	return int64(newInt)
}
