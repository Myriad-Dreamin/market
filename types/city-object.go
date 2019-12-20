package types

type CityObject struct {
	Province string `json:"province"`
	Name string `json:"name"`
	ID string `json:"id"`
}

var CityObjectMap map[string]CityObject

func SetCityObjectMap(cmp map[string]CityObject) {
	CityObjectMap = cmp
}
