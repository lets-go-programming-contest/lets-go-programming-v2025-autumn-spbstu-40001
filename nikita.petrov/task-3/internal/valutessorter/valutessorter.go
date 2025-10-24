package valutessorter

import (
	"github.com/Nekich06/task-3/internal/currencyrate"
)

type ByValue currencyrate.CurrencyRate

func (myCurrRate ByValue) Len() int {
	return len(myCurrRate.Valutes)
}

func (myCurrRate ByValue) Swap(i, j int) {
	myCurrRate.Valutes[i], myCurrRate.Valutes[j] = myCurrRate.Valutes[j], myCurrRate.Valutes[i]
}

func (myCurrRate ByValue) Less(i, j int) bool {
	return myCurrRate.Valutes[i].Value > myCurrRate.Valutes[j].Value
}
