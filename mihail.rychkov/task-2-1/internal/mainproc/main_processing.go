package mainproc

import (
	"errors"
	"fmt"
	"io"
)

var ErrUnknownWishSign = errors.New("unknown comparison sign")

func ProcessWish(istream io.Reader, wish *Wish) error {
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

func ProcessDepartmentWishes(istream io.Reader, logstream io.Writer) (Wish, error) {
	var (
		nWishes    uint
		commonWish = Wish{15, 30}
	)

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
