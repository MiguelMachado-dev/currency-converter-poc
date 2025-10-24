package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type CurrencyRates struct {
	Base  string                     `json:"base"`
	Date  string                     `json:"date"`
	Rates map[string]decimal.Decimal `json:"rates"`
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Erro: Faltando argumentos. \nUso: ./currency-converter <valor> <moeda>")
	}
	val := strings.ReplaceAll(os.Args[1], ".", "")
	val = strings.ReplaceAll(val, ",", ".")
	curr := strings.ToUpper(os.Args[2])

	valorEmReais, err := decimal.NewFromString(val)
	check(err)

	f, err := os.Open("rates.json")
	check(err)
	defer f.Close()

	byteValue, err := io.ReadAll(f)
	check(err)

	var currencyRates CurrencyRates
	err = json.Unmarshal(byteValue, &currencyRates)
	check(err)

	currRate, ok := currencyRates.Rates[curr]
	if !ok {
		log.Fatalf("Moeda %s não encontrada", curr)
	}

	exchange := valorEmReais.Mul(currRate)
	fmt.Println(exchange.StringFixed(2))
}
