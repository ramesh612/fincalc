package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"sort"
)

type MetaData struct {
	symbol        string
	lastRefreshed time.Time
}

type DailyData struct {
	openPrice          float32
	highPrice          float32
	lowPrice           float32
	closePrice         float32
	adjustedClosePrice float32
	volume             int
	dividendAmount     float32
	splitCoefficient   float32
}

type StockData struct {
	metaData     *MetaData
	dailyDataMap map[time.Time]*DailyData
}

const (
	SYMBOL         = "2. Symbol"
	LAST_REFRESHED = "3. Last Refreshed"
	TIME_ZONE      = "5. Time Zone"
)

const (
	OPEN              = "1. open"
	HIGH              = "2. high"
	LOW               = "3. low"
	CLOSE             = "4. close"
	ADJUSTED_CLOSE    = "5. adjusted close"
	VOLUME            = "6. volume"
	DIVIDEND_AMOUNT   = "7. dividend amount"
	SPLIT_COEFFICIENT = "8. split coefficient"
)

const (
	DATE_FORMAT = "2006-01-02"
)

func getData(symbol string) string {
	// url := "https://alpha-vantage.p.rapidapi.com/query?interval=5min&function=TIME_SERIES_INTRADAY&symbol=MSFT&datatype=json&output_size=compact"
	// url := "https://alpha-vantage.p.rapidapi.com/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=MSFT&datatype=json"
	url := "https://alpha-vantage.p.rapidapi.com/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=" + symbol + "&datatype=json&output_size=compact"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "cc4c30cf92msh8a62763f9a03861p1d896cjsn6596481e7261")
	req.Header.Add("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func getDate(dateStr string, timeZone string) time.Time {
	location, _ := time.LoadLocation(timeZone)
	dateObj, _ := time.ParseInLocation(DATE_FORMAT, dateStr, location)
	return dateObj
}

func parseData(jsonStr string) *StockData {
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)

	// get lastRefreshedDate and timeZone
	metadata_map, _ := data["Meta Data"].(map[string]interface{})

	symbol := strings.ToUpper(metadata_map[SYMBOL].(string))
	lastRefreshed := metadata_map[LAST_REFRESHED].(string)
	timeZone := metadata_map[TIME_ZONE].(string)
	lastRefreshedDate := getDate(lastRefreshed, timeZone)

	var stockData StockData
	stockData.metaData = new(MetaData)
	stockData.metaData.symbol = symbol
	stockData.metaData.lastRefreshed = lastRefreshedDate

	// get stock data by day
	timeseries_daily_map, _ := data["Time Series (Daily)"].(map[string]interface{})
	stockData.dailyDataMap = make(map[time.Time]*DailyData)
	for k, v := range timeseries_daily_map {
		stock_data_map, _ := v.(map[string]interface{})
		dataDate := getDate(k, timeZone)
		var dailyData DailyData
		dailyData.openPrice = String2Float(stock_data_map[OPEN].(string))
		dailyData.lowPrice = String2Float(stock_data_map[LOW].(string))
		dailyData.highPrice = String2Float(stock_data_map[HIGH].(string))
		dailyData.closePrice = String2Float(stock_data_map[CLOSE].(string))
		dailyData.adjustedClosePrice = String2Float(stock_data_map[ADJUSTED_CLOSE].(string))
		dailyData.volume = String2Int(stock_data_map[VOLUME].(string))
		dailyData.dividendAmount = String2Float(stock_data_map[DIVIDEND_AMOUNT].(string))
		dailyData.splitCoefficient = String2Float(stock_data_map[SPLIT_COEFFICIENT].(string))
		stockData.dailyDataMap[dataDate] = &dailyData
	}
	return &stockData
}

func main() {
	if len(os.Args) > 1 {
		body := getData(os.Args[1])
		stockData := parseData(body)
		var keys []time.Time
		for k := range stockData.dailyDataMap {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].Before(keys[j])
		})
		fmt.Printf("Ticker: %s, Last Refreshed: %v\n",
			stockData.metaData.symbol, stockData.metaData.lastRefreshed)
		fmt.Printf("%10s %8s %8s %8s %8s %8s %9s %8s %8s\n",
			"DATE", "OPEN", "LOW", "HIGH", "CLOSE", "ADJUSTED", "VOLUME", "DIVIDEND", "SPLIT")
		for _, k := range keys {
			fmt.Printf("%10s ", k.Format(DATE_FORMAT))
			v := stockData.dailyDataMap[k]
			fmt.Printf("%8.2f %8.2f %8.2f %8.2f %8.2f %9d %8.2f %8.2f\n",
				v.openPrice,
				v.lowPrice,
				v.highPrice,
				v.closePrice,
				v.adjustedClosePrice,
				v.volume,
				v.dividendAmount,
				v.splitCoefficient)
		}
	} else {
		fmt.Println("this program requires a stock symbol as argument")
	}
}

func String2Float(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func String2Int(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
