package isbn

import (
	"errors"
)

func IsValidISBN10(isbn string) (bool, error) {
	if len(isbn) != 10 {
		return false, errors.New("Bad formatted ISBN")
	}

	sum := 0
	for idx, c := range isbn {
		num := int(c) - '0'
		if num < 0 || num > 9 {
			return false, errors.New("ISBN contains at least one non-numerical character")
		}

		sum += (10 - idx) * num
	}

	return sum % 11 == 0, nil
}