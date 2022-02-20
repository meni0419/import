package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlStr := "https://testdb.kpi-drive.ru/_api/facts/save_fact"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "d0f00715-09ad-4808-b7a7-a7208e90bdec"
	var stags string
	var indicator_to_mo_id string
	var fact_value string

	stags = "sss"
	indicator_to_mo_id = "995"
	fact_value = "212"
	form := url.Values{}
	form.Add("period_start", "2022-02-01")
	form.Add("period_end", "2022-02-28")
	form.Add("period_key", "month")
	form.Add("auth_user_id", "4")
	form.Add("indicator_to_mo_fact_id", "0")
	form.Add("fact_time", "2022-02-28")
	form.Add("value", fact_value)
	form.Add("indicator_to_mo_id", indicator_to_mo_id)
	form.Add("comment", "приложение ММ на Го")
	form.Add("plan", "0")
	form.Add("request_id", "")
	form.Add("supertags", `[{"tag":{"id":2000,"name":"Валюта","key":"З","values":[{"title":"$","value":"$"},{"title":"грн.","value":"грн."},{"title":"руб.","value":"руб."}],"values_source":0},"value":"`+stags+`","value_title":null}]`)

	// Create a new request using http
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}
