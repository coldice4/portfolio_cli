package main

import (
	"os"
	"strings"
	"testing"
)

var testValues = make(map[string]Decimal)

func setup() {
	testValues["123.45"] = Decimal{Base: 12345, Factor:100}
	testValues["12345.45"] = Decimal{Base: 1234545, Factor:100}
	testValues["123.00"] = Decimal{Base: 123, Factor:1}
	testValues["0.20"] = Decimal{Base: 2, Factor:10}
	testValues["0.2"] = Decimal{Base: 2, Factor:10}
	testValues["0.0003"] = Decimal{Base: 3, Factor:10000}
	testValues["-0.03"] = Decimal{Base: -3, Factor:100}
	testValues["123"] = Decimal{Base: 123, Factor:1}
	testValues["-123"] = Decimal{Base: -123, Factor:1}
	testValues["0.64"] = Decimal{Base: 64, Factor: 100}
	testValues["0.123"] = Decimal{Base: 123, Factor:1000}
	testValues["0"] = Decimal{Base:0, Factor:1}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestStringToPrice(t *testing.T) {
	for str, price := range testValues {
		p, err := StringToDecimal(str)
		if p.Base != price.Base || p.Factor != price.Factor {
			t.Errorf("%s\tBase: %d != %d\tFactor: %d != %d\n", str, p.Base, price.Base, p.Factor, price.Factor)
		}
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDecimal_String(t *testing.T) {
	for expectedStr, price := range testValues {
		expectedStr = strings.TrimRight(expectedStr, "0")
		expectedStr = strings.TrimSuffix(expectedStr, ".")

		// edge case for 0
		if expectedStr == "" {
			expectedStr = "0"
		}

		str := price.String()
		if str != expectedStr {
			t.Errorf("%s != %s", str, expectedStr)
		}
	}
}