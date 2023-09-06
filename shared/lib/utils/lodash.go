package utils

func Includes[T comparable](array []T, value T) int {
	if len(array) == 0 {
		return -1
	}
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func CutStrSuffix(str string, n int) string {
	var res []uint8
	for i := 0; i < len(str)-n; i++ {
		res = append(res, str[i])
	}
	return string(res)
}
