package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("How much do you want to invest? ")
	fmt.Scan(&investmentAmount)

	fmt.Print("How many years do you want to hold to your investment? ")
	fmt.Scan(&years)

	fmt.Print("What is the expected return rate? ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Expected value: %f\n", futureValue)
	fmt.Printf("Expected inflated value: %f\n", futureRealValue)
}

// calculates values based on investment values
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (value, realvalue float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
