package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type PortfolioLine struct {
	Date     time.Time
	ISIN     string
	Price    decimal.Decimal
	Quantity decimal.Decimal
	Dividend decimal.Decimal
	Taxes    decimal.Decimal
	Fees     decimal.Decimal
}

type Security struct {
	Price decimal.Decimal
	Quantity decimal.Decimal
	Dividend decimal.Decimal
	Taxes decimal.Decimal
	Fees decimal.Decimal
	Symbol string
}

type Portfolio struct {
	Transactions []PortfolioLine
	ISIN map[string]Security
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

func getPortfolioLine(line []string) (err error) {
	var portfolioLine PortfolioLine
	if portfolioLine.Date, err = time.Parse("20060102", line[0]); err != nil {
		return err
	}
	portfolioLine.ISIN = line[1]
	portfolioLine.Price, err = decimal.NewFromString(line[2])
	portfolioLine.Quantity, err = decimal.NewFromString(line[3])
	portfolioLine.Dividend, err = decimal.NewFromString(line[4])
	portfolioLine.Taxes, err = decimal.NewFromString(line[5])
	portfolioLine.Fees, err = decimal.NewFromString(line[6])

	portfolio.Transactions = append(portfolio.Transactions, portfolioLine)
	return err
}

func getMappingLine(line []string) (err error) {
	var sec Security
	sec.Symbol = line[1]
	portfolio.ISIN[line[0]] = sec
	fmt.Printf("%+v\n", portfolio.ISIN[line[0]])
	return err
}

func (p *Portfolio) WriteToFS() (err error) {
	portfolioFilename := "portfolio.csv"
	portfolioFile, err := os.OpenFile(portfolioFilename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer portfolioFile.Close()

	portfolioFile.WriteString("transactions\n")
	w := csv.NewWriter(portfolioFile)
	for _, line := range p.Transactions {
		if err := w.Write(line.CSV()); err != nil {
			return err
		}
	}
	w.Flush()

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
	p.ISIN = make(map[string]Security)

	lineReader := bufio.NewReader(portfolioFile)

	var lineHandler func([]string)(error)

	for {
		line, err := lineReader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		//remove the new line because it causes confusion later on
		line = strings.TrimSuffix(line, "\n")

		fields := strings.Split(line, ",")
		
		switch fields[0] {
		case "transactions":
			lineHandler = getPortfolioLine
		case "mapping":
			lineHandler = getMappingLine
		default:
			err = lineHandler(fields)
		}
	}

	fmt.Printf("%s\n", "Loaded transaction from file")
	return nil
}

func getUserInput(stdin io.Reader) (input string) {
	reader := bufio.NewReader(stdin)
	input, _ = reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func getUserInputDecimal(stdin io.Reader) (input decimal.Decimal, err error) {
	reader := bufio.NewReader(stdin)
	inputStr, err := reader.ReadString('\n')
	inputStr = strings.TrimSuffix(inputStr, "\n")
	if inputStr == "" {
		inputStr = "0"
	}
	input, err = decimal.NewFromString(inputStr)

	return input, err
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
	if line.Price, err = getUserInputDecimal(os.Stdin); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Quantity: ")
	if line.Quantity, err = getUserInputDecimal(os.Stdin); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Dividend: ")
	if line.Dividend, err = getUserInputDecimal(os.Stdin); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Taxes: ")
	if line.Taxes, err = getUserInputDecimal(os.Stdin); err != nil {
		fmt.Print(err)
	}

	fmt.Print("Fees: ")
	if line.Fees, err = getUserInputDecimal(os.Stdin); err != nil {
		fmt.Print(err)
	}

	return line, err
}

func (p *Portfolio) PrintTransactions() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "Date\tISIN\tPrice\tQuantity\tDividend\tTaxes\tFees\t\n")
	for _, line := range p.Transactions {
		//fmt.Printf("%+v\n", line)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
			line.Date.Format("2006-01-02"),
			line.ISIN,
			line.Price.String(),
			line.Quantity.String(),
			line.Dividend.String(),
			line.Taxes.String(),
			line.Fees.String())
	}
	w.Flush()
}

func (p *Portfolio) PrintStatus() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "ISIN\tPrice\tQuantity\tDividend\tTaxes\tFees\tSymbol\t\n")
	for _, line := range p.Transactions {
		//var sec Security
		sec := p.ISIN[line.ISIN]
		sec.Quantity = p.ISIN[line.ISIN].Quantity.Add(line.Quantity)
		sec.Fees = p.ISIN[line.ISIN].Fees.Add(line.Fees)
		sec.Taxes = p.ISIN[line.ISIN].Taxes.Add(line.Taxes)
		sec.Dividend = p.ISIN[line.ISIN].Taxes.Add(line.Dividend)

		p.ISIN[line.ISIN] = sec
	}

	for isin, status := range p.ISIN {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n",
			isin,
			status.Price,
			status.Quantity,
			status.Dividend,
			status.Taxes,
			status.Fees,
			status.Symbol)
	}

	w.Flush()
}
