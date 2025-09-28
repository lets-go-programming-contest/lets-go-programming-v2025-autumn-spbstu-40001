package mainproc;

import "errors";

type Wish struct {
	minTemperature, maxTemperature int;
}
func (wish *Wish) IncludeMin(temperature int) {
	if (wish.minTemperature < temperature) {
		wish.minTemperature = temperature;
	}
}
func (wish *Wish) IncludeMax(temperature int) {
	if (wish.maxTemperature > temperature) {
		wish.maxTemperature = temperature;
	}
}
var ErrNoOptimum = errors.New("optimal temperature does not exist");
func (wish *Wish) GetOptimum() (int, error) {
	if (wish.minTemperature > wish.maxTemperature) {
		return 0, ErrNoOptimum;
	}
	return wish.minTemperature, nil;
}
