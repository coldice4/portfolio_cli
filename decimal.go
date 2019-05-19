package main

import (
	"math"
	"strconv"
	"strings"
)

type Decimal struct {
	Base int
	Factor int
}

func StringToDecimal(str string) (p Decimal, err error) {
	if str == "" {
		p.Base = 0
		p.Factor = 1
		return p, nil
	}

	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
	}

	p.Factor = calculateFactor(str)

	baseStr := strings.Replace(str, ".", "", 1)
	p.Base, err = strconv.Atoi(baseStr)

	return p, err
}

func (d *Decimal) String() (str string) {
	count := len(strconv.Itoa(d.Factor))
	baseStr := strconv.Itoa(d.Base)

	if d.Base == 0 {
		return "0"
	}

	if d.Factor > d.Base {
		zeroStr := strings.Repeat("0", count - 1)

		if baseStr[0] == '-' {
			baseStr = strings.TrimPrefix(baseStr, "-")
			str = "-" + zeroStr + baseStr
		} else {
			str = zeroStr + baseStr
		}

		commaPos := len(str) - count + 1
		str = str[:commaPos] + "." + str[commaPos:]
	} else {
		commaPos := len(baseStr) - count + 1

		str = baseStr[:commaPos] + "." + baseStr[commaPos:]
	}

	//clean str of trailing commas
	str = strings.TrimSuffix(str, ".")

	str = strings.TrimLeft(str, "0")
	if str[0] == '.' {
		str = "0" + str
	}

	return str
}

func calculateFactor(s string) (factor int) {
	decimalPointIndex := strings.Index(s, ".")
	if decimalPointIndex == -1 {
		return 1
	}
	decimalPlaces := len(s) - decimalPointIndex - 1
	factor = int(math.Pow10(decimalPlaces))

	return factor
}