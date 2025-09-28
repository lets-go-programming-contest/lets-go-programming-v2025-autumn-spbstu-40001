package mainproc;

import "fmt";
import "errors";
import "io";

var ErrUnknownWishSign = errors.New("unknown comparison sign");
func ProcessWish(istream io.Reader, wish *Wish) error {
	var sign string;
	var temperature int;
	_, err := fmt.Fscan(istream, &sign, &temperature);
	switch {
	case err != nil:
		return fmt.Errorf("failed to process wish: %w", err);
	case sign == ">=":
		wish.IncludeMin(temperature);
		return nil;
	case sign == "<=":
		wish.IncludeMax(temperature);
		return nil;
	default:
		return ErrUnknownWishSign;
	}
}

func ProcessDepartmentWishes(istream io.Reader, logstream io.Writer) (Wish, error) {
	var nWishes uint;
	var commonWish = Wish{15, 30};
	_, err := fmt.Fscan(istream, &nWishes);
	if (err != nil) {
		return commonWish, fmt.Errorf("failed to scan wishes count: %w", err);
	}

	for range(nWishes) {
		err = ProcessWish(istream, &commonWish);
		if (err != nil) { return commonWish, err; }

		if (logstream != nil) {
			temperature, err := commonWish.GetOptimum();
			if (err != nil) {
				fmt.Fprintln(logstream, -1);
			} else {
				fmt.Fprintln(logstream, temperature);
			}
		}
	}
	return commonWish, nil;
}
