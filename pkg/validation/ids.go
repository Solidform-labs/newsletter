package validation

import "strconv"

func ParseNumericID(id string) (bool, int) {
	intID, err := strconv.Atoi(id)
	return err == nil, intID
}
