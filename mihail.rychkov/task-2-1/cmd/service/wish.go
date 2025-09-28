import "errors"

type Wish struct {
	minTemperature, maxTemperature int;
}
func (wish *Wish) includeMin(temperature int) {
	if (wish.minTemperature < temperature) {
		wish.minTemperature = temperature;
	}
}
func (wish *Wish) includeMax(temperature int) {
	if (wish.maxTemperature > temperature) {
		wish.minTemperature = temperature;
	}
}
var ErrNoOptimum = errors.New("optimal temperature does not exist");
func (wish *Wish) getOptimum() (int, error) {
	if (wish.minTemperature > wish.maxTemperature) {
		return 0, ErrNoOptimum;
	}
	return wish.minTemperature, nil;
}
