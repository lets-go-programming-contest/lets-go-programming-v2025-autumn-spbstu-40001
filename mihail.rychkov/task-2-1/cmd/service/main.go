package main

import (
	"errors"
	"fmt"
	"io"
	"os"

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
		fmt.Println("Failed to read departments count", err)

		return
	}

	for range nDepartments {
		_, err := ProcessDepartmentWishes(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ProcessDepartmentWishes(istream io.Reader, logstream io.Writer) (wish.Wish, error) {
	var (
		nWishes    uint
		commonWish wish.Wish
	)

	commonWish.Init(minSupportedTemperature, maxSupportedTemperature)

	_, err := fmt.Fscan(istream, &nWishes)
	if err != nil {
		return commonWish, fmt.Errorf("failed to scan wishes count: %w", err)
	}

	for range nWishes {
		err = ProcessWish(istream, &commonWish)
		if err != nil {
			return commonWish, err
		}

		if logstream != nil {
			temperature, err := commonWish.GetOptimum()
			if err != nil {
				_, err = fmt.Fprintln(logstream, -1)
			} else {
				_, err = fmt.Fprintln(logstream, temperature)
			}

			if err != nil {
				return commonWish, fmt.Errorf("failed to print log: %w", err)
			}
		}
	}

	return commonWish, nil
}

func ProcessWish(istream io.Reader, wish *wish.Wish) error {
	var (
		sign        string
		temperature int
	)

	_, err := fmt.Fscan(istream, &sign, &temperature)

	switch {
	case err != nil:
		return fmt.Errorf("failed to process wish: %w", err)
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
