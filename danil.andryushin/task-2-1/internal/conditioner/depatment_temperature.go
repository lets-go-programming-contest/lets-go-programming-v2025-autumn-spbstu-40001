package conditioner

import (
	"errors"
	"fmt"
	"io"
)

var ErrInvalidEmployeesCount = errors.New("invalid employees count")

func CalcDepartmentTemperature(reader io.Reader, writer io.Writer) (int, error) {
	var (
		nEmployees  uint
		wish        = Wish{15, 30}
		temperature int
	)

	_, err := fmt.Fscan(reader, &nEmployees)
	if err != nil {
		return 0, ErrInvalidEmployeesCount
	}

	for range nEmployees {
		temperature, err = fulfilWish(&wish, reader)
		if err != nil {
			return 0, err
		}

		if writer != nil {
			_, err = fmt.Fprintln(writer, temperature)
			if err != nil {
				return 0, fmt.Errorf("invalid output stream: %w", err)
			}
		}
	}

	return temperature, nil
}

var (
	ErrReadWish        = errors.New("failed to read employee wish")
	ErrUnknownOperator = errors.New("unknown operator")
)

func fulfilWish(wish *Wish, reader io.Reader) (int, error) {
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
		wish.UpdateMinTemp(currentTemp)
	case "<=":
		wish.UpdateMaxTemp(currentTemp)
	default:
		return 0, ErrUnknownOperator
	}

	temp, err := wish.GetTemp()
	if err != nil {
		temp = -1
	}

	return temp, nil
}
