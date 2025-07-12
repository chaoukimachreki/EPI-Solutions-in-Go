package main

import (
    "fmt"
    "errors"
   // "math"
)

type tradeSummary struct {
    boughtAt float64
    soldAt   float64
    profit   float64
}

func (t tradeSummary) String() string {
    return fmt.Sprintf(
        "📈 Trade Summary\n------------------\nBought at: %.2f\nSold at: %.2f\nProfit: %.2f\n",
        t.boughtAt, t.soldAt, t.profit)
}

func buyAndSellStockOnce(prices []float64) (tradeSummary, error) {
    if len(prices) == 0 {
        return tradeSummary{}, errors.New("please add an array of prices")
    }

    minPriceSoFar := prices[0]
    maxProfit := float64(0)
    buyPrice := float64(0)
    sellPrice := float64(0)

    for _, price := range prices {
        potentialProfit := price - minPriceSoFar
        if potentialProfit > maxProfit {
            maxProfit = potentialProfit
            buyPrice = minPriceSoFar
            sellPrice = price
        }
        if price < minPriceSoFar {
            minPriceSoFar = price
        }
    }

    return tradeSummary{
        boughtAt: buyPrice,
        soldAt:   sellPrice,
        profit:   maxProfit,
    }, nil
}

func main() {
    stockPrices := []float64{20.0, 18.0, 25.0, 17.0}
    tradeSummary, err := buyAndSellStockOnce(stockPrices)
    if err != nil {
        fmt.Printf("❗ Error: %v\n", err)
        return
    }
    fmt.Println(tradeSummary)
}

