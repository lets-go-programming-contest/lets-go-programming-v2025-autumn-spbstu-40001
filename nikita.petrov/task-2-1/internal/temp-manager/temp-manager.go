package internal

import "errors"

var SetValueError error = errors.New("can't set new value")

type TempManager struct {
	maxTemp int
	minTemp int
	optTemp int
}

func (TM *TempManager) Init(maxValue int, minValue int) {
	TM.maxTemp = maxValue
	TM.minTemp = minValue
	TM.optTemp = minValue
}

func (TM *TempManager) GetCurrentOptimalTemp() int {
	return TM.optTemp
}

func (TM *TempManager) SetNewOptimalTemp(condition string, newTemp int) error {
	switch condition {
	case ">=":
		if newTemp > TM.minTemp {
			TM.minTemp = newTemp
		}
		if TM.minTemp > TM.maxTemp {
			return SetValueError
		}
		TM.optTemp = TM.minTemp
	case "<=":
		if newTemp < TM.maxTemp {
			TM.maxTemp = newTemp
		}
		if TM.maxTemp < TM.minTemp {
			return SetValueError
		}
	}
	return nil
}
