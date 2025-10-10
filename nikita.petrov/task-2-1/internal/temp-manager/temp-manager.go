package internal

import "errors"

var ErrSetValue error = errors.New("can't set new value")

type TempManager struct {
	maxTemp int
	minTemp int
}

func (tm *TempManager) Init(maxValue int, minValue int) {
	tm.maxTemp = maxValue
	tm.minTemp = minValue
}

func (tm *TempManager) GetCurrentOptimalTemp() int {
	return tm.minTemp
}

func (tm *TempManager) SetNewOptimalTemp(condition string, newTemp int) error {
	switch condition {
	case ">=":
		if newTemp > tm.minTemp {
			tm.minTemp = newTemp
		}

		if tm.minTemp > tm.maxTemp {
			return ErrSetValue
		}
	case "<=":
		if newTemp < tm.maxTemp {
			tm.maxTemp = newTemp
		}

		if tm.maxTemp < tm.minTemp {
			return ErrSetValue
		}
	}

	return nil
}
