package randommer

import "fmt"

func GetCultures() []Culture {
	path := req.buildPath(CULTURES)
	res, _ := req.Get(path)
	defer res.Body.Close()
	var cultureArr []apiMapResponse
	parseRequestBody(res.Body, &cultureArr)
	cultures := makeCultures(cultureArr)
	return cultures
}

func GetRandomAddresses(quantity int, culture string) []RandomAddress {
	path := req.buildPath(fmt.Sprintf(RANDOM_ADDRESS, quantity, culture))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var randomAddresses []RandomAddress
	parseRequestBody(res.Body, &randomAddresses)
	return randomAddresses
}
