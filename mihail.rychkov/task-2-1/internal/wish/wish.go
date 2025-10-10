package wish

import "errors"

var ErrNoOptimum = errors.New("optimal temperature does not exist")

func New(minTemperature, maxTemperature int) Wish {
	return Wish{minTemperature, maxTemperature}
}

type Wish struct {
	minTemperature, maxTemperature int
}

func (wish *Wish) IncludeMin(temperature int) {
	if wish.minTemperature < temperature {
		wish.minTemperature = temperature
	}
}

func (wish *Wish) IncludeMax(temperature int) {
	if wish.maxTemperature > temperature {
		wish.maxTemperature = temperature
	}
}

func (wish *Wish) GetOptimum() (int, error) {
	if wish.minTemperature > wish.maxTemperature {
		return 0, ErrNoOptimum
	}

	return wish.minTemperature, nil
}
