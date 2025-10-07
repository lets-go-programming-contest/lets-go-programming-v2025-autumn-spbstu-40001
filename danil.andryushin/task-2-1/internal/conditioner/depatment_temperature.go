package conditioner

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidEmployeesCount = errors.New("invalid employees count")
	ErrReadWish              = errors.New("read employee wish")
)

const minTemp, maxTemp, wrongTemp int = 15, 30, -1

func CalcDepartmentTemperature(reader io.Reader, writer io.Writer) (*Wish, error) {
	var (
		nEmployees uint
		wish       = Wish{minTemp, maxTemp}
	)

	_, err := fmt.Fscan(reader, &nEmployees)
	if err != nil {
		return nil, ErrInvalidEmployeesCount
	}

	for range nEmployees {
		var (
			currentTemp int
			operator    string
		)

		_, err := fmt.Fscan(reader, &operator, &currentTemp)
		if err != nil {
			return nil, ErrReadWish
		}

		err = fulfilWish(&wish, operator, currentTemp)
		if err != nil {
			return nil, err
		}

		if writer != nil {
			temperature, err := wish.GetTemp()
			if err != nil {
				_, err = fmt.Fprintln(writer, wrongTemp)
			} else {
				_, err = fmt.Fprintln(writer, temperature)
			}

			if err != nil {
				return nil, fmt.Errorf("invalid output stream: %w", err)
			}
		}
	}

	return &wish, nil
}

var ErrUnknownOperator = errors.New("unknown operator")

func fulfilWish(wish *Wish, operator string, currentTemp int) error {
	switch operator {
	case ">=":
		wish.UpdateMinTemp(currentTemp)
	case "<=":
		wish.UpdateMaxTemp(currentTemp)
	default:
		return ErrUnknownOperator
	}

	return nil
}
