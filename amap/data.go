package amap


type address struct {
	FormattedAddress string `json:"formatted_address"`
	//Country          string `json:"country"`
	//Province         string `json:"province"`
	//City             string `json:"city"`
}


type result struct {
	Status   string    `json:"status"`
	Info     string    `json:"info"`
	Infocode string    `json:"infocode"`
	Count    string    `json:"count"`
	Geocodes []address `json:"geocodes"`
}

type Client struct {
	Key 	string
	Secret  string
	ApiBase string
}
