package utils

import "strconv"

func ParseFloat32(s string) (float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(f64), nil
}
