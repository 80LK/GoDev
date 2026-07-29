package utils

func Map[T any, R any](slice []T, callback func(T, int, []T) (R, error)) ([]R, error) {
	ret := []R{}
	for i, item := range slice {
		r, err := callback(item, i, slice)
		if err != nil {
			return nil, err
		}
		ret = append(ret, r)
	}
	return ret, nil
}
