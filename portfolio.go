package main

import (
	"encoding/csv"
	"fmt"
	"io"
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
	output = append(output, l.Date.Format("20060102"))
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

func getPortfolioLine(line []string) (portfolioLine PortfolioLine, err error) {
	if portfolioLine.Date, err = time.Parse("20060102", line[0]); err != nil {
		return portfolioLine, err
	}
	portfolioLine.ISIN = line[1]
	portfolioLine.Price.Base, err = strconv.Atoi(line[2])
	portfolioLine.Price.Factor, err = strconv.Atoi(line[3])
	portfolioLine.Quantity.Base, err = strconv.Atoi(line[4])
	portfolioLine.Quantity.Factor, err = strconv.Atoi(line[5])
	portfolioLine.Dividend.Base, err = strconv.Atoi(line[6])
	portfolioLine.Dividend.Factor, err = strconv.Atoi(line[7])
	portfolioLine.Taxes.Base, err = strconv.Atoi(line[8])
	portfolioLine.Taxes.Factor, err = strconv.Atoi(line[9])
	portfolioLine.Taxes.Base, err = strconv.Atoi(line[10])
	portfolioLine.Taxes.Factor, err = strconv.Atoi(line[11])
	return portfolioLine, err
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

func (p *Portfolio) ReadFromFS() (err error) {
	portfolioFilename := "portfolio.csv"
	portfolioFile, err := os.OpenFile(portfolioFilename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer portfolioFile.Close()

	r := csv.NewReader(portfolioFile)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		portfolioLine, err := getPortfolioLine(line)
		if err != nil {
			return err
		}
		p.Transactions = append(p.Transactions, portfolioLine)
	}
	return nil
}
