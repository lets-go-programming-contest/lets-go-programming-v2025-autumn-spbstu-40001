package tmanager

import "errors"

var ErrSetNewOptTemp error = errors.New("cannot set new optimal temp")

type TempManager struct {
	maxTemp int
	minTemp int
}

func New(maxValue int, minValue int) TempManager {
	return TempManager{maxValue, minValue}
}

func (tm *TempManager) SetAndGetNewOptimalTemp(condition string, newTemp int) (int, error) {
	switch condition {
	case ">=":
		if newTemp > tm.minTemp {
			tm.minTemp = newTemp
		}

		if tm.minTemp > tm.maxTemp {
			return -1, nil
		}
	case "<=":
		if newTemp < tm.maxTemp {
			tm.maxTemp = newTemp
		}

		if tm.maxTemp < tm.minTemp {
			return -1, nil
		}
	default:
		return 0, ErrSetNewOptTemp
	}

	return tm.minTemp, nil
}
