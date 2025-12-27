package commander

import (
	"errors"
	"fmt"

	"github.com/Aapng-cmd/task-2-1/internal/freezer"
	"github.com/Aapng-cmd/task-2-1/internal/models"
)

var (
	ErrInvalidTemperature    = errors.New("invalid temperature")
	ErrInvalidEmployeesCount = errors.New("invalid employees count")
	ErrWTF                   = errors.New("smth strange and wicked has benn done")
)

func CalcForDepartment(wrapper *models.Wrapper) error {
	var nNumberOfDepartures int

	_, err := fmt.Fscan(wrapper.Cin, &nNumberOfDepartures)
	if err != nil || nNumberOfDepartures < 0 {
		return fmt.Errorf("%w", err)
	}

	var kEmployeesCount int

	for range nNumberOfDepartures {
		_, err := fmt.Fscan(wrapper.Cin, &kEmployeesCount)
		if err != nil || kEmployeesCount < 0 {
			return ErrInvalidEmployeesCount
		}

		err = CalcForEmployee(wrapper, kEmployeesCount)
		if err != nil {
			return err
		}
	}

	return nil
}

func CalcForEmployee(wrapper *models.Wrapper, kEmployeesCount int) error {
	var (
		optimalTemp      int
		border           string
		minTemp, maxTemp = 15, 30
	)

	for range kEmployeesCount {
		_, err := fmt.Fscan(wrapper.Cin, &border, &optimalTemp)
		if err != nil {
			return ErrInvalidTemperature
		}

		finalTemp, err := freezer.CalcForEmployee(&minTemp, &maxTemp, optimalTemp, border)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		_, err = fmt.Fprintf(wrapper.Cout, "%d\n", finalTemp)
		if err != nil {
			return ErrWTF
		}
	}

	return nil
}
