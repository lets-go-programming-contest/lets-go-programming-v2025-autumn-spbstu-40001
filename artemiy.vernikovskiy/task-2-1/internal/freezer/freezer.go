package freezer

import (
	"errors"
	"fmt"
	"io"
)

func Calc(reader io.Reader) (int, error) {
	var (
		optimalTemp               int
		border                    string
		minTemp, maxTemp, errTemp = 15, 30, -1
	)

	var (
		ErrInvalidTemperature = errors.New("Invalid temperature")
		ErrWrongOperator      = errors.New("Wrong operator")
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

	return minTemp, nil
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
	var (
		kEmployeesCount          int
		ErrInvalidEmployeesCount = errors.New("Invalid employees count")
	)

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
