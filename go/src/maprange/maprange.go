package maprange

import (
	"errors"
)

func MapLinear(value float32, fromRangeStart float32, fromRangeEnd float32, toRangeStart float32, toRangeEnd float32) (float32, error) {
	if value < fromRangeStart || value > fromRangeEnd {
		return 0, errors.New("Out of original range")
	}

	/*
		given:
		value = v
		[fromRangeStart, fromRangeEnd] = [a1, a2]
		[toRangeStart, toRangeEnd] = [b1, b2]

		then:
		                   (v - a1)(b2 - b1)
		mappedValue = b1 + -----------------
		                       (a2 - a1)
	*/
	mappedValue := toRangeStart + (((value - fromRangeStart) * (toRangeEnd - toRangeStart)) / (fromRangeEnd - fromRangeStart))
	return mappedValue, nil
}