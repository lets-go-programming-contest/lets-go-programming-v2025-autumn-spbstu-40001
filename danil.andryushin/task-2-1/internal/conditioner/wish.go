package conditioner

import "errors"

type Conditioner struct {
	minTemp, maxTemp int
}

func (cond *Conditioner) UpdateMaxTemp(temp int) {
	cond.maxTemp = min(cond.maxTemp, temp)
}

func (cond *Conditioner) UpdateMinTemp(temp int) {
	cond.minTemp = max(cond.minTemp, temp)
}

var ErrInvalidTemp = errors.New("invalid temperature")

func (cond *Conditioner) GetTemp() (int, error) {
	if cond.maxTemp >= cond.minTemp {
		return cond.minTemp, nil
	}

	return 0, ErrInvalidTemp
}
