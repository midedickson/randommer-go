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

// finance
type CryptoAddressType string
type CryptoAddress struct {
	Address string
	Type    string
}
type Country struct {
	Name        string
	CountryCode string
}

func newCryptoAddress(cryptoAddressJson apiMapResponse) *CryptoAddress {
	return &CryptoAddress{
		Address: cryptoAddressJson["address"].(string),
		Type:    cryptoAddressJson["type"].(string),
	}
}

func makeCountries(countryArr []apiMapResponse) []Country {
	var countries []Country
	for _, country := range countryArr {
		countries = append(countries, Country{
			Name:        country["name"].(string),
			CountryCode: country["countryCode"].(string),
		})
	}
	return countries

}

type AccountNumber string

// names
type Name string
