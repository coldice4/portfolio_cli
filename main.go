package main

import (
	"fmt"
	"time"
)

var tickers = []string{
	"C050.DE",
	"PND.F",
	"BAYN.DE",
	"59A.F",
}

var history = make(map[string][]HistoryLine)

func main() {
	fmt.Println("PORTFOLIO_CLI v0.1")

	/*for _, ticker := range tickers {
		ScheduleAVGetWeekly(ticker)
	}

	fmt.Printf("%+v\n", history)*/



	var portfolio Portfolio
	line := PortfolioLine{Date: time.Now(), ISIN: "ABC123XYZ", Price: price{Base: 2413, Factor: 100}, Quantity: price{Base:200, Factor: 1}}
	portfolio.Transactions = append(portfolio.Transactions, line)
	if err := portfolio.WriteToFS(); err != nil {
		panic(err)
	}
}
