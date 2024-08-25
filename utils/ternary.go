package utils

func TernaryInt(condition bool, a int, b int) int {
	if condition { return a }
	return b
}

func TernaryString(condition bool, a string, b string) string {
	if condition { return a }
	return b
}