package pkg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ExchangeResponse struct {
	Rates map[string]float32
}

type ExchangeRate struct {
	Currency string
	Rate     float32
}

const USD = "USD"
const GBP = "GBP"
const RUB = "RUB"

func GetExchanges() ([]ExchangeRate, error) {
	resp, err := http.Get("https://api.exchangeratesapi.io/latest")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var ex ExchangeResponse

	err = json.Unmarshal(body, &ex)
	if err != nil {
		log.Fatal(err)
	}

	rates := []ExchangeRate{}
	rates = append(rates, ExchangeRate{Currency: USD, Rate: ex.Rates[USD]})
	rates = append(rates, ExchangeRate{Currency: GBP, Rate: ex.Rates[GBP]})
	rates = append(rates, ExchangeRate{Currency: RUB, Rate: ex.Rates[RUB]})

	return rates, nil
}

// func main() {
// 	resp, err := http.Get("https://api.exchangeratesapi.io/latest")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var ex ExchangeResponse

// 	err = json.Unmarshal(body, &ex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%s %f\n", "USD", ex.Rates["USD"])
// 	fmt.Printf("%s %f\n", "GBP", ex.Rates["GBP"])
// 	fmt.Printf("%s %f\n", "RUB", ex.Rates["RUB"])
// }
