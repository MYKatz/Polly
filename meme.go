package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"image"
	//"image/gif"
	//"golang.org/x/image/font"
)

func getMeme(search string) string {
	var media string
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			media = ""
		}
	}()
	var response map[string]interface{}
	escaped := url.PathEscape(search)
	url := fmt.Sprintf("https://api.tenor.com/v1/search?q=%s&key=%s&limit=1", escaped, tenorKey)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		panic(fmt.Errorf("Tenor request failed: %s", err))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("Parse error: %s", err))
	}
	json.Unmarshal(body, &response)
	fmt.Println(response)
	media = response["results"].([]interface{})[0].(map[string]interface{})["media"].([]interface{})[0].(map[string]interface{})["gif"].(map[string]interface{})["url"].(string) //jesus christ what the fuck
	return media
}
