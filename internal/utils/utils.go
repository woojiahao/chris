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

func FuzzyMatch(target string, input string) bool {
	for i := 0; i < len(input); i++ {
		if target[i] != input[i] {
			return false
		}
	}

	return true
}

func Any[V comparable](lst []V, predicate func(V) bool) bool {
	for _, v := range lst {
		if predicate(v) {
			return true
		}
	}

	return false
}

func Map[V comparable, K comparable](lst []V, operation func(V) K) []K {
	var mapped []K
	for _, v := range lst {
		mapped = append(mapped, operation(v))
	}

	return mapped
}

func StrToFloat(str string) *float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil
	}
	return &num
}
