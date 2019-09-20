package main

import (
	"errors"
	"github.com/shopspring/decimal"
)

type Transaction struct {

}

type CurrencyAmount struct {
	symbol string
	exchangeRate decimal.Decimal
	euro decimal.Decimal
	foreignCurrencyAmount decimal.Decimal
}

func (cA *CurrencyAmount) Symbol() (string) {
	return cA.symbol
}

func (cA *CurrencyAmount) EuroValue() (decimal.Decimal, error) {
	var euroValue decimal.Decimal
	var err error

	if cA.euro.IsZero() && (cA.foreignCurrencyAmount.IsZero() || cA.exchangeRate.IsZero()) {
		err = errors.New("missing values, can't calculate euro value")
	}

	if cA.euro.IsZero() == false && cA.exchangeRate.IsZero() && cA.foreignCurrencyAmount.IsZero() {
		euroValue = cA.euro
	}

	if cA.euro.IsZero() && cA.exchangeRate.IsZero() == false && cA.foreignCurrencyAmount.IsZero() == false {
		euroValue = cA.exchangeRate.Mul(cA.foreignCurrencyAmount)
	}

	return euroValue, err
}
