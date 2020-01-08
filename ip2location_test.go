package main

import (
	"fmt"
	"github.com/ip2location/ip2location-go"
)

// This an example of how to use the IP2Location Go Package to
// query the BIN database file to get geolocation data.
func main() {
	ip2location.Open("./IPV6-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE.BIN")
	ip := "8.8.8.8"

	results := ip2location.Get_all(ip)

	fmt.Printf("country_short: %s\n", results.Country_short)
	fmt.Printf("country_long: %s\n", results.Country_long)
	fmt.Printf("region: %s\n", results.Region)
	fmt.Printf("city: %s\n", results.City)
	fmt.Printf("isp: %s\n", results.Isp)
	fmt.Printf("latitude: %f\n", results.Latitude)
	fmt.Printf("longitude: %f\n", results.Longitude)
	fmt.Printf("domain: %s\n", results.Domain)
	fmt.Printf("zipcode: %s\n", results.Zipcode)
	fmt.Printf("timezone: %s\n", results.Timezone)
	fmt.Printf("netspeed: %s\n", results.Netspeed)
	fmt.Printf("iddcode: %s\n", results.Iddcode)
	fmt.Printf("areacode: %s\n", results.Areacode)
	fmt.Printf("weatherstationcode: %s\n", results.Weatherstationcode)
	fmt.Printf("weatherstationname: %s\n", results.Weatherstationname)
	fmt.Printf("mcc: %s\n", results.Mcc)
	fmt.Printf("mnc: %s\n", results.Mnc)
	fmt.Printf("mobilebrand: %s\n", results.Mobilebrand)
	fmt.Printf("elevation: %f\n", results.Elevation)
	fmt.Printf("usagetype: %s\n", results.Usagetype)
	fmt.Printf("api version: %s\n", ip2location.Api_version())

	ip2location.Close()
	// Output:
	// country_short: US
	// country_long: United States
	// region: California
	// city: Mountain View
	// isp: Google LLC
	// latitude: 37.405991
	// longitude: -122.078514
	// domain: google.com
	// zipcode: 94043
	// timezone: -07:00
	// netspeed: T1
	// iddcode: 1
	// areacode: 650
	// weatherstationcode: USCA0746
	// weatherstationname: Mountain View
	// mcc: -
	// mnc: -
	// mobilebrand: -
	// elevation: 32.000000
	// usagetype: DCH
	// api version: 8.2.0
}
