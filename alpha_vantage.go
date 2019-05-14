package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var muxAPI sync.Mutex
var muxSchedule sync.Mutex
var scheduledSymbols map[string]int

type ResponseData struct {
	MetaData        MetaData
	TimeSeriesEntry []TimeSeriesEntry
}

type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type TimeSeriesEntry struct {
	Time   time.Time
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume int    `json:"5. volume,string"`
}

type HistoryLine struct {
	Date time.Time
	Ticker string
	ISIN string
	Close price
}

type jsonMap map[string]TimeSeriesEntry

//multithreading wrapper for AVGetWeekly
func ScheduleAVGetWeekly(symbol string) {
	muxSchedule.Lock()
	_, ok := scheduledSymbols[symbol]
	if ok == true {
		muxSchedule.Unlock()
		return
	}
	scheduledSymbols = make(map[string]int)
	scheduledSymbols[symbol] = 1
	muxSchedule.Unlock()

	muxAPI.Lock()

	start := time.Now()
	history[symbol] = AVGetWeekly(symbol)
	end := time.Now()

	neededPause := 12*time.Second - end.Sub(start)
	time.Sleep(neededPause)

	muxAPI.Unlock()
}

func AVGetWeekly(symbol string) ([]HistoryLine) {
	var historyLines []HistoryLine

	url := "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=" + symbol + "&apikey=DSWJ3IXD54JK3LFL&datatype=json"
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Errorf("AlphaVantage: %s", err)
	}

	var responseData ResponseData

	jsonData, _ := ioutil.ReadAll(response.Body)
	jsonRaw := map[string]json.RawMessage{}
	err = json.Unmarshal(jsonData, &jsonRaw)
	if err != nil {
		log.Errorf("AlphaVantage: %s", err)
	}

	var jsonMap jsonMap
	err = json.Unmarshal(jsonRaw["Time Series (Daily)"], &jsonMap)
	if err != nil {
		log.Errorf("AlphaVantage: %s", err)
	}

	responseData.TimeSeriesEntry = []TimeSeriesEntry{}
	for key, _ := range jsonMap {
		loc, _ := time.LoadLocation(responseData.MetaData.TimeZone)
		timeP, _ := time.ParseInLocation("2006-01-02", key, loc)
		tmp := jsonMap[key]

		//calculate decimal places
		factor := calculateFactor(tmp.Open)

		var historyLine HistoryLine
		historyLine.Ticker = symbol
		historyLine.Date = timeP
		historyLine.Close.Factor = factor
		historyLine.Close.Base = convertStringPriceToIntegerPrice(tmp.Close, factor)
		historyLines = append(historyLines, historyLine)




	}
	return historyLines
}

func convertStringPriceToIntegerPrice(stringPrice string, factor int) (intPrice int) {
	floatPrice, err := strconv.ParseFloat(stringPrice, 64)
	if err != nil {
		log.Errorf("AlphaVantage: %s", err)
	}
	return int(floatPrice * float64(factor))
}

func calculateFactor(s string) (factor int) {
	decimalPointIndex := strings.Index(s, ".")
	if decimalPointIndex == -1 {
		log.Errorf("decimal point not found")
	}
	decimalPlaces := len(s) - decimalPointIndex - 1
	factor = int(math.Pow10(decimalPlaces))
	return factor
}
