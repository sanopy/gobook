package ints

import "errors"

func minElement(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("vals should have element")
	}
	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func maxElement(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("vals should have element")
	}
	max := vals[0]
	for _, val := range vals[1:] {
		if max < val {
			max = val
		}
	}
	return max, nil
}

func min(min int, vals ...int) int {
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max(max int, vals ...int) int {
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}
