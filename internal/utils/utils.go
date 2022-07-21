package utils

func In[V comparable](lst []V, target V) bool {
	for _, cur := range lst {
		if cur == target {
			return true
		}
	}

	return false
}
