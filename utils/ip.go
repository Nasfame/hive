package utils

import (
	"io/ioutil"
	"net/http"
)

func GetPublicIP() string {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(ip)
}
