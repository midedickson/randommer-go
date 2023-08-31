package randommer

import "fmt"

func GetCardTypes() []CardType {
	path := req.buildPath("Card/Types")
	res, _ := req.Get(path)
	defer res.Body.Close()
	var types []CardType
	parseRequestBody(res.Body, &types)
	return types
}

func GetRandomCard() Card {
	path := req.buildPath("Card")
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cardJson apiMapResponse
	parseRequestBody(res.Body, &cardJson)
	card := *newCard(cardJson)
	return card
}

func GetCardByType(cardType string) Card {
	path := req.buildPath(fmt.Sprintf("Card?type=%s", cardType))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cardJson apiMapResponse
	parseRequestBody(res.Body, &cardJson)
	card := *newCard(cardJson)
	return card
}

func GetRandomNames(nameType string, quantity int) []Name {
	path := req.buildPath(fmt.Sprintf("Name?nameType=%s&quantity=%d", nameType, quantity))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var names []Name
	parseRequestBody(res.Body, &names)
	return names
}
