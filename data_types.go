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

type ValidatorResult struct {
	IsValid bool
	Message string
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

func newValidatorResult(validatorResponse apiMapResponse) *ValidatorResult {
	message, ok := validatorResponse["errorMessage"].(string)
	if !ok {
		message = "VAT is valid!"
	}
	return &ValidatorResult{
		IsValid: validatorResponse["isValid"].(bool),
		Message: message,
	}
}

type AccountNumber string

// names
type Name string

// misc
// Culture represents struct for timezone code and full area name
type Culture struct {
	Code string
	Name string
}

func makeCultures(cultureArr []apiMapResponse) []Culture {
	var cultures []Culture
	for _, culture := range cultureArr {
		cultures = append(cultures, Culture{
			Code: culture["code"].(string),
			Name: culture["name"].(string),
		})
	}
	return cultures
}

type RandomAddress string
