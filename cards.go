package randommer

import "fmt"

func GetCardTypes() []CardType {
	path := req.buildPath(CARD_TYPES)
	res, _ := req.Get(path)
	defer res.Body.Close()
	var types []CardType
	parseRequestBody(res.Body, &types)
	return types
}

func GetRandomCard() Card {
	path := req.buildPath(CARD)
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cardJson apiMapResponse
	parseRequestBody(res.Body, &cardJson)
	card := *newCard(cardJson)
	return card
}

func GetCardByType(cardType string) Card {
	path := req.buildPath(fmt.Sprintf(CARD_BY_TYPE, cardType))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cardJson apiMapResponse
	parseRequestBody(res.Body, &cardJson)
	card := *newCard(cardJson)
	return card
}
