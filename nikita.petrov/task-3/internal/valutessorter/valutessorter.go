package valutessorter

import (
	"github.com/Nekich06/task-3/internal/currencyrate"
)

type ByValue currencyrate.CurrencyRate

func (myCurrRate ByValue) Len() int {
	return len(myCurrRate.Valute)
}

func (myCurrRate ByValue) Swap(i, j int) {
	myCurrRate.Valute[i], myCurrRate.Valute[j] = myCurrRate.Valute[j], myCurrRate.Valute[i]
}

func (myCurrRate ByValue) Less(i, j int) bool {
	return myCurrRate.Valute[i].Value > myCurrRate.Valute[j].Value
}
