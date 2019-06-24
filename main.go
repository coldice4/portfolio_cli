package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tickers = []string{
	"C050.DE",
	"PND.F",
	"BAYN.DE",
	"59A.F",
}

var history = make(map[string][]HistoryLine)

var portfolio Portfolio

func main() {
	fmt.Println("PORTFOLIO_CLI v0.1")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err := inputController(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	/*for _, ticker := range tickers {
		ScheduleAVGetWeekly(ticker)
	}

	fmt.Printf("%+v\n", history)*/
}

func inputController(input string) (err error) {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	switch args[0] {
	case "load":
		return portfolio.ReadFromFS()
	case "save":
		return portfolio.WriteToFS()
	case "add":
		line, err := inputPortfolioLine()
		if err == nil {
			fmt.Printf("%+v\n", line)
			portfolio.Transactions = append(portfolio.Transactions, line)
		}
		return err
	case "transactions": {
		portfolio.PrintTransactions()
		return err
	}
	case "calc": {
		portfolio.CalculateHistory()
	}
	case "status": {
		portfolio.PrintStatus()
		return err
	}
	case "exit":
		os.Exit(0)
	}
	return nil
}
