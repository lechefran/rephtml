package rephtml

import "regexp"

/*
Internal parsing function to remove all spaces from a byte array
*/
func strip(bytes []byte) []byte {
	re := regexp.MustCompile("\\s+")
	return re.ReplaceAll(bytes, nil)
}

/*
Internal function that creates tabs based on the value of t
*/
func tabs(t int) string {
	res := ""
	for i := 0; i < t; i++ {
		res += tab
	}
	return res
}
