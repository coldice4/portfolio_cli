package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type PortfolioLine struct {
	Date     time.Time
	ISIN     string
	Price    Decimal
	Quantity Decimal
	Dividend Decimal
	Taxes    Decimal
	Fees     Decimal
}

type Portfolio struct {
	Transactions []PortfolioLine
}

func (l *PortfolioLine) CSV() (output []string) {
	output = append(output, l.Date.Format("20060102"))
	output = append(output, l.ISIN)
	output = append(output, l.Price.String())
	output = append(output, l.Quantity.String())
	output = append(output, l.Dividend.String())
	output = append(output, l.Taxes.String())
	output = append(output, l.Fees.String())
	return output
}

func getPortfolioLine(line []string) (portfolioLine PortfolioLine, err error) {
	if portfolioLine.Date, err = time.Parse("20060102", line[0]); err != nil {
		return portfolioLine, err
	}
	portfolioLine.ISIN = line[1]
	portfolioLine.Price, err = StringToDecimal(line[2])
	portfolioLine.Quantity, err = StringToDecimal(line[3])
	portfolioLine.Dividend, err = StringToDecimal(line[4])
	portfolioLine.Taxes, err = StringToDecimal(line[5])
	portfolioLine.Fees, err = StringToDecimal(line[6])
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
	fmt.Printf("%s\n", "Saved transactions to file")
	return nil
}

func (p *Portfolio) ReadFromFS() (err error) {
	portfolioFilename := "portfolio.csv"
	portfolioFile, err := os.OpenFile(portfolioFilename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer portfolioFile.Close()

	p.Transactions = nil

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
	fmt.Printf("%s\n", "Loaded transaction from file")
	return nil
}

func getUserInput(stdin io.Reader) (input string) {
	reader := bufio.NewReader(stdin)
	input, _ = reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func inputPortfolioLine() (line PortfolioLine, err error) {
	fmt.Print("ISIN: ")
	line.ISIN = getUserInput(os.Stdin)

	fmt.Print("Date (YYYY-MM-DD): ")
	line.Date, err = time.Parse("2006-01-02", getUserInput(os.Stdin))
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("Price: ")
	if line.Price, err = StringToDecimal(getUserInput(os.Stdin)); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Quantity: ")
	if line.Quantity, err = StringToDecimal(getUserInput(os.Stdin)); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Dividend: ")
	if line.Dividend, err = StringToDecimal(getUserInput(os.Stdin)); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Taxes: ")
	if line.Taxes, err = StringToDecimal(getUserInput(os.Stdin)); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Fees: ")
	if line.Fees, err = StringToDecimal(getUserInput(os.Stdin)); err != nil {
		fmt.Print(err)
	}

	return line, err
}

func (p *Portfolio) PrintTransactions() {
	for _, line := range p.Transactions {
		//fmt.Printf("%+v\n", line)
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			line.Date.Format("2006-01-02"),
			line.ISIN,
			line.Price.String(),
			line.Quantity.String(),
			line.Dividend.String(),
			line.Taxes.String(),
			line.Fees.String())
	}
}
