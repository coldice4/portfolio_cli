package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestCurrencyAmount_EuroValue(t *testing.T) {
	// all fields empty
	cA := CurrencyAmount{}
	value, err := cA.EuroValue()
	if value.IsZero() != true || err == nil {
		fmt.Printf("Empty: %+v\n", cA)
		fmt.Printf("value: %s\terr: %s\n", value, err)
		t.Fail()
	}

	// euro is filled
	dec, _ := decimal.NewFromString("13.37")
	cA = CurrencyAmount{euro: dec}
	value, err = cA.EuroValue()
	if value != dec || err != nil {
		fmt.Printf("Euro is filled: %+v\n", cA)
		fmt.Printf("value: %s\terr: %s\n", value, err)
		t.Fail()
	}

	// foreign is filled
	dec1, _ := decimal.NewFromString("13.37")
	dec2, _ := decimal.NewFromString("1.23")
	checkDec, _ := decimal.NewFromString("16.4451")
	cA = CurrencyAmount{foreignCurrencyAmount: dec1, exchangeRate: dec2}
	value, err = cA.EuroValue()
	if value.Equal(checkDec) == false || err != nil {
		fmt.Printf("Foreign is filled: %+v\n", cA)
		fmt.Printf("value: %s\tcheck: %s\terr: %s\n", value, checkDec, err)
		t.Fail()
	}

	// foreign is partially filled
	cA = CurrencyAmount{foreignCurrencyAmount: dec1}
	value, err = cA.EuroValue()
	if value.IsZero() == false || err == nil {
		fmt.Printf("Foreign is partially filled: %+v\n", cA)
		fmt.Printf("value: %s\terr: %s\n", value, err)
		t.Fail()
	}
}
