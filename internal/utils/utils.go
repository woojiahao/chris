package utils

import "strconv"

func In[V comparable](lst []V, target V) bool {
	for _, cur := range lst {
		if cur == target {
			return true
		}
	}

	return false
}

func StrToFloat(str string) *float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil
	}
	return &num
}
