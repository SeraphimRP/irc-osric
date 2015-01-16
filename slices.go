package main

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func removeItemInSlice(a string, list []string) bool {
	for c, b := range list {
		if b == a {
			list[c] = ""
			return true
		}
	}

	return false
}
