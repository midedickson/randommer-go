package randommer

type apiMapResponse map[string]interface{}

type CardType string

type Card struct {
	Type       string
	IsoDate    string
	FullName   string
	CardNumber string
	Cvv        string
	Pin        float64
}

func newCard(cardJson apiMapResponse) *Card {
	return &Card{
		Type:       cardJson["type"].(string),
		IsoDate:    cardJson["date"].(string),
		FullName:   cardJson["fullName"].(string),
		CardNumber: cardJson["cardNumber"].(string),
		Cvv:        cardJson["cvv"].(string),
		Pin:        cardJson["pin"].(float64),
	}
}

type Name string
