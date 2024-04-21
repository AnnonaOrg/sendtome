package common

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must2[T any](v T, err error) T {
	Must(err)
	return v
}

func Contains[T comparable](v T, s []T) bool {
	for _, t := range s {
		if v == t {
			return true
		}
	}
	return false
}
