package md

type Status struct {
	Name    string        `json:"name"`
	Address int           `json:"address"`
	Datas   []*StatusData `json:"data"`
}

type StatusData struct {
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type Value struct {
	Name           string `json:"name"`
	Address        int    `json:"address"`
	Decimal_places string `json:"decimal_places"`
}
