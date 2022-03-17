package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Address struct {
	Query      string
	Country    string
	RegionName string
	City       string
	Zip        string
	Lat        float32
	Lon        float32
	Isp        string
}

type Options struct {
	Ip string
}

func getOptions() Options {
	ip := flag.String("ip", "", "ip address to scan")

	flag.Parse()

	return Options{*ip}
}

func fetchResponse(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

// empty string will fetch current IP
func fetchAddress(ip string) {
	var response Address
	resp := fetchResponse("http://ip-api.com/json/" + ip)

	json.Unmarshal([]byte(resp), &response)

	print(response, ip)
}

func print(data Address, ip string) {
	fmt.Println("IP: ", data.Query)
	fmt.Println("country: ", data.Country)
	fmt.Println("region: ", data.RegionName)
	fmt.Println("city: ", data.City)
	fmt.Println("zip: ", data.Zip)
	fmt.Println("lat: ", data.Lat)
	fmt.Println("long: ", data.Lon)
	fmt.Println("isp: ", data.Isp)
}

func main() {
	options := getOptions()

	fetchAddress(options.Ip)
}
