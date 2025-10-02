package conditioner

import (
	"errors"
	"fmt"
	"io"
)

var ErrInvalidEmployeesCount = errors.New("invalid employees count")

func MakeDepartmentTemperature(reader io.Reader, writer io.Writer) (int, error) {
	var (
		nEmployees  uint
		conditioner = Conditioner{
			minTemp: 15,
			maxTemp: 30,
		}
		temperature int
	)

	_, err := fmt.Fscan(reader, &nEmployees)
	if err != nil {
		return 0, ErrInvalidEmployeesCount
	}

	for range nEmployees {
		temperature, err = fulfilWish(&conditioner, reader)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Fprintln(writer, temperature)
		if err != nil {
			return 0, fmt.Errorf("invalid output stream: %w", err)
		}
	}

	return temperature, nil
}

var (
	ErrReadWish        = errors.New("failed to read employee wish")
	ErrUnknownOperator = errors.New("unknown operator")
)

func fulfilWish(conditioner *Conditioner, reader io.Reader) (int, error) {
	var (
		currentTemp int
		operator    string
	)

	_, err := fmt.Fscan(reader, &operator, &currentTemp)
	if err != nil {
		return 0, ErrReadWish
	}

	switch operator {
	case ">=":
		conditioner.UpdateMinTemp(currentTemp)
	case "<=":
		conditioner.UpdateMaxTemp(currentTemp)
	default:
		return 0, ErrUnknownOperator
	}

	temp, err := conditioner.GetTemp()
	if err != nil {
		temp = -1
	}

	return temp, nil
}
