package main

import "fmt"

type (
	Plsql struct {
		name string
	}

	Mongo struct {
		name string
	}

	DbMethods interface {
		Manipute() string
	}
)

func (pl Plsql) Manipute() string {
	return "some plsql db related stuff"
}

func (md *Mongo) manipulate() string {
	return "some mongo related stuff"
}

func createOrUpdateData(db DbMethods) {
	db.Manipute()
}

func main() {
	// Instatiate with plsql
	// Instantiate with mongo
	a := 2
	ps := Plsql{
		name: "s",
	}
	fmt.Println(a, ps)
}

func maxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			// If diff of current stock with minPrice is greater
			// update the profit
			profit = prices[i] - minPrice
		}
	}

	return profit
}
