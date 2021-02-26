package brokerage

import (
	"fmt"
)

// stock broker interface
type StockBroker interface {
	commission() float64
	availableExchange() []string
}

// struct of various stock brokers
type StandardChartered struct {
	ForexRate float64
	BrokerageRate float64
	TransactionAmount float64
	AmountToUsd float64
}

type Ibkr struct {
	TransactionAmount float64
	ClearingFee float64
	FixMonthlyPlatformFee float64
	ForexRate float64
	FixBrokerageFee float64
	AmountToUsd float64
}

// commission method on standard chartered struct
func (sc StandardChartered) commission() float64 {
	return (sc.ForexRate * sc.AmountToUsd + 10) + (sc.TransactionAmount * sc.BrokerageRate)
}

// commission method on ibkr struct
func (ibkr Ibkr) commission() float64 {
	return ibkr.ClearingFee + ibkr.FixMonthlyPlatformFee + ibkr.ForexRate * ibkr.AmountToUsd + ibkr.FixBrokerageFee
}

// availableExchange method on standard chartered struct
func (sc StandardChartered) availableExchange() []string {
	return []string{"singapore", "usa", "hong kong"}
}

// availableExchange method on ibkr struct
func (ibkr Ibkr) availableExchange() []string {
	return []string{ "london", "uae"}
}

func BrokerInformation(broker StockBroker) {
	fmt.Println("----- commission -----")
	fmt.Println(broker.commission())
	fmt.Println("----- available exchange -----")
	fmt.Println(broker.availableExchange())
}