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

func CalcForEmployee(reader io.Reader, kEmployeesCount int) error {
	var (
		optimalTemp               int
		border                    string
		minTemp, maxTemp, errTemp = 15, 30, -1
	)

	for range kEmployeesCount {
		_, err := fmt.Fscan(reader, &border, &optimalTemp)
		if err != nil {
			return ErrInvalidTemperature
		}

		switch border {
		case ">=":
			minTemp = max(minTemp, optimalTemp)
		case "<=":
			maxTemp = min(maxTemp, optimalTemp)
		default:
			return ErrWrongOperator
		}

		if maxTemp < minTemp {
			fmt.Println(errTemp)
		} else {
			fmt.Println(minTemp)
		}
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
