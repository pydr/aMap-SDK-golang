package amap

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// 查询字符串组装
func genQueryStr(address, city, key string) string {
	var queryMap = make(map[string]string)

	queryMap["address"] = address
	queryMap["city"] = city
	queryMap["key"] = key

	var queryKeys []string
	for k := range queryMap {
		queryKeys = append(queryKeys, k)
	}

	sort.Strings(queryKeys)

	var queryStr string
	for _, queryKey := range queryKeys {
		queryStr += queryKey + "=" + queryMap[queryKey] + "&"
	}

	queryStr = strings.Trim(queryStr, "&")

	return queryStr
}

// 生成签名
func genSign(queryStr, secret string) string {

	str := queryStr + secret

	m := md5.New()
	m.Write([]byte(str))

	sign := hex.EncodeToString(m.Sum(nil))

	return sign
}

// 请求服务
func request(server, query, sign string) (address string) {

	url := server + query + "&" + sign

	//logs.Warn("url： ", url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	//logs.Warn(string(body))

	var result result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	if result.Status == "1" {
		if len(result.Geocodes) > 0 {
			address = result.Geocodes[0].FormattedAddress
		}
	}

	return

}
