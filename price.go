package main

import (
	"math"
	"strconv"
	"strings"
)

type price struct {
	Base int
	Factor int
}

func StringToPrice(str string) (p price, err error) {
	if strings.Contains(str, ".") {
		str = strings.Trim(str, "0")
	}

	p.Factor = calculateFactor(str)

	baseStr := strings.Replace(str, ".", "", 1)
	p.Base, err = strconv.Atoi(baseStr)

	return p, err
}

// deprecated
/*func convertStringPriceToIntegerPrice(stringPrice string, factor int) (intPrice int) {
	floatPrice, err := strconv.ParseFloat(stringPrice, 64)
	if err != nil {
		log.Errorf("AlphaVantage: %s", err)
	}
	return int(floatPrice * float64(factor))
}*/

func calculateFactor(s string) (factor int) {
	decimalPointIndex := strings.Index(s, ".")
	if decimalPointIndex == -1 {
		return 1
	}
	decimalPlaces := len(s) - decimalPointIndex - 1
	factor = int(math.Pow10(decimalPlaces))

	return factor
}