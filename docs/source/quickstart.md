# Quickstart

## Dependencies

This library requires IP2Location BIN database to function. You may download the BIN database at

-   IP2Location LITE BIN Data (Free): <https://lite.ip2location.com>
-   IP2Location Commercial BIN Data (Comprehensive):
    <https://www.ip2location.com>

## IPv4 BIN vs IPv6 BIN

Use the IPv4 BIN file if you just need to query IPv4 addresses.

Use the IPv6 BIN file if you need to query BOTH IPv4 and IPv6 addresses.

## Installation

To install this module type the following:

```bash

go get github.com/ip2location/ip2location-go/v9

```

## Sample Codes

### Query geolocation information from BIN database

You can query the geolocation information from the IP2Location BIN database as below:

```go
package main

import (
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
)

func main() {
	db, err := ip2location.OpenDB("./IP-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE-ADDRESSTYPE-CATEGORY-DISTRICT-ASN.BIN")
	
	if err != nil {
		fmt.Print(err)
		return
	}
	ip := "8.8.8.8"
	results, err := db.Get_all(ip)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
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
	fmt.Printf("addresstype: %s\n", results.Addresstype)
	fmt.Printf("category: %s\n", results.Category)
	fmt.Printf("district: %s\n", results.District)
	fmt.Printf("asn: %s\n", results.Asn)
	fmt.Printf("as: %s\n", results.As)
	fmt.Printf("api version: %s\n", ip2location.Api_version())
	
	db.Close()
}
```

### Processing IP address using IP Tools class

You can manupulate IP address, IP number and CIDR as below:

```go
package main

import (
	"github.com/ip2location/ip2location-go/v9"
	"fmt"
	"math/big"
)

func main() {
	t := ip2location.OpenTools()
	
	ip := "8.8.8.8"
	res := t.IsIPv4(ip)
	
	fmt.Printf("Is IPv4: %t\n", res)
	
	ipnum, err := t.IPv4ToDecimal(ip)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("IPNum: %v\n", ipnum)
	}

	ip2 := "2600:1f18:45b0:5b00:f5d8:4183:7710:ceec"
	res2 := t.IsIPv6(ip2)
	
	fmt.Printf("Is IPv6: %t\n", res2)

	ipnum2, err := t.IPv6ToDecimal(ip2)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("IPNum: %v\n", ipnum2)
	}
	
	ipnum3 := big.NewInt(42534)
	res3, err := t.DecimalToIPv4(ipnum3)
	
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("IPv4: %v\n", res3)
	}
	
	ipnum4, ok := big.NewInt(0).SetString("22398978840339333967292465152", 10)
	if ok {
		res4, err := t.DecimalToIPv6(ipnum4)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("IPv6: %v\n", res4)
		}
	}
	
	ip3 := "2600:1f18:045b:005b:f5d8:0:000:ceec"
	res5, err := t.CompressIPv6(ip3)
	
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Compressed: %v\n", res5)
	}
	
	ip4 := "::45b:05b:f5d8:0:000:ceec"
	res6, err := t.ExpandIPv6(ip4)
	
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Expanded: %v\n", res6)
	}
	
	res7, err := t.IPv4ToCIDR("10.0.0.0", "10.10.2.255")
	
	if err != nil {
		fmt.Print(err)
	} else {
		for _, element := range res7 {
			fmt.Println(element)
		}
	}
	
	res8, err := t.IPv6ToCIDR("2001:4860:4860:0000:0000:0000:0000:8888", "2001:4860:4860:0000:eeee:ffff:ffff:ffff")
	
	if err != nil {
		fmt.Print(err)
	} else {
		for _, element := range res8 {
			fmt.Println(element)
		}
	}
	
	res9, err := t.CIDRToIPv4("123.245.99.13/26")
	
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("IPv4 Range: %v\n", res9)
	}
	
	res10, err := t.CIDRToIPv6("2002:1234::abcd:ffff:c0a8:101/62")
	
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("IPv6 Range: %v\n", res10)
	}
}
```

### List down country information

You can query country information for a country from IP2Location Country Information CSV file as below:

```go
package main

import (
	"github.com/ip2location/ip2location-go"
	"fmt"
)

func main() {
	c, err := ip2location.OpenCountryInfo("./IP2LOCATION-COUNTRY-INFORMATION.CSV")

	if err != nil {
		fmt.Print(err)
		return
	}

	res, err := c.GetCountryInfo("US")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("country_code: %s\n", res[0].Country_code)
	fmt.Printf("country_name: %s\n", res[0].Country_name)
	fmt.Printf("country_alpha3_code: %s\n", res[0].Country_alpha3_code)
	fmt.Printf("country_numeric_code: %s\n", res[0].Country_numeric_code)
	fmt.Printf("capital: %s\n", res[0].Capital)
	fmt.Printf("country_demonym: %s\n", res[0].Country_demonym)
	fmt.Printf("total_area: %s\n", res[0].Total_area)
	fmt.Printf("population: %s\n", res[0].Population)
	fmt.Printf("idd_code: %s\n", res[0].Idd_code)
	fmt.Printf("currency_code: %s\n", res[0].Currency_code)
	fmt.Printf("currency_name: %s\n", res[0].Currency_name)
	fmt.Printf("currency_symbol: %s\n", res[0].Currency_symbol)
	fmt.Printf("lang_code: %s\n", res[0].Lang_code)
	fmt.Printf("lang_name: %s\n", res[0].Lang_name)
	fmt.Printf("cctld: %s\n", res[0].Cctld)
	fmt.Print("==============================================\n")

	res2, err := c.GetCountryInfo()

	if err != nil {
		fmt.Print(err)
		return
	}

	for _, v := range res2 {
		fmt.Printf("country_code: %s\n", v.Country_code)
		fmt.Printf("country_name: %s\n", v.Country_name)
		fmt.Printf("country_alpha3_code: %s\n", v.Country_alpha3_code)
		fmt.Printf("country_numeric_code: %s\n", v.Country_numeric_code)
		fmt.Printf("capital: %s\n", v.Capital)
		fmt.Printf("country_demonym: %s\n", v.Country_demonym)
		fmt.Printf("total_area: %s\n", v.Total_area)
		fmt.Printf("population: %s\n", v.Population)
		fmt.Printf("idd_code: %s\n", v.Idd_code)
		fmt.Printf("currency_code: %s\n", v.Currency_code)
		fmt.Printf("currency_name: %s\n", v.Currency_name)
		fmt.Printf("currency_symbol: %s\n", v.Currency_symbol)
		fmt.Printf("lang_code: %s\n", v.Lang_code)
		fmt.Printf("lang_name: %s\n", v.Lang_name)
		fmt.Printf("cctld: %s\n", v.Cctld)
		fmt.Print("==============================================\n")
	}
}
```

### List down region information

You can get the region code by country code and region name from IP2Location ISO 3166-2 Subdivision Code CSV file as below:

```go
package main

import (
	"github.com/ip2location/ip2location-go"
	"fmt"
)

func main() {
	r, err := ip2location.OpenRegionInfo("./IP2LOCATION-ISO3166-2.CSV")

	if err != nil {
		fmt.Print(err)
		return
	}

	res, err := r.GetRegionCode("US", "California")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("region code: %s\n", res)
}
```