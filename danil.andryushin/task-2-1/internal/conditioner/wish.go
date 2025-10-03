package conditioner

import "errors"

type Wish struct {
	minTemp, maxTemp int
}

func (wish *Wish) UpdateMaxTemp(temp int) {
	wish.maxTemp = min(wish.maxTemp, temp)
}

func (wish *Wish) UpdateMinTemp(temp int) {
	wish.minTemp = max(wish.minTemp, temp)
}

var ErrInvalidTemp = errors.New("invalid temperature")

func (wish *Wish) GetTemp() (int, error) {
	if wish.maxTemp >= wish.minTemp {
		return wish.minTemp, nil
	}

	return 0, ErrInvalidTemp
}
