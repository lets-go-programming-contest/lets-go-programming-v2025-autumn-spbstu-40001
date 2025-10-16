package tmanager

type TempManager struct {
	maxTemp int
	minTemp int
}

func New(maxValue int, minValue int) TempManager {
	return TempManager{maxValue, minValue}
}

func (tm *TempManager) SetAndGetNewOptimalTemp(condition string, newTemp int) int {
	switch condition {
	case ">=":
		if newTemp > tm.minTemp {
			tm.minTemp = newTemp
		}

		if tm.minTemp > tm.maxTemp {
			return -1
		}
	case "<=":
		if newTemp < tm.maxTemp {
			tm.maxTemp = newTemp
		}

		if tm.maxTemp < tm.minTemp {
			return -1
		}
	default:
		return -1
	}

	return tm.minTemp
}
