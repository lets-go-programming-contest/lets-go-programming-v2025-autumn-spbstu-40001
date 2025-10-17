package controller

import "errors"

var (
	ErrUnknownOperator = errors.New("unknown operator")
	ErrInvalidTemp     = errors.New("invalid temperature")
)

func New(minTemperature, maxTemperature int) TemperatureController {
	return TemperatureController{minTemperature, maxTemperature}
}

type TemperatureController struct {
	minTemperature, maxTemperature int
}

type Desire struct {
	Sign               string
	DesiredTemperature int
}

func (obj *TemperatureController) ChangeTemperature(desire Desire) error {
	switch desire.Sign {
	case ">=":
		obj.updateMinTemperature(desire.DesiredTemperature)
	case "<=":
		obj.updateMaxTemperature(desire.DesiredTemperature)
	default:
		return ErrUnknownOperator
	}

	return nil
}

func (obj *TemperatureController) GetTemperature() (int, error) {
	if obj.maxTemperature >= obj.minTemperature {
		return obj.minTemperature, nil
	}

	return 0, ErrInvalidTemp
}

func (obj *TemperatureController) updateMaxTemperature(temp int) {
	obj.maxTemperature = min(obj.maxTemperature, temp)
}

func (obj *TemperatureController) updateMinTemperature(temp int) {
	obj.minTemperature = max(obj.minTemperature, temp)
}
