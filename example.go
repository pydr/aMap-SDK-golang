package main

import (
	"aMap-SDK-golang/amap"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	client := amap.NewClient("yourKey", "yourSecret")

	var o *os.File
	o, err := os.OpenFile("addr_new.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		println("open file failed.")
		return
	}

	buf, err := ioutil.ReadFile("addr.txt")

	if err != nil {
		println("read failed.", err)
		return
	}

	content := string(buf)

	address := strings.Split(content, "\n")

	defer o.Close()
	for _, addr := range address {
		formatAddr, formatted := client.Address(addr)
		formatAddr = strings.Replace(formatAddr, "|", "", -1)

		if formatted || formatAddr != "" {
			println("正在写入： ", formatAddr)
			formatAddr = formatAddr + "\n"
			_, err = o.WriteString(formatAddr)
			if err != nil {
				println("%v 写入失败: %v", formatAddr, err)
				return
			}
		} else {
			println("正在写入： ", addr)
			addr = addr + "\n"
			_, err = o.WriteString(addr)
			if err != nil {
				println("%v 写入失败: %v", formatAddr, err)
				return
			}
		}
	}

	return
}
