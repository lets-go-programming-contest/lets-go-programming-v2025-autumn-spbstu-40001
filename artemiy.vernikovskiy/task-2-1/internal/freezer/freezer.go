package freezer

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidTemperature    = errors.New("invalid temperature")
	ErrWrongOperator         = errors.New("wrong operator")
	ErrInvalidEmployeesCount = errors.New("invalid employees count")
)

func Calc(reader io.Reader) (int, error) {
	var (
		optimalTemp               int
		border                    string
		minTemp, maxTemp, errTemp = 15, 30, -1
	)

	_, err := fmt.Fscan(reader, &border, &optimalTemp)
	if err != nil {
		return errTemp, ErrInvalidTemperature
	}

	switch border {
	case ">=":
		minTemp = max(minTemp, optimalTemp)
	case "<=":
		maxTemp = min(maxTemp, optimalTemp)
	default:
		return errTemp, ErrWrongOperator
	}

	if maxTemp < minTemp {
		return errTemp, nil
	}

	return maxTemp, nil
}

func CalcForEmployee(reader io.Reader, kEmployeesCount int) error {
	for range kEmployeesCount {
		resultOfCalc, err := Calc(reader)
		if err != nil {
			return err
		}

		fmt.Println(resultOfCalc)
	}

	return nil
}

func CalcForDepartment(reader io.Reader, nNumberOfOlimpic int) error {
	var kEmployeesCount int

	for range nNumberOfOlimpic {
		_, err := fmt.Fscan(reader, &kEmployeesCount)
		if err != nil || kEmployeesCount < 0 {
			return ErrInvalidEmployeesCount
		}

		err = CalcForEmployee(reader, kEmployeesCount)
		if err != nil {
			return err
		}
	}

	return nil
}
