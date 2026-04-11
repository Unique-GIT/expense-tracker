package utils

import (
	"fmt"
	"strconv"
)

func ParseFloat32(s string) (float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(f64), nil
}

func StringToInt32(s string) (int32, error) {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid int16 value: %w", err)
	}
	return int32(val), nil
}
