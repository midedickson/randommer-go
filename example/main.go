package main

import (
	"fmt"

	"github.com/Double-DOS/randommer-go"
)

func main() {
	randommer.Init("fba3fcbcaea346fc8f34acbb3d3efc60")
	card_types := randommer.GetCardTypes()
	fmt.Println(card_types)
	randomCard := randommer.GetRandomCard()
	fmt.Println(randomCard)
	cardByType := randommer.GetCardByType("Visa")
	fmt.Println(cardByType)
	fullNames := randommer.GetRandomNames("fullname", 10)
	fmt.Println(fullNames)
	firstNames := randommer.GetRandomNames("firstname", 10)
	fmt.Println(firstNames)
	surNames := randommer.GetRandomNames("surname", 10)
	fmt.Println(surNames)
	cryptoAddressTypes := randommer.GetCryptoAdressTypes()
	fmt.Println(cryptoAddressTypes)
	cryptoAddressByType := randommer.GetCryptoAdressByType(string(cryptoAddressTypes[0]))
	fmt.Println(cryptoAddressByType)
	availableCountries := randommer.GetAvailabeCountries()
	fmt.Println(availableCountries)
	accountNumber := randommer.GetAccountNumberByCountry(availableCountries[0].CountryCode)
	fmt.Println(accountNumber)
}
