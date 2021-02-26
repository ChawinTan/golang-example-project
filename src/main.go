package main

import (
	"fmt"
	"app/brokerage"
)

func reverseString(s []byte) (res []byte) {
	for i := len(s) - 1; i > -1; i-- {
		res = append(res, s[i])
	}
	return res
}

func main() {
	s := []byte{}
	s = append(s, 'h')
	s = append(s, 'e')
	s = append(s, 'l')
	s = append(s, 'l')
	s = append(s, 'o')
	fmt.Println(string(reverseString(s)))

	scb := brokerage.StandardChartered{
		ForexRate: 0.02,
		BrokerageRate: 0.03,
		TransactionAmount: 10000,
		AmountToUsd: 10000}

	ibkr := brokerage.Ibkr{
		TransactionAmount: 10000,
		ClearingFee: 2,
		FixMonthlyPlatformFee: 10,
		ForexRate: 0.0002,
		FixBrokerageFee: 2,
		AmountToUsd: 10000}

	fmt.Println("--- standard chartered ---")
	brokerage.BrokerInformation(scb)
	fmt.Println("--- interactive broker ---")
	brokerage.BrokerInformation(ibkr)
}