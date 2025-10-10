package freezer

import (
	"errors"
)

var ErrWrongOperator = errors.New("wrong operator")

func CalcForEmployee(minTemp *int, maxTemp *int, optimalTemp int, border string) (int, error) {
	errTemp := -1

	switch border {
	case ">=":
		*minTemp = max(*minTemp, optimalTemp)
	case "<=":
		*maxTemp = min(*maxTemp, optimalTemp)
	default:
		return 0, ErrWrongOperator
	}

	if *maxTemp < *minTemp {
		return errTemp, nil
	} else {
		return *minTemp, nil
	}
}
