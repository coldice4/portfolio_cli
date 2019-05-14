package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type PortfolioLine struct {
	Date time.Time
	ISIN string
	Price price
	Quantity price
	Dividend price
	Taxes price
	Fees price
}

type Portfolio struct {
	Transactions []PortfolioLine
}

func (l *PortfolioLine) CSV() (output []string) {
	output = append(output, l.Date.String())
	output = append(output, l.ISIN)
	output = append(output, strconv.Itoa(l.Price.Base))
	output = append(output, strconv.Itoa(l.Price.Factor))
	output = append(output, strconv.Itoa(l.Quantity.Base))
	output = append(output, strconv.Itoa(l.Quantity.Factor))
	output = append(output, strconv.Itoa(l.Dividend.Base))
	output = append(output, strconv.Itoa(l.Dividend.Factor))
	output = append(output, strconv.Itoa(l.Taxes.Base))
	output = append(output, strconv.Itoa(l.Taxes.Factor))
	output = append(output, strconv.Itoa(l.Fees.Base))
	output = append(output, strconv.Itoa(l.Fees.Factor))
	fmt.Printf("%+v\n", output)
	return output
}

func (p *Portfolio) WriteToFS() (err error) {
	portfolioFilename := "portfolio.csv"
	portfolioFile, err := os.OpenFile(portfolioFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer portfolioFile.Close()

	w := csv.NewWriter(portfolioFile)
	defer w.Flush()
	for _, line := range p.Transactions {
		if err := w.Write(line.CSV()); err != nil {
			return err
		}
	}
	return nil
}
