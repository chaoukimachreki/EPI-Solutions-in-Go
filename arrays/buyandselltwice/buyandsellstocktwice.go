package main

import (
    "fmt"
    "math"
)

type TradeSummary struct {
    FirstBuyPrice     int
    FirstSellPrice    int
    SecondBuyPrice    int
    SecondSellPrice   int
    TotalProfit       int
}

func MaxProfitWith2Transactions(prices []int) TradeSummary {
    minPrice := math.MaxInt32
    maxProfitAfterFirstSell := 0
    maxProfitLeftAfterSecondBuy := math.MinInt32
    maxProfitAfterSecondSell := 0

    summary := TradeSummary{}

    for _, p := range prices {
        // First Buy
        if p < minPrice {
            minPrice = p
            summary.FirstBuyPrice = p
        }

        // First Sell
        if p - minPrice > maxProfitAfterFirstSell {
            maxProfitAfterFirstSell = p - minPrice
            summary.FirstSellPrice = p
        }

        // Second Buy
        if maxProfitAfterFirstSell - p > maxProfitLeftAfterSecondBuy {
            maxProfitLeftAfterSecondBuy = maxProfitAfterFirstSell - p
            summary.SecondBuyPrice = p
        }

        // Second Sell
        if p + maxProfitLeftAfterSecondBuy > maxProfitAfterSecondSell {
            maxProfitAfterSecondSell = p + maxProfitLeftAfterSecondBuy
            summary.SecondSellPrice = p
        }
    }

    summary.TotalProfit = maxProfitAfterSecondSell
    return summary
}

func main() {
    prices := []int{3, 5, 1, 100, 4, 9, 2, 10}
    result := MaxProfitWith2Transactions(prices)

    fmt.Printf("Trade Summary:\n")
    fmt.Printf("First Buy  at: %d\n", result.FirstBuyPrice)
    fmt.Printf("First Sell at: %d\n", result.FirstSellPrice)
    fmt.Printf("Second Buy at: %d\n", result.SecondBuyPrice)
    fmt.Printf("Second Sell at: %d\n", result.SecondSellPrice)
    fmt.Printf("Total Profit: %d\n", result.TotalProfit)
}

