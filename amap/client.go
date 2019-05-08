package amap

const apiBase = "https://restapi.amap.com/v3/geocode/geo?"


func NewClient(key, secret string) *Client {
	client := &Client{
		Key: 		key,
		Secret: 	secret,
		ApiBase:	apiBase,
	}

	return client
}


func (c *Client) Address(text string) (address string, formatted bool) {

	queryStr := genQueryStr(text, "", c.Key)
	sign := genSign(queryStr, c.Secret)
	address = request(c.ApiBase, queryStr, sign)

	if address != "" {
		formatted = true
	}

	return address, formatted
}
