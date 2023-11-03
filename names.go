package randommer

import "fmt"

func GetRandomNames(nameType string, quantity int) []Name {
	path := req.buildPath(fmt.Sprintf(NAMES, nameType, quantity))
	res, _ := req.Get(path)
	defer res.Body.Close()
	var names []Name
	parseRequestBody(res.Body, &names)
	return names
}
