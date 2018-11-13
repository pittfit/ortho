package ast

import (
	"strconv"
)

func stringsForNumericRange(startBuf []byte, endBuf []byte, stepBuf []byte) ([]string, error) {
	numbers, err := numbersForRange(startBuf, endBuf, stepBuf)
	if err != nil {
		return nil, err
	}

	var strings []string

	for _, n := range numbers {
		strings = append(strings, strconv.Itoa(n))
	}

	return strings, nil
}

func numbersForRange(startBuf []byte, endBuf []byte, stepBuf []byte) ([]int, error) {
	var err error
	var start, end, step int

	start, err = strconv.Atoi(string(startBuf))
	if err != nil {
		return nil, err
	}

	end, err = strconv.Atoi(string(endBuf))
	if err != nil {
		return nil, err
	}

	step = 0

	if len(stepBuf) > 0 {
		step, err = strconv.Atoi(string(stepBuf))
		if err != nil {
			return nil, err
		}
	}

	return numbers(start, end, step), nil
}

func numbers(start int, end int, step int) []int {
	if start == end {
		return []int{start}
	}

	if step == 0 {
		step = 1
	} else if step < 0 {
		start, end = end, start
		step = -step
	}

	var nums []int

	if start < end {
		for n := start; n <= end; n += step {
			nums = append(nums, n)
		}
	} else {
		for n := start; n >= end; n -= step {
			nums = append(nums, n)
		}
	}

	return nums
}
