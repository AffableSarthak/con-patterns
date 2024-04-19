package main

import "fmt"

func main() {
	const s = "สวัสดี"
	for idx, runeValue := range s {
		fmt.Printf("%q starts at %d\n", runeValue, idx)
	}
}
