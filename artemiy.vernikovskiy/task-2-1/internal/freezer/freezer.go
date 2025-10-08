package freezer

import (
	"errors"
	"fmt"
	"io"
)

var minTemp, maxTemp, errTemp = 15, 30, -1

var (
	optimalTemp int
	border      string
)

func Calc(reader io.Reader) (int, error) {
	_, err := fmt.Fscan(reader, &border, &optimalTemp)
	if err != nil {
		return errTemp, errors.New("Invalid temperature")
	}

	switch border {
	case ">=":
		minTemp = max(minTemp, optimalTemp)
	case "<=":
		maxTemp = min(maxTemp, optimalTemp)
	default:
		return errTemp, errors.New("Wrong operator")
	}

	if maxTemp < minTemp {
		return errTemp, nil
	}

	return minTemp, nil
}

func CalcForEmployee(reader io.Reader, kEmployeesCount int) error {
	// var resultOfCalc int
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
			return errors.New("Invalid employees count")
		}

		err = CalcForEmployee(reader, kEmployeesCount)
		if err != nil {
			return err
		}
	}

	return nil
}
