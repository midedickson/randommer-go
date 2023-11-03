package randommer

import "fmt"

func GetCryptoAdressTypes() []CryptoAddressType {
	path := req.buildPath(CRYPTOADRESS_TYPES)
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cryptoAddressTypes []CryptoAddressType
	parseRequestBody(res.Body, &cryptoAddressTypes)
	return cryptoAddressTypes
}

func GetCryptoAdressByType(cryptoType string) CryptoAddress {
	path := req.buildPath(fmt.Sprintf(CRYPTO_ADDRESS_BY_TYPE, cryptoType))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cryptoAddressJson apiMapResponse
	parseRequestBody(res.Body, &cryptoAddressJson)
	cryptoAddress := newCryptoAddress(cryptoAddressJson)
	return *cryptoAddress
}

func GetAvailabeCountries() []Country {
	path := req.buildPath(AVAILABLE_COUNTRIES)
	res, _ := req.Get(path)
	defer res.Body.Close()
	var countryJson []apiMapResponse
	parseRequestBody(res.Body, &countryJson)
	country := makeCountries(countryJson)
	return country
}

func GetAccountNumberByCountry(countryCode string) AccountNumber {
	path := req.buildPath(fmt.Sprintf(BANK_BY_COUNTRIES, countryCode))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var accountNumber AccountNumber
	parseRequestBody(res.Body, &accountNumber)
	return accountNumber
}

func ValidateVAT(countryCode, vat string) ValidatorResult {
	path := req.buildPath(fmt.Sprintf(VAT_VALIDATOR, countryCode, vat))
	res, _ := req.Post(path, "")
	defer res.Body.Close()
	var validatorResultMapResponse apiMapResponse
	parseRequestBody(res.Body, &validatorResultMapResponse)
	validatorResult := newValidatorResult(validatorResultMapResponse)
	return *validatorResult
}
