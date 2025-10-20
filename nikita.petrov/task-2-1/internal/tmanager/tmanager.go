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
	case "<=":
		if newTemp < tm.maxTemp {
			tm.maxTemp = newTemp
		}
	default:
		return 0, ErrSetNewOptTemp
	}

	if tm.maxTemp < tm.minTemp {
		return -1, nil
	}

	return tm.minTemp, nil
}
