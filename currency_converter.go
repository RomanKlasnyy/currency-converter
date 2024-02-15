package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	cache := make(map[string]float64)

	var x string
	fmt.Print("Enter your currency code: ")
	fmt.Scan(&x)
	x = strings.ToUpper(x)

	r, err := http.Get("http://www.floatrates.com/daily/" + x + ".json")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if x != "USD" {
		cache["USD"] = getValueFromJSON(body, "usd")
	}
	if x != "EUR" {
		cache["EUR"] = getValueFromJSON(body, "eur")
	}

	for {
		var y string
		fmt.Print("Enter exchange currency: ")
		fmt.Scan(&y)
		y = strings.ToUpper(y)
		if y == "" {
			break
		}

		var numStr string
		fmt.Printf("Enter amount of %s: ", x)
		fmt.Scan(&numStr)

		var num float64
		if strings.Contains(numStr, ".") {
			num, err = strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Println("Invalid input:", err)
				continue
			}
		} else {
			numInt, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Invalid input:", err)
				continue
			}
			num = float64(numInt)
		}

		fmt.Println("Checking the cache...")
		if rate, ok := cache[y]; ok {
			fmt.Println("Oh! It is in the cache!")
			fmt.Printf("You received %.2f %s\n", rate*num, y)
		} else {
			fmt.Println("Sorry, but it is not in the cache!")
			cache[y] = getValueFromJSON(body, strings.ToLower(y))
			fmt.Printf("You received %.2f %s\n", cache[y]*num, y)
		}
	}
}

func getValueFromJSON(body []byte, currency string) float64 {
	value := 0.0
	if strings.Contains(string(body), currency) {
		startIndex := strings.Index(string(body), currency+"\":{\"rate\":") + len(currency) + 9
		endIndex := strings.Index(string(body)[startIndex:], ",")
		rateStr := string(body)[startIndex : startIndex+endIndex]
		value, _ = strconv.ParseFloat(rateStr, 64)
	}
	return value
}
