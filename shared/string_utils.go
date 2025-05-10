package shared

import "strconv"

func ParseStringToInt(number string) (int, error) {
	value, err := strconv.Atoi(number)
	if err != nil {
		return 0, err
	}
	return value, nil
}
