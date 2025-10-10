package internal

import "errors"

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
	case "<=":
		if newTemp >= TM.minTemp {
			if newTemp < TM.maxTemp {
				TM.maxTemp = newTemp
			}
		} else {
			return errors.New("can't set new value")
		}
	case ">=":
		if newTemp <= TM.maxTemp {
			if newTemp > TM.minTemp {
				TM.minTemp = newTemp
				TM.optTemp = TM.minTemp
			}
		} else {
			return errors.New("can't set new value")
		}
	}

	return nil
}
