[![Go Report Card](https://goreportcard.com/badge/github.com/ip2location/ip2location-go/v9)](https://goreportcard.com/report/github.com/ip2location/ip2location-go/v9)


# IP2Location Go Package

This Go package provides a fast lookup of country, region, city, latitude, longitude, ZIP code, time zone, ISP, domain name, connection type, IDD code, area code, weather station code, station name, mcc, mnc, mobile brand, elevation, usage type, address type and IAB category from IP address by using IP2Location database. This package uses a file based database available at IP2Location.com. This database simply contains IP blocks as keys, and other information such as country, region, city, latitude, longitude, ZIP code, time zone, ISP, domain name, connection type, IDD code, area code, weather station code, station name, mcc, mnc, mobile brand, elevation, usage type, address type and IAB category as values. It supports both IP address in IPv4 and IPv6.

This package can be used in many types of projects such as:

 - select the geographically closest mirror
 - analyze your web server logs to determine the countries of your visitors
 - credit card fraud detection
 - software export controls
 - display native language and currency 
 - prevent password sharing and abuse of service 
 - geotargeting in advertisement

The database will be updated in monthly basis for the greater accuracy. Free LITE databases are available at https://lite.ip2location.com/ upon registration.

The paid databases are available at https://www.ip2location.com under Premium subscription package.

As an alternative, this package can also call the IP2Location Web Service. This requires an API key. If you don't have an existing API key, you can subscribe for one at the below:

https://www.ip2location.com/web-service/ip2location

## Installation

```
go get github.com/ip2location/ip2location-go/v9
```

## QUERY USING THE BIN FILE

## Dependencies

This package requires IP2Location BIN data file to function. You may download the BIN data file at
* IP2Location LITE BIN Data (Free): https://lite.ip2location.com
* IP2Location Commercial BIN Data (Comprehensive): https://www.ip2location.com


## IPv4 BIN vs IPv6 BIN

Use the IPv4 BIN file if you just need to query IPv4 addresses.

Use the IPv6 BIN file if you need to query BOTH IPv4 and IPv6 addresses.


## Methods

Below are the methods supported in this package.

|Method Name|Description|
|---|---|
|OpenDB|Initialize the package with the BIN file.|
|Get_all|Returns the geolocation information in an object.|
|Get_country_short|Returns the country code.|
|Get_country_long|Returns the country name.|
|Get_region|Returns the region name.|
|Get_city|Returns the city name.|
|Get_isp|Returns the ISP name.|
|Get_latitude|Returns the latitude.|
|Get_longitude|Returns the longitude.|
|Get_domain|Returns the domain name.|
|Get_zipcode|Returns the ZIP code.|
|Get_timezone|Returns the time zone.|
|Get_netspeed|Returns the net speed.|
|Get_iddcode|Returns the IDD code.|
|Get_areacode|Returns the area code.|
|Get_weatherstationcode|Returns the weather station code.|
|Get_weatherstationname|Returns the weather station name.|
|Get_mcc|Returns the mobile country code.|
|Get_mnc|Returns the mobile network code.|
|Get_mobilebrand|Returns the mobile brand.|
|Get_elevation|Returns the elevation in meters.|
|Get_usagetype|Returns the usage type.|
|Get_addresstype|Returns the address type.|
|Get_category|Returns the IAB category.|
|Close|Closes BIN file.|

## Usage

```go
package main

import (
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
)

func main() {
	db, err := ip2location.OpenDB("./IP-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE-ADDRESSTYPE-CATEGORY.BIN")
	
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
	fmt.Printf("api version: %s\n", ip2location.Api_version())
	
	db.Close()
}
```

## QUERY USING THE IP2LOCATION WEB SERVICE

## Methods
Below are the methods supported in this package.

|Method Name|Description|
|---|---|
|OpenWS| 3 input parameters:<ol><li>IP2Location API Key.</li><li>Package (WS1 - WS25)</li></li><li>Use HTTPS or HTTP</li></ol> |
|LookUp|Query IP address. This method returns an object containing the geolocation info. <ul><li>country_code</li><li>country_name</li><li>region_name</li><li>city_name</li><li>latitude</li><li>longitude</li><li>zip_code</li><li>time_zone</li><li>isp</li><li>domain</li><li>net_speed</li><li>idd_code</li><li>area_code</li><li>weather_station_code</li><li>weather_station_name</li><li>mcc</li><li>mnc</li><li>mobile_brand</li><li>elevation</li><li>usage_type</li><li>address_type</li><li>category</li><li>continent<ul><li>name</li><li>code</li><li>hemisphere</li><li>translations</li></ul></li><li>country<ul><li>name</li><li>alpha3_code</li><li>numeric_code</li><li>demonym</li><li>flag</li><li>capital</li><li>total_area</li><li>population</li><li>currency<ul><li>code</li><li>name</li><li>symbol</li></ul></li><li>language<ul><li>code</li><li>name</li></ul></li><li>idd_code</li><li>tld</li><li>is_eu</li><li>translations</li></ul></li><li>region<ul><li>name</li><li>code</li><li>translations</li></ul></li><li>city<ul><li>name</li><li>translations</li></ul></li><li>geotargeting<ul><li>metro</li></ul></li><li>country_groupings</li><li>time_zone_info<ul><li>olson</li><li>current_time</li><li>gmt_offset</li><li>is_dst</li><li>sunrise</li><li>sunset</li></ul></li><ul>|
|GetCredit|This method returns the web service credit balance in an object.|

## Usage

```go
package main

import (
	"github.com/ip2location/ip2location-go/v9"
	"fmt"
)

func main() {
	apikey := "YOUR_API_KEY"
	apipackage := "WS25"
	usessl := true
	addon := "continent,country,region,city,geotargeting,country_groupings,time_zone_info" // leave blank if no need
	lang := "en" // leave blank if no need
	
	ws, err := ip2location.OpenWS(apikey, apipackage, usessl)
	
	if err != nil {
		fmt.Print(err)
		return
	}
	ip := "8.8.8.8"
	res, err := ws.LookUp(ip, addon, lang)

	if err != nil {
		fmt.Print(err)
		return
	}

	if res.Response != "OK" {
		fmt.Printf("Error: %s\n", res.Response)
	} else {
		// standard results
		fmt.Printf("Response: %s\n", res.Response)
		fmt.Printf("CountryCode: %s\n", res.CountryCode)
		fmt.Printf("CountryName: %s\n", res.CountryName)
		fmt.Printf("RegionName: %s\n", res.RegionName)
		fmt.Printf("CityName: %s\n", res.CityName)
		fmt.Printf("Latitude: %f\n", res.Latitude)
		fmt.Printf("Longitude: %f\n", res.Longitude)
		fmt.Printf("ZipCode: %s\n", res.ZipCode)
		fmt.Printf("TimeZone: %s\n", res.TimeZone)
		fmt.Printf("Isp: %s\n", res.Isp)
		fmt.Printf("Domain: %s\n", res.Domain)
		fmt.Printf("NetSpeed: %s\n", res.NetSpeed)
		fmt.Printf("IddCode: %s\n", res.IddCode)
		fmt.Printf("AreaCode: %s\n", res.AreaCode)
		fmt.Printf("WeatherStationCode: %s\n", res.WeatherStationCode)
		fmt.Printf("WeatherStationName: %s\n", res.WeatherStationName)
		fmt.Printf("Mcc: %s\n", res.Mcc)
		fmt.Printf("Mnc: %s\n", res.Mnc)
		fmt.Printf("MobileBrand: %s\n", res.MobileBrand)
		fmt.Printf("Elevation: %d\n", res.Elevation)
		fmt.Printf("UsageType: %s\n", res.UsageType)
		fmt.Printf("AddressType: %s\n", res.AddressType)
		fmt.Printf("Category: %s\n", res.Category)
		fmt.Printf("CategoryName: %s\n", res.CategoryName)
		fmt.Printf("CreditsConsumed: %d\n", res.CreditsConsumed)
		
		// continent addon
		fmt.Printf("Continent => Name: %s\n", res.Continent.Name)
		fmt.Printf("Continent => Code: %s\n", res.Continent.Code)
		fmt.Printf("Continent => Hemisphere: %+v\n", res.Continent.Hemisphere)
		
		// country addon
		fmt.Printf("Country => Name: %s\n", res.Country.Name)
		fmt.Printf("Country => Alpha3Code: %s\n", res.Country.Alpha3Code)
		fmt.Printf("Country => NumericCode: %s\n", res.Country.NumericCode)
		fmt.Printf("Country => Demonym: %s\n", res.Country.Demonym)
		fmt.Printf("Country => Flag: %s\n", res.Country.Flag)
		fmt.Printf("Country => Capital: %s\n", res.Country.Capital)
		fmt.Printf("Country => TotalArea: %s\n", res.Country.TotalArea)
		fmt.Printf("Country => Population: %s\n", res.Country.Population)
		fmt.Printf("Country => IddCode: %s\n", res.Country.IddCode)
		fmt.Printf("Country => Tld: %s\n", res.Country.Tld)
		fmt.Printf("Country => IsEu: %t\n", res.Country.IsEu)
		
		fmt.Printf("Country => Currency => Code: %s\n", res.Country.Currency.Code)
		fmt.Printf("Country => Currency => Name: %s\n", res.Country.Currency.Name)
		fmt.Printf("Country => Currency => Symbol: %s\n", res.Country.Currency.Symbol)
		
		fmt.Printf("Country => Language => Code: %s\n", res.Country.Language.Code)
		fmt.Printf("Country => Language => Name: %s\n", res.Country.Language.Name)
		
		// region addon
		fmt.Printf("Region => Name: %s\n", res.Region.Name)
		fmt.Printf("Region => Code: %s\n", res.Region.Code)
		
		// city addon
		fmt.Printf("City => Name: %s\n", res.City.Name)
		
		// geotargeting addon
		fmt.Printf("Geotargeting => Metro: %s\n", res.Geotargeting.Metro)
		
		// country_groupings addon
		for i, s := range res.CountryGroupings {
			fmt.Printf("CountryGroupings => #%d => Acronym: %s\n", i, s.Acronym)
			fmt.Printf("CountryGroupings => #%d => Name: %s\n", i, s.Name)
		}
		
		// time_zone_info addon
		fmt.Printf("TimeZoneInfo => Olson: %s\n", res.TimeZoneInfo.Olson)
		fmt.Printf("TimeZoneInfo => CurrentTime: %s\n", res.TimeZoneInfo.CurrentTime)
		fmt.Printf("TimeZoneInfo => GmtOffset: %d\n", res.TimeZoneInfo.GmtOffset)
		fmt.Printf("TimeZoneInfo => IsDst: %s\n", res.TimeZoneInfo.IsDst)
		fmt.Printf("TimeZoneInfo => Sunrise: %s\n", res.TimeZoneInfo.Sunrise)
		fmt.Printf("TimeZoneInfo => Sunset: %s\n", res.TimeZoneInfo.Sunset)
	}

	res2, err := ws.GetCredit()

	if err != nil {
		fmt.Print(err)
		return
	}
	
	fmt.Printf("Credit Balance: %d\n", res2.Response)
}
```

## IPTOOLS CLASS

## Methods
Below are the methods supported in this package.

|Method Name|Description|
|---|---|
|func (t *IPTools) IsIPv4(IP string) bool|Returns true if string contains an IPv4 address. Otherwise false.|
|func (t *IPTools) IsIPv6(IP string) bool|Returns true if string contains an IPv6 address. Otherwise false.|
|func (t *IPTools) IPv4ToDecimal(IP string) (*big.Int, error)|Returns the IP number for an IPv4 address.|
|func (t *IPTools) IPv6ToDecimal(IP string) (*big.Int, error)|Returns the IP number for an IPv6 address.|
|func (t *IPTools) DecimalToIPv4(IPNum *big.Int) (string, error)|Returns the IPv4 address for the supplied IP number.|
|func (t *IPTools) DecimalToIPv6(IPNum *big.Int) (string, error)|Returns the IPv6 address for the supplied IP number.|
|func (t *IPTools) CompressIPv6(IP string) (string, error)|Returns the IPv6 address in compressed form.|
|func (t *IPTools) ExpandIPv6(IP string) (string, error)|Returns the IPv6 address in expanded form.|
|func (t *IPTools) IPv4ToCIDR(IPFrom string, IPTo string) ([]string, error)|Returns a list of CIDR from the supplied IPv4 range.|
|func (t *IPTools) IPv6ToCIDR(IPFrom string, IPTo string) ([]string, error)|Returns a list of CIDR from the supplied IPv6 range.|
|func (t *IPTools) CIDRToIPv4(CIDR string) ([]string, error)|Returns the IPv4 range from the supplied CIDR.|
|func (t *IPTools) CIDRToIPv6(CIDR string) ([]string, error)|Returns the IPv6 range from the supplied CIDR.|

## Usage

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

## COUNTRY CLASS

## Methods
Below are the methods supported in this package.

|Method Name|Description|
|---|---|
|func OpenCountryInfo(csvFile string) (*CI, error)|Expect a IP2Location Country Information CSV file. This database is free for download at https://www.ip2location.com/free/country-information|
|func (c *CI) GetCountryInfo(countryCode ...string) ([]CountryInfoRecord, error)|Returns the country information for specified country or all countries.|

## Usage

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

## REGION CLASS

## Methods
Below are the methods supported in this package.

|Method Name|Description|
|---|---|
|func OpenRegionInfo(csvFile string) (*RI, error)|Expect a IP2Location ISO 3166-2 Subdivision Code CSV file. This database is free for download at https://www.ip2location.com/free/iso3166-2|
|func (r *RI) GetRegionCode(countryCode string, regionName string) (string, error)|Returns the region code for the supplied country code and region name.|

## Usage

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