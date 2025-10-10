package main

import (
	"errors"
	"fmt"

	"github.com/Rychmick/task-2-1/internal/wish"
)

var ErrUnknownWishSign = errors.New("unknown comparison sign")

const (
	minSupportedTemperature = 15
	maxSupportedTemperature = 30
)

func main() {
	var nDepartments uint

	_, err := fmt.Scan(&nDepartments)
	if err != nil {
		fmt.Println("failed to read departments count:", err)

		return
	}

	for range nDepartments {
		var nWishes uint

		_, err = fmt.Scan(&nWishes)
		if err != nil {
			fmt.Println("failed to scan wishes count:", err)

			return
		}

		commonWish := wish.New(minSupportedTemperature, maxSupportedTemperature)

		for range nWishes {
			err = ProcessWishFromStdin(&commonWish)
			if err != nil {
				fmt.Println("failed to process wish:", err)

				return
			}

			temperature, err := commonWish.GetOptimum()
			if err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(temperature)
			}
		}
	}
}

func ProcessWishFromStdin(wish *wish.Wish) error {
	var (
		sign        string
		temperature int
	)

	_, err := fmt.Scan(&sign, &temperature)

	switch {
	case err != nil:
		return fmt.Errorf("cannot read wish: %w", err)
	case sign == ">=":
		wish.IncludeMin(temperature)

		return nil
	case sign == "<=":
		wish.IncludeMax(temperature)

		return nil
	default:
		return ErrUnknownWishSign
	}
}
