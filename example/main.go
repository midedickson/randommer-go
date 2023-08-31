package main

import (
	"fmt"

	"github.com/Double-DOS/randommer-go"
)

func main() {
	randommer.Init("4311161c65594d53a565ccc6c250049a")
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

}
