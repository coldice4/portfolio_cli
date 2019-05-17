package main

import "testing"

func TestStringToPrice(t *testing.T) {
	var testValues = make(map[string]price)
	testValues["123.45"] = price{Base:12345, Factor:100}
	testValues["123.00"] = price{Base:123, Factor:1}
	testValues["0.20"] = price{Base:2, Factor:10}
	testValues["0.2"] = price{Base:2, Factor:10}
	testValues["123"] = price{Base:123, Factor:1}
	testValues["-123"] = price{Base:-123, Factor:1}

	for str, price := range testValues {
		p, err := StringToPrice(str)
		if p.Base != price.Base || p.Factor != price.Factor {
			t.Errorf("%s\tBase: %d != %d\tFactor: %d != %d\n", str, p.Base, price.Base, p.Factor, price.Factor)
		}
		if err != nil {
			t.Error(err)
		}
	}
}
